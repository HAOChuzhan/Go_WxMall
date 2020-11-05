package v2

/*
相关微信登录验证和订单
*/
import (
	"encoding/json"
	"net/http"
	"wx-gin-master/models/auth"
	"wx-gin-master/models/order"
	"wx-gin-master/models/user"
	"wx-gin-master/models/wallet"
	"wx-gin-master/models/wxorder"
	"wx-gin-master/pkg/app"
	"wx-gin-master/pkg/e"
	"wx-gin-master/pkg/logging"
	"wx-gin-master/pkg/setting"
	"wx-gin-master/pkg/util"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/medivhzhan/weapp"
	"github.com/medivhzhan/weapp/payment"
	"github.com/unknwon/com"
)

func WxLogin(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	httpCode := http.StatusInternalServerError
	errCode := e.InternalServerError

	valid.Required(c.Query("code"), "code").Message("code 必须")
	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.BadRequest, valid.Errors)
		return
	}

	wxCode := c.Query("code") //此处的是查询字符串query string，此处获取code，再去交换res
	res, err := weapp.Login(setting.WxappSetting.AppID, setting.WxappSetting.Secret, wxCode)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ErrorWxLoginFain, valid.Errors)
		return
	}
	if res.UnionID != "" {
		if userInfo := user.QueryUserByUnionid(res.UnionID); !(userInfo.ID > 0) {
			find := user.QueryUserByOpenid(res.OpenID)
			if find.ID > 0 {
				user.AddUnionid(find.ID, res.UnionID)
			} else {
				newUser := user.User{Openid: res.OpenID, Unionid: res.UnionID}
				if !user.Create(&newUser) {
					httpCode = http.StatusInternalServerError
					errCode = e.ErrorCreateUserFail
				}
			}
		}
	} else if res.OpenID != "" {
		if userInfo := user.QueryUserByOpenid(res.OpenID); !(userInfo.ID > 0) {
			newUser := user.User{Openid: res.OpenID, Unionid: res.UnionID}
			if !user.Create(&newUser) {
				httpCode = http.StatusInternalServerError
				errCode = e.ErrorCreateUserFail
			}
		}
	}
	if !user.PutSsk(res.OpenID, res.SessionKey) {
		httpCode = http.StatusInternalServerError
		errCode = e.ErrorPutSskFain
	}
	id := auth.CheckOpenid(res.OpenID)
	data := make(map[string]interface{})
	if id > 0 {
		token, err := util.GenerateToken(id)
		if err != nil {
			httpCode = http.StatusInternalServerError
			errCode = e.ErrorAuthToken
		} else {
			httpCode = http.StatusOK
			errCode = e.Created
			data["token"] = token
		}
	}
	appG.Response(httpCode, errCode, data)
}

func WxGetUserInfo(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	userId := c.MustGet("AuthData").(*util.Claims).User.ID

	valid.Required(c.PostForm("rawData"), "rawData").Message("rawData 必须")
	valid.Required(c.PostForm("signature"), "signature").Message("signature 必须")
	valid.Required(c.PostForm("encryptedData"), "encryptedData").Message("encryptedData 必须")
	valid.Required(c.PostForm("iv"), "iv").Message("iv 必须")
	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.BadRequest, valid.Errors)
		return
	}

	userInfo := user.QueryUserById(userId)
	rawData := c.PostForm("rawData")
	encryptedData := c.PostForm("encryptedData")
	signature := c.PostForm("signature")
	iv := c.PostForm("iv")
	ssk := userInfo.SessionKey
	ui, err := weapp.DecryptUserInfo(rawData, encryptedData, signature, iv, ssk)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ErrorDecryptUserInfoFain, nil)
		return
	}
	UserInfo := user.User{
		Avatar:   ui.Avatar,
		Nickname: ui.Nickname,
		Sex:      ui.Gender,
	}
	if !user.Update(&userInfo, UserInfo) {
		appG.Response(http.StatusInternalServerError, e.ErrorUpdateUserFail, nil)
		return
	}
	userInfo = user.QueryUserById(userId)
	appG.Response(http.StatusOK, e.OK, userInfo)
}

func WxGetPhone(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	userId := c.MustGet("AuthData").(*util.Claims).User.ID

	valid.Required(c.PostForm("encryptedData"), "encryptedData").Message("encryptedData 必须")
	valid.Required(c.PostForm("iv"), "iv").Message("iv 必须")
	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.BadRequest, valid.Errors)
		return
	}

	userInfo := user.QueryUserById(userId)
	encryptedData := c.PostForm("encryptedData")
	iv := c.PostForm("iv")
	ssk := userInfo.SessionKey

	phone, err := weapp.DecryptPhoneNumber(ssk, encryptedData, iv)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ErrorDecryptUserInfoFain, nil)
		return
	}

	if !user.UpdateColumn(&userInfo, "mobile", phone.PurePhoneNumber) {
		appG.Response(http.StatusInternalServerError, e.ErrorUpdateUserFail, nil)
		return
	}

	userInfo = user.QueryUserById(userId)
	appG.Response(http.StatusOK, e.OK, userInfo)
}

func WxTopup(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	userId := c.MustGet("AuthData").(*util.Claims).User.ID

	valid.Required(c.PostForm("amount"), "amount").Message("amount 必须")
	valid.Numeric(c.PostForm("amount"), "amount").Message("amount 必须是有效数值,单位为分")
	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.BadRequest, valid.Errors)
		return
	}

	amount := com.StrTo(c.PostForm("amount")).MustInt()
	sumPay := float64(amount / 100)
	tradeNo := string(com.RandomCreateBytes(32))
	receiveUrl := "https://xxx.xxx.xxx"
	Body := "充值 " + com.ToStr(sumPay) + "RMB"
	userInfo := user.QueryUserById(userId)

	form := payment.Order{
		AppID:      setting.WxappSetting.AppID,
		MchID:      setting.WxappSetting.MchId,
		Body:       Body,
		NotifyURL:  receiveUrl,
		OpenID:     userInfo.Openid,
		OutTradeNo: tradeNo,
		TotalFee:   amount,
		Detail:     Body,
		Attach:     Body,
	}

	res, err := form.Unify(setting.WxappSetting.PayKey)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ErrorTopupFain, nil)
		return
	}

	params, err := payment.GetParams(res.AppID, setting.WxappSetting.PayKey, res.NonceStr, res.PrePayID)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ErrorGetParamsFain, nil)
		return
	}

	orderInfo := wxorder.WxOrder{
		UserId:     userId,
		OutTradeNo: tradeNo,
		NonceStr:   res.NonceStr,
		Sign:       res.Sign,
		Body:       Body,
		Detail:     form.Detail,
		Attach:     form.Attach,
		SumPay:     sumPay,
		TotalFee:   form.TotalFee,
	}
	if !wxorder.Create(&orderInfo) {
		appG.Response(http.StatusInternalServerError, e.ErrorCreateWxOrderFain, nil)
		return
	}
	appG.Response(http.StatusOK, e.OK, params)
}

func WxTopupCallback(c *gin.Context) {
	err := payment.HandlePaidNotify(c.Writer, c.Request, func(ntf payment.PaidNotify) (bool, string) {
		tradeNo := ntf.OutTradeNo
		wxOrder := wxorder.QueryByTradeNo(tradeNo)
		if wallet.TopUpBalance(wxOrder.SumPay, wxOrder.UserId) {
			wxOrder.TransactionId = ntf.TransactionID
			if wxorder.Update(&wxOrder) {
				if wallet.TopUpBalance(wxOrder.SumPay, wxOrder.UserId) && wxorder.Destroy(&wxOrder) {
					return true, ""
				}
			}
		}
		return false, "服务器内部更新错误"
	})
	if err != nil {
		logging.Info(err)
	}
}

type Info struct {
	Oid    int `json:"oid"`
	UserId int `json:"userId"`
}

func WxPayForOrder(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	userId := c.MustGet("AuthData").(*util.Claims).User.ID

	valid.Required(c.PostForm("oid"), "oid").Message("oid(Order Id) 必须")
	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.BadRequest, valid.Errors)
		return
	}

	oid := com.StrTo(c.PostForm("oid")).MustInt()
	if !order.IsOwner(userId, oid) {
		appG.Response(http.StatusInternalServerError, e.ErrorIsOwnerFail, nil)
		return
	}

	orderInfo := order.QueryOrderById(oid)

	totalFee := int(orderInfo.SumPay * 100)
	tradeNo := string(com.RandomCreateBytes(32))
	receiveUrl := "https://xxx.xxx.xxx"
	jsonData := Info{
		orderInfo.ID, userId,
	}
	jsonStr, err := json.Marshal(&jsonData)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ErrorJsonMarshalFail, nil)
		return
	}
	Body := string(jsonStr)
	userInfo := user.QueryUserById(userId)

	form := payment.Order{
		AppID:      setting.WxappSetting.AppID,
		MchID:      setting.WxappSetting.MchId,
		Body:       Body,
		NotifyURL:  receiveUrl,
		OpenID:     userInfo.Openid,
		OutTradeNo: tradeNo,
		TotalFee:   totalFee,
		Detail:     Body,
		Attach:     Body,
	}

	res, err := form.Unify(setting.WxappSetting.PayKey)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ErrorTopupFain, nil)
		return
	}

	params, err := payment.GetParams(res.AppID, setting.WxappSetting.PayKey, res.NonceStr, res.PrePayID)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ErrorGetParamsFain, nil)
		return
	}

	wxOrder := wxorder.WxOrder{
		UserId:     userId,
		OutTradeNo: tradeNo,
		NonceStr:   res.NonceStr,
		Sign:       res.Sign,
		Body:       Body,
		Detail:     form.Detail,
		Attach:     form.Attach,
		SumPay:     orderInfo.SumPay,
		TotalFee:   form.TotalFee,
	}
	if !wxorder.Create(&wxOrder) {
		appG.Response(http.StatusInternalServerError, e.ErrorCreateWxOrderFain, nil)
		return
	}
	appG.Response(http.StatusOK, e.OK, params)
}

func WxPayForOrderCallback(c *gin.Context) {
	err := payment.HandlePaidNotify(c.Writer, c.Request, func(ntf payment.PaidNotify) (bool, string) {
		tradeNo := ntf.OutTradeNo
		wxOrder := wxorder.QueryByTradeNo(tradeNo)
		if wallet.TopUpBalance(wxOrder.SumPay, wxOrder.UserId) {
			wxOrder.TransactionId = ntf.TransactionID
			if wxorder.Update(&wxOrder) {
				var info Info
				err := json.Unmarshal([]byte(wxOrder.Body), &info)
				if err != nil {
					logging.Info(err)
					return false, "解析Body失败"
				}
				if order.Settlement(info.UserId, info.Oid) && wxorder.Destroy(&wxOrder) {
					return true, ""
				}
			}
		}
		return false, "服务器内部更新错误"
	})
	if err != nil {
		logging.Info(err)
	}
}

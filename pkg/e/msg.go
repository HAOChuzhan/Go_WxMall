package e

var MsgFlags = map[int]string{
	OK:                  "请求成功",
	Created:             "成功请求并创建了新的资源",
	NoContent:           "无内容，服务器成功处理，但未返回内容",
	BadRequest:          "请求参数错误",
	Unauthorized:        "请求要求用户的身份认证",
	NotFound:            "未找到资源",
	InternalServerError: "服务器内部错误",

	OKUpdateBanner: "Banner修改成功",

	ErrorCheckAdminFail:  "用户名或密码错误",
	ErrorCreateProFain:   "添加新商品失败",
	ErrorUpDateProFain:   "修改商品信息失败",
	ErrorDestroyProFain:  "删除商品失败",
	ErrorUpdateOrderFain: "修改订单信息失败",

	ErrorWxLoginFain:         "微信服务器登录失败",
	ErrorPutSskFain:          "存放Sessionkey失败",
	ErrorDecryptUserInfoFain: "解析用户信息失败",
	ErrorTopupFain:           "支付失败",
	ErrorGetParamsFain:       "获取支付参数失败",
	ErrorCreateWxOrderFain:   "未能插入微信订单表",

	ErrorCreateUserFail:        "新建用户失败",
	ErrorLoginUserFail:         "用户登录失败",
	ErrorUpdateUserFail:        "用户信息更新失败",
	ErrorGetFavoritesListFail:  "获取收藏夹失败",
	ErrorAddFavoriteFail:       "添加商品到收藏夹失败",
	ErrorDestroyFavoriteFail:   "从收藏夹移除商品失败",
	ErrorPutInCartFail:         "添加到购物车失败",
	ErrorDropFromCartFail:      "删除购物车商品失败",
	ErrorDecreaseNumbFail:      "减少购物车商品数量失败",
	ErrorCreateAddressFail:     "新建收货地址失败",
	ErrorDestroyAddressFail:    "删除收货地址失败",
	ErrorUpdateAddressFail:     "更新收货地址失败",
	ErrorSetAddressDefaultFail: "设置默认收货地址失败",
	ErrorCreatedOrderFail:      "新建订单失败",
	ErrorDestroyOrderFail:      "销毁订单失败",
	ErrorSettlementFail:        "支付失败，余额不足",
	ErrorIsOwnerFail:           "订单与用户信息不符",

	ErrorAuthCheckTokenFail:    "Token鉴权失败",
	ErrorAuthCheckTokenTimeout: "Token已超时",
	ErrorAuthToken:             "Token生成失败",
	ErrorAuth:                  "Token错误",

	ErrorUploadSaveImageFail:    "保存图片失败",
	ErrorUploadCheckImageFail:   "检查图片失败",
	ErrorUploadCheckImageFormat: "校验图片错误，图片格式或大小有问题",

	ErrorJsonMarshalFail: "编码Json失败",
}

// GetMsg 根据代码获取错误信息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[InternalServerError]
}

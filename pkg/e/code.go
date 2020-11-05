package e

const (
	OK                  = 200 // 请求成功。一般用于GET与POST请求
	Created             = 201 // 已创建。成功请求并创建了新的资源
	NoContent           = 204 // 无内容。服务器成功处理，但未返回内容。在未更新网页的情况下，可确保浏览器继续显示当前文档
	BadRequest          = 400 // 请求参数错误
	Unauthorized        = 401 // 请求要求用户的身份认证
	NotFound            = 404 // 服务器无法根据客户端的请求找到资源（网页）。通过此代码，网站设计人员可设置"您所请求的资源无法找到"的个性页面
	InternalServerError = 500 // 服务器内部错误，无法完成请求

	OKUpdateBanner = 40001 //Banner修改成功

	ErrorCheckAdminFail  = 50001 // 用户名或密码错误
	ErrorCreateProFain   = 50002 // 添加新商品失败
	ErrorUpDateProFain   = 50003 // 修改商品信息失败
	ErrorDestroyProFain  = 50004 // 删除商品失败
	ErrorUpdateOrderFain = 50005 // 修改订单信息失败

	ErrorWxLoginFain         = 60001 // 微信服务器登录失败
	ErrorPutSskFain          = 60002 // 存放session key失败
	ErrorDecryptUserInfoFain = 60003 // 解析用户信息失败
	ErrorTopupFain           = 60004 // 支付失败
	ErrorGetParamsFain       = 60005 // 获取支付参数失败
	ErrorCreateWxOrderFain   = 60006 // 未能插入微信订单表

	ErrorCreateUserFail        = 10001 // 新建用户失败
	ErrorLoginUserFail         = 10002 // 用户登录失败
	ErrorUpdateUserFail        = 10003 // 用户信息更新失败
	ErrorGetFavoritesListFail  = 10004 // 获取收藏夹失败
	ErrorAddFavoriteFail       = 10005 // 添加商品到收藏夹失败
	ErrorDestroyFavoriteFail   = 10006 // 从收藏夹移除商品失败
	ErrorPutInCartFail         = 10007 // 添加到购物车失败
	ErrorDropFromCartFail      = 10008 // 删除购物车商品失败
	ErrorDecreaseNumbFail      = 10009 // 减少购物车商品数量失败
	ErrorCreateAddressFail     = 10010 // 新建收货地址失败
	ErrorDestroyAddressFail    = 10011 // 删除收货地址失败
	ErrorUpdateAddressFail     = 10012 // 更新收货地址失败
	ErrorSetAddressDefaultFail = 10013 // 设置默认收货地址失败
	ErrorCreatedOrderFail      = 10014 // 新建订单失败
	ErrorDestroyOrderFail      = 10015 // 销毁订单失败
	ErrorSettlementFail        = 10016 // 支付失败，余额不足
	ErrorIsOwnerFail           = 10017 // 订单与用户信息不符

	ErrorAuthCheckTokenFail    = 20001 // Token鉴权失败
	ErrorAuthCheckTokenTimeout = 20002 // Token已超时
	ErrorAuthToken             = 20003 // Token生成失败
	ErrorAuth                  = 20004 // Token错误

	ErrorUploadSaveImageFail    = 30001 // 保存图片失败
	ErrorUploadCheckImageFail   = 30002 // 检查图片失败
	ErrorUploadCheckImageFormat = 30003 // 校验图片错误，图片格式或大小有问题

	ErrorJsonMarshalFail = 70001 // 编码Json失败
)

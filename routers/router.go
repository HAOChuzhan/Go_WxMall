package routers

import (
	"wx-gin-master/controller/api/admincont"
	v1 "wx-gin-master/controller/api/v1"
	v2 "wx-gin-master/controller/api/v2"
	_ "wx-gin-master/docs"
	"wx-gin-master/middleware/jwt"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// InitRouter 初始化路由器
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger()) //全局中间件
	r.Use(gin.Recovery())
	/*
		默认使用Logger和Recover中间件。
		Logger是负责进行打印并输出日志的中间件,
		方便开发者进行程序调试;
		Recovery中间件的作如果程序执行过程中遇到panc中断了服务,则 Recovery会恢复程序执行,
		并返回服务器500内误。通常情况下,我们使用默认的gin.Defaul创建 Engine实例。
	*/
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//注册路由组apiv1
	apiv1 := r.Group("/api/v1")
	{
		// User
		userRoute := apiv1.Group("/user")
		{
			userRoute.POST("/register", v1.CreateUser)       // ***
			userRoute.POST("/login", v1.LoginUser)           // ***
			userRoute.GET("/index", jwt.JWT(), v1.ShowUser)  // ***
			userRoute.PUT("/edit", jwt.JWT(), v1.UpdateUser) // ***

			// Favorites
			userRoute.GET("/favorite", jwt.JWT(), v1.FavoritesList)                   // ***
			userRoute.POST("/favorite", jwt.JWT(), v1.FavoritesCreate)                // ***
			userRoute.DELETE("/favorite/product/:id", jwt.JWT(), v1.FavoritesDestroy) // ***

			// Shopping Cart
			userRoute.GET("/shoppingCart", jwt.JWT(), v1.IndexCart)
			userRoute.POST("/shoppingCart", jwt.JWT(), v1.CreateCart)
			userRoute.DELETE("/shoppingCart/:id", jwt.JWT(), v1.DestroyCart)
			userRoute.DELETE("/shoppingCart/:id/number", jwt.JWT(), v1.DecreaseCart)
			/*
				// Addresses
				userRoute.GET("/address", jwt.JWT(), v1.IndexAddress)
				userRoute.POST("/address", jwt.JWT(), v1.CreateAddress)
				userRoute.DELETE("/address/:id", jwt.JWT(), v1.DestroyAddress)
				userRoute.PUT("/address/:id/default", jwt.JWT(), v1.SetDefaultAddress)
				userRoute.PUT("/address/:id", jwt.JWT(), v1.UpdateAddress)
			*/
			// Order
			userRoute.GET("/order", jwt.JWT(), v1.ListOrder)
			userRoute.POST("/order", jwt.JWT(), v1.CreateOrder)
			userRoute.DELETE("/order/:id", jwt.JWT(), v1.DestroyOrder)
			userRoute.POST("/order/:id/pay", jwt.JWT(), v1.PayOrder)
			userRoute.GET("/order/:id", jwt.JWT(), v1.ViewOrderDetailsOrder)

		}
		// Products
		productRoute := apiv1.Group("/product")
		{
			//productRoute.GET("/index", v1.IndexProduct)
			//productRoute.GET("/search", v1.SearchProduct)
			productRoute.POST("/indexproduct", v1.ShowProduct)
		}
		apiv1.POST("/banner/createbanner", v1.CreateBanner)
		apiv1.POST("/banner/updatebanner", v1.UpdateBanner)
		apiv1.POST("/banner/deletebanner", v1.DeleteBanner)
		apiv1.POST("/banner/indexbanner", v1.IndexBanner)

		apiv1.POST("/banner/createradio", v1.CreatRadio)
		apiv1.POST("/banner/updateradio", v1.UpdateRadio)
		apiv1.POST("/banner/deleteradio", v1.DeleteRadio)
		apiv1.POST("/banner/indexradio", v1.IndexRadio)

		apiv1.POST("/banner/createactivity", v1.CreateActivity)
		apiv1.POST("/banner/updateactivity", v1.UpdateActivity)
		apiv1.POST("/banner/deleteactivity", v1.DeleteActivity)
		apiv1.POST("/banner/indexactivity", v1.IndexActivity)
	}

	apiv2 := r.Group("/api/v2")
	{
		v2wxUser := apiv2.Group("/wx/user")
		{
			v2wxUser.GET("/", v2.WxLogin)
			v2wxUser.POST("/", jwt.JWT(), v2.WxGetUserInfo)
			v2wxUser.POST("/phone", jwt.JWT(), v2.WxGetPhone)
		}

		v2wxPay := apiv2.Group("/pay")
		{
			v2wxPay.POST("/topup", jwt.JWT(), v2.WxTopup)
			v2wxPay.Any("/topup/callback", v2.WxTopupCallback)
			v2wxPay.POST("/order", v2.WxPayForOrder)
			v2wxPay.Any("/order/callback", v2.WxPayForOrderCallback)
		}
	}

	// 注册后台路由
	r.POST("/adminUserLogin", admincont.LoginAdmin) // ***
	adminApi := r.Group("/admin")
	adminApi.Use(jwt.Admin())
	{
		adminApi.GET("/info", admincont.ShowAdmin) // ***
		// Products
		adminProductRoute := adminApi.Group("/product")
		{
			adminProductRoute.POST("/createproduct", admincont.CreateProduct) // ***
			adminProductRoute.PUT("/:oid", admincont.UpdateProduct)           // ***
			adminProductRoute.DELETE("/:oid", admincont.DestroyProduct)       // ***
		}

		adminOrderRoute := adminApi.Group("/order")
		{
			adminOrderRoute.POST("/:oid/express", admincont.UpdateOrder)
		}
	}

	return r
}

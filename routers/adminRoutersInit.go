package routers

import (
	"demonwuwen/mallitying/controllers/admin"
	"demonwuwen/mallitying/middlewares"
	"github.com/gin-gonic/gin"
)

func AdminRoutersInit(r *gin.Engine) {
	//middlewares.InitMiddleware中间件

	adminLogRouters := r.Group("/admin")
	{
		adminLogRouters.GET("/login", admin.LoginController{}.Index)
		adminLogRouters.GET("/captcha", admin.LoginController{}.Captcha)
		adminLogRouters.POST("/doLogin", admin.LoginController{}.DoLogin)
	}

	adminRouters := r.Group("/admin", middlewares.InitAdminAuthMiddleware)
	//adminRouters := r.Group("/admin")
	{
		adminRouters.GET("/", admin.MainController{}.Index)
		adminRouters.GET("/welcome", admin.MainController{}.Welcome)
		adminRouters.GET("/changeStatus", admin.MainController{}.ChangeStatus)
		adminRouters.GET("/changeNum", admin.MainController{}.ChangeNum)

		//adminRouters.GET("/login", admin.LoginController{}.Index)
		//adminRouters.GET("/captcha", admin.LoginController{}.Captcha)
		//adminRouters.POST("/doLogin", admin.LoginController{}.DoLogin)
		adminRouters.GET("/loginOut", admin.LoginController{}.LoginOut)

		adminRouters.GET("/manager", admin.ManagerController{}.Index)
		adminRouters.GET("/manager/add", admin.ManagerController{}.Add)
		adminRouters.POST("/manager/doAdd", admin.ManagerController{}.DoAdd)
		adminRouters.GET("/manager/edit", admin.ManagerController{}.Edit)
		adminRouters.POST("/manager/doEdit", admin.ManagerController{}.DoEdit)
		adminRouters.GET("/manager/delete", admin.ManagerController{}.Delete)
		//
		adminRouters.GET("/focus", admin.FocusController{}.Index)
		adminRouters.GET("/focus/add", admin.FocusController{}.Add)
		adminRouters.POST("/focus/doAdd", admin.FocusController{}.DoAdd)
		adminRouters.GET("/focus/edit", admin.FocusController{}.Edit)
		adminRouters.POST("/focus/doEdit", admin.FocusController{}.DoEdit)
		adminRouters.GET("/focus/delete", admin.FocusController{}.Delete)
		//
		adminRouters.GET("/role", admin.RoleController{}.Index)
		adminRouters.GET("/role/add", admin.RoleController{}.Add)
		adminRouters.POST("/role/doAdd", admin.RoleController{}.DoAdd)
		adminRouters.GET("/role/edit", admin.RoleController{}.Edit)
		adminRouters.POST("/role/doEdit", admin.RoleController{}.DoEdit)
		adminRouters.GET("/role/delete", admin.RoleController{}.Delete)
		adminRouters.GET("/role/auth", admin.RoleController{}.Auth)
		adminRouters.POST("/role/doAuth", admin.RoleController{}.DoAuth)

		adminRouters.GET("/access", admin.AccessController{}.Index)
		adminRouters.GET("/access/add", admin.AccessController{}.Add)
		adminRouters.POST("/access/doAdd", admin.AccessController{}.DoAdd)
		adminRouters.GET("/access/edit", admin.AccessController{}.Edit)
		adminRouters.POST("/access/doEdit", admin.AccessController{}.DoEdit)
		adminRouters.GET("/access/delete", admin.AccessController{}.Delete)

		adminRouters.GET("/goodsCate", admin.GoodsCateController{}.Index)
		adminRouters.GET("/goodsCate/add", admin.GoodsCateController{}.Add)
		adminRouters.POST("/goodsCate/doAdd", admin.GoodsCateController{}.DoAdd)
		adminRouters.GET("/goodsCate/edit", admin.GoodsCateController{}.Edit)
		adminRouters.POST("/goodsCate/doEdit", admin.GoodsCateController{}.DoEdit)
		adminRouters.GET("/goodsCate/delete", admin.GoodsCateController{}.Delete)
	}
}

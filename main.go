package main

import (
	"demonwuwen/mallitying/routers"
	"demonwuwen/mallitying/utils"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"html/template"

	_ "net/http/pprof"
)

func main() {
	//创建默认路由引擎
	r := gin.Default()
	//	自定义模版函数
	r.SetFuncMap(template.FuncMap{
		//"UnixToTime":utils.UnixToTime(), 不能要括号
		"UnixToTime": utils.UnixToTime,
	})

	//	加载模版，放在配置路由前
	r.LoadHTMLGlob("templates/**/**/*")
	//配置静态web目录   第一个参数表示路由, 第二个参数表示映射的目录
	r.Static("/static", "./static")

	//	创建给予cookie的存储引擎，
	store := cookie.NewStore([]byte("cookieSecret"))
	//	配置session中间价，存储引擎尾store，
	r.Use(sessions.Sessions("mysession", store))

	routers.AdminRoutersInit(r)

	r.Run(":8080")
}

package middlewares

import (
	"demonwuwen/mallitying/models"
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func InitAdminAuthMiddleware(c *gin.Context) {
	//进行权限判断 没有登录的用户 不能进入后台管理中心

	//1、获取Url访问的地址  /admin/captcha
	//2、获取Session里面保存的用户信息

	//3、判断Session中的用户信息是否存在，如果不存在跳转到登录页面（注意需要判断） 如果存在继续向下执行

	//4、如果Session不存在，判断当前访问的URl是否是login doLogin captcha，如果不是跳转到登录页面，如果是不行任何操作

	//	1、获取Url访问地址  /admin/captcha?=t
	pathname := strings.Split(c.Request.URL.String(), "?")[0]
	// 2、获取Session里面保存的用户信息
	session := sessions.Default(c)
	userinfo := session.Get("userinfo")
	userinfoStr, ok := userinfo.(string)
	if ok {
		var userinfoStruct []models.Manager
		err := json.Unmarshal([]byte(userinfoStr), &userinfoStruct)
		if err != nil || !(len(userinfoStruct) > 0) && userinfoStruct[0].Username != "" {
			if pathname != "/admin/login" && pathname != "/admin/doLogin" && pathname != "/admin/captcha" {
				c.Redirect(http.StatusFound, "/admin/login")
			}
		} else {
			//用户登陆成功 权限判断
		}
	} else {
		if pathname != "/admin/login" && pathname != "/admin/doLogin" && pathname != "admin/captcha" {
			c.Redirect(http.StatusFound, "/admin/login")
		}
	}

}

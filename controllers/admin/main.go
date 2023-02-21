package admin

import (
	"demonwuwen/mallitying/models"
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type MainController struct {
}

func (con MainController) Index(c *gin.Context) {
	//获取userinfo 对应的session
	session := sessions.Default(c)
	userinfo := session.Get("userinfo")
	//类型断言 来判断 userinfo是不是一个string
	userinfoStr, ok := userinfo.(string)

	if ok {
		//获取用户信息
		var userinfoStruct []models.Manager
		json.Unmarshal([]byte(userinfoStr), &userinfoStruct)
		fmt.Println(userinfoStruct)
		//获取所有权限
		accessList := []models.Access{}
		//排序，按模块顺序排序
		//models.DB.Where("module_id = ?", 0).Preload("AccessItem").Find(&accessList)
		models.DB.Where("module_id = ?", 0).Preload("AccessItem", func(db *gorm.DB) *gorm.DB {
			return db.Order("access.sort DESC")
		}).Order("sort DESC").Find(&accessList)

		//获取当前角色的权限，并把权限id放在map里
		roleAccess := []models.RoleAccess{}
		models.DB.Where("role_id = ?", userinfoStruct[0].RoleId).Find(&roleAccess)
		roleAccessMap := make(map[int]int)
		for _, v := range roleAccess {
			roleAccessMap[v.AccessId] = v.AccessId
		}

		//遍历所有权限数据，判断当前权限的id是否在角色权限的map对象中，
		for i := 0; i < len(accessList); i++ {
			if _, ok := roleAccessMap[accessList[i].Id]; ok {
				accessList[i].Checked = true
			}
			for j := 0; j < len(accessList[i].AccessItem); j++ {
				if _, ok := roleAccessMap[accessList[i].AccessItem[j].Id]; ok {
					accessList[i].AccessItem[j].Checked = true
				}
			}
		}

		c.HTML(http.StatusOK, "admin/main/index.html", gin.H{
			"username":   userinfoStruct[0].Username,
			"accessList": accessList,
			"isSuper":    userinfoStruct[0].IsSuper,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"username": "session不存在",
		})
	}

}

func (con MainController) Welcome(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/main/welcome.html", gin.H{})
}

// ChangeStatus 公共修改状态
func (con MainController) ChangeStatus(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "参数传入错误",
		})
		return
	}
	table := c.Query("table")
	field := c.Query("field")

	err = models.DB.Exec("update "+table+" set "+field+" = ABS( "+field+" -1) where id = ?", id).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "修改失败请重试",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "修改数据成功",
	})
	//c.HTML(http.StatusOK, "admin/main/welcome.html", gin.H{})
}

func (con MainController) ChangeNum(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "参数传入错误",
		})
		return
	}
	table := c.Query("table")
	field := c.Query("field")
	num := c.Query("num")

	err = models.DB.Exec("update "+table+" set "+field+"="+num+" where id=?", id).Error
	fmt.Println("err = ", err)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "修改失败请重试",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "修改数据成功",
	})
	//c.HTML(http.StatusOK, "admin/main/welcome.html", gin.H{})
}

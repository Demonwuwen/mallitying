package admin

import (
	"demonwuwen/mallitying/models"
	"demonwuwen/mallitying/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type ManagerController struct {
	BaseController
}

func (con ManagerController) Index(c *gin.Context) {
	managerList := []models.Manager{}
	models.DB.Preload("Role").Find(&managerList)

	c.HTML(http.StatusOK, "admin/manager/index.html", gin.H{
		"managerList": managerList,
	})
}

func (con ManagerController) Add(c *gin.Context) {
	roleList := []models.Role{}
	models.DB.Find(&roleList)
	c.HTML(http.StatusOK, "admin/manager/add.html", gin.H{
		"roleList": roleList,
	})
}

func (con ManagerController) DoAdd(c *gin.Context) {
	roleId, err := strconv.Atoi(c.PostForm("role_id"))
	if err != nil {
		con.Error(c, "参数传入错误", "/admin/manager/add")
	}
	username := strings.Trim(c.PostForm("username"), " ")
	password := strings.Trim(c.PostForm("password"), " ")
	email := strings.Trim(c.PostForm("email"), " ")
	mobile := strings.Trim(c.PostForm("mobile"), " ")
	if len(username) < 2 || len(password) < 6 {
		con.Error(c, "用户名或密码长度不符", "/admin/manager/add")
	}
	managerList := []models.Manager{}
	models.DB.Where("username=?", username).Find(&managerList)
	if len(managerList) > 0 {
		con.Error(c, "管理员已存在", "/admin/manager/add")
		return
	}

	manager := models.Manager{
		Username: username,
		Password: utils.Md5(password),
		Mobile:   mobile,
		Email:    email,
		Status:   1,
		RoleId:   roleId,
		AddTime:  int(utils.GetUnix()),
	}
	err1 := models.DB.Create(&manager).Error
	if err1 != nil {
		con.Error(c, "管理员增加失败", "/admin/manager/add")
		return
	}
	con.Success(c, "管理员添加成功", "/admin/manager")

}

func (con ManagerController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		con.Error(c, "传入参数错误", "/admin/manager")
		return
	}
	manager := models.Manager{Id: id}
	err = models.DB.Delete(&manager).Error
	if err != nil {
		con.Error(c, "删除失败", "/admin/manager")
		return
	}
	con.Success(c, "删除成功", "/admin/manager")
}

func (con ManagerController) Edit(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		con.Error(c, "传入参数错误", "/admin/manager")
		return
	}
	manager := models.Manager{Id: id}
	models.DB.Find(&manager)

	roleList := []models.Role{}
	models.DB.Find(&roleList)
	c.HTML(http.StatusOK, "admin/manager/edit.html", gin.H{
		"manager":  manager,
		"roleList": roleList,
	})
}

func (con ManagerController) DoEdit(c *gin.Context) {
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		con.Error(c, "传入参数错误", "/admin/manager")
		return
	}
	roleId, err1 := strconv.Atoi(c.PostForm("role_id"))
	if err1 != nil {
		con.Error(c, "传入参数错误", "/admin/manager")
		return
	}
	username := strings.Trim(c.PostForm("username"), " ")
	password := strings.Trim(c.PostForm("password"), " ")
	email := strings.Trim(c.PostForm("email"), " ")
	mobile := strings.Trim(c.PostForm("mobile"), " ")
	if len(mobile) > 11 {
		con.Error(c, "mobile length illegal ", "/admin/manager")
		return
	}
	manager := models.Manager{Id: id}
	models.DB.Find(&manager)
	manager.Username = username
	//manager.Password = password
	manager.Mobile = mobile
	manager.Email = email
	manager.RoleId = roleId
	if password != "" {
		if len(password) < 6 {
			con.Error(c, "密码长度不合法，不小于6位数", "/admin/manager/edit?id="+strconv.Itoa(id))
			return
		}
		manager.Password = utils.Md5(password)
	}

	err2 := models.DB.Save(&manager).Error
	if err2 != nil {
		con.Error(c, "修改数据失败", "/admin/manager/edit?id="+strconv.Itoa(id))
		return
	}
	con.Success(c, "数据修改成功", "/admin/manager")

}

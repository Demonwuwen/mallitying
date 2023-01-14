package admin

import (
	"demonwuwen/mallitying/models"
	"demonwuwen/mallitying/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type RoleController struct {
	BaseController
}

func (con RoleController) Index(c *gin.Context) {
	roleList := []models.Role{}
	models.DB.Find(&roleList)

	c.HTML(http.StatusOK, "admin/role/index.html", gin.H{
		"roleList": roleList,
	})
}

func (con RoleController) Add(c *gin.Context) {
	roleList := []models.Role{}
	models.DB.Find(&roleList)
	c.HTML(http.StatusOK, "admin/role/add.html", gin.H{
		"roleList": roleList,
	})
}

func (con RoleController) DoAdd(c *gin.Context) {
	title := strings.Trim(c.PostForm("title"), " ")

	description := strings.Trim(c.PostForm("description"), " ")

	role := models.Role{}
	role.Title = title
	role.Description = description
	role.AddTime = int(utils.GetUnix())

	err1 := models.DB.Create(&role).Error
	if err1 != nil {
		con.Error(c, "角色增加失败", "/admin/role/add")
		return
	}
	con.Success(c, "角色添加成功", "/admin/role")

}

func (con RoleController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		con.Error(c, "传入参数错误", "/admin/role")
		return
	}
	role := models.Role{Id: id}
	err = models.DB.Delete(&role).Error
	if err != nil {
		con.Error(c, "删除失败", "/admin/role")
		return
	}
	con.Success(c, "删除成功", "/admin/role")
}

func (con RoleController) Edit(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		con.Error(c, "传入参数错误", "/admin/role")
		return
	}
	role := models.Role{Id: id}
	models.DB.Find(&role)

	roleList := []models.Role{}
	models.DB.Find(&roleList)
	c.HTML(http.StatusOK, "admin/role/edit.html", gin.H{
		"role":     role,
		"roleList": roleList,
	})
}

func (con RoleController) DoEdit(c *gin.Context) {
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		con.Error(c, "传入参数错误", "/admin/role")
		return
	}
	title := strings.Trim(c.PostForm("title"), " ")
	if title == "" {
		con.Error(c, "角色标题不能为空", "/admin/role/edit?id="+strconv.Itoa(id))
		return
	}
	description := strings.Trim(c.PostForm("description"), " ")

	role := models.Role{}
	models.DB.Find(&role)
	role.Title = title
	role.Description = description
	err2 := models.DB.Save(&role).Error
	if err2 != nil {
		con.Error(c, "修改数据失败", "/admin/role/edit?id="+strconv.Itoa(id))
		return
	}
	con.Success(c, "数据修改成功", "/admin/role")

}

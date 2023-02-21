package admin

import (
	"demonwuwen/mallitying/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type AccessController struct {
	BaseController
}

func (con AccessController) Index(c *gin.Context) {
	accessList := []models.Access{}
	//models.DB.Find(&accessList)
	models.DB.Where("module_id=?", 0).Preload("AccessItem").Find(&accessList)

	c.HTML(http.StatusOK, "admin/access/index.html", gin.H{
		"accessList": accessList,
	})
}

func (con AccessController) Add(c *gin.Context) {
	accessList := []models.Access{}
	models.DB.Where("module_id=?", 0).Find(&accessList)
	c.HTML(http.StatusOK, "admin/access/add.html", gin.H{
		"accessList": accessList,
	})
}

func (con AccessController) DoAdd(c *gin.Context) {
	moduleName := strings.Trim(c.PostForm("module_name"), " ")
	accessType, err0 := strconv.Atoi(c.PostForm("type"))
	actionName := strings.Trim(c.PostForm("action_name"), " ")
	url := strings.Trim(c.PostForm("url"), " ")
	moduleId, err3 := strconv.Atoi(c.PostForm("module_id"))
	sort, err1 := strconv.Atoi(c.PostForm("sort"))
	description := c.PostForm("description")
	status, err2 := strconv.Atoi(c.PostForm("status"))
	if err0 != nil || err1 != nil || err2 != nil || err3 != nil {
		con.Error(c, "传入参数错误", "/admin/access/add")
		return
	}
	if moduleName == "" {
		con.Error(c, "模块名称不能为空", "/admin/access/add")
		return
	}
	access := models.Access{
		ModuleName:  moduleName,
		ActionName:  actionName,
		Type:        accessType,
		Url:         url,
		ModuleId:    moduleId,
		Sort:        sort,
		Description: description,
		Status:      status,
	}
	err4 := models.DB.Create(&access).Error
	if err4 != nil {
		con.Error(c, "添加权限数据失败", "/admin/access/add")
		return
	}
	con.Success(c, "添加数据成功", "/admin/access")
}

func (con AccessController) Edit(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		con.Error(c, "数据传输错误", "/admin/access")
		return
	}
	access := models.Access{Id: id}
	models.DB.Find(&access)
	c.HTML(http.StatusOK, "admin/access/edit.html", gin.H{
		"access": access,
	})

}
func (con AccessController) DoEdit(c *gin.Context) {
	id, err1 := strconv.Atoi(c.PostForm("id"))
	//if err1 != nil {
	//	fmt.Println("err1 =", err1)
	//}
	moduleName := strings.Trim(c.PostForm("module_name"), " ")
	accessType, err2 := strconv.Atoi(c.PostForm("type"))
	//if err2 != nil {
	//	fmt.Println("err2 =", err2)
	//}
	actionName := strings.Trim(c.PostForm("action_name"), " ")
	url := strings.Trim(c.PostForm("url"), " ")
	moduleId, err3 := strconv.Atoi(c.PostForm("module_id"))
	//if err3 != nil {
	//	fmt.Println("err3 =", err1)
	//}
	sort, err4 := strconv.Atoi(c.PostForm("sort"))
	//if err4 != nil {
	//	fmt.Println("err4 =", err4)
	//}
	description := strings.Trim(c.PostForm("description"), " ")
	status, err5 := strconv.Atoi(c.PostForm("status"))
	//if err5 != nil {
	//	fmt.Println("err5 =", err5)
	//}
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil {
		con.Error(c, "传输ID数据错误", "/admin/access/edit?id="+strconv.Itoa(id))
		return
	}

	if moduleName == "" {
		con.Error(c, "模块名称不能为空", "/admin/access/edit?id="+strconv.Itoa(id))
		return
	}
	if accessType != 1 && actionName == "" {
		con.Error(c, "操作名称不能为空", "/admin/access/edit?id="+strconv.Itoa(id))
		return
	}

	access := models.Access{Id: id}
	models.DB.Find(&access)

	access.ModuleName = moduleName
	access.Type = accessType
	access.ActionName = actionName
	access.Url = url
	access.ModuleId = moduleId
	access.Sort = sort
	access.Description = description
	access.Status = status
	err6 := models.DB.Save(&access).Error
	if err6 != nil {
		con.Error(c, "数据修改失败", "/admin/access/edit")
		return
	}
	con.Success(c, "数据更新成功", "/admin/access")
}
func (con AccessController) Delete(c *gin.Context) {
	//c.String(http.StatusOK, "-delete-")
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/access")
	} else {
		//获取我们要删除的数据
		access := models.Access{Id: id}
		models.DB.Find(&access)
		if access.ModuleId == 0 { //顶级模块
			accessList := []models.Access{}
			models.DB.Where("module_id = ?", access.Id).Find(&accessList)
			if len(accessList) > 0 {
				con.Error(c, "当前模块下面有菜单或者操作，请删除菜单或者操作以后再来删除这个数据", "/admin/access")
			} else {
				models.DB.Delete(&access)
				con.Success(c, "删除数据成功", "/admin/access")
			}
		} else { //操作 或者菜单
			models.DB.Delete(&access)
			con.Success(c, "删除数据成功", "/admin/access")
		}

	}
}

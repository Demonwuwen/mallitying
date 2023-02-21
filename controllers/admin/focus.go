package admin

import (
	"demonwuwen/mallitying/models"
	"demonwuwen/mallitying/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type FocusController struct {
	BaseController
}

func (con FocusController) Index(c *gin.Context) {
	focusList := []models.Focus{}
	models.DB.Find(&focusList)
	c.HTML(http.StatusOK, "admin/focus/index.html", gin.H{
		"focusList": focusList,
	})
}

func (con FocusController) Add(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/focus/add.html", gin.H{})
}
func (con FocusController) DoAdd(c *gin.Context) {
	title := c.PostForm("title")
	fmt.Println("title = ", title)
	focusType, err1 := strconv.Atoi(c.PostForm("focus_type"))
	link := c.PostForm("link")
	fmt.Println("link = ", link)

	sort, err2 := strconv.Atoi(c.PostForm("sort"))
	status, err3 := strconv.Atoi(c.PostForm("status"))
	fmt.Println("sort = ", sort)
	fmt.Println("status = ", status)

	if err1 != nil || err3 != nil {
		con.Error(c, "非法请求", "/admin/focus/add")
	}
	if err2 != nil {
		con.Error(c, "请输入正确的排序值", "/admin/focus/add")
	}
	//上传文件
	focusImgSrc, err4 := utils.UploadImg(c, "focus_img")
	if err4 != nil {
		fmt.Println(err4)
	}
	fmt.Println(" 上传文件")
	fmt.Println(" 上传文件")

	focus := models.Focus{
		Title:     title,
		FocusType: focusType,
		FocusImg:  focusImgSrc,
		Link:      link,
		Sort:      sort,
		Status:    status,
		AddTime:   int(utils.GetUnix()),
	}
	err5 := models.DB.Create(&focus).Error
	if err5 != nil {
		con.Error(c, "增加轮播图失败", "/admin/focus/add")
	} else {
		con.Success(c, "增加轮播图成功", "/admin/focus")
	}

}

func (con FocusController) Edit(c *gin.Context) {
	id, err1 := strconv.Atoi(c.Query("id"))
	if err1 != nil {
		con.Error(c, "传入参数错误", "/admin/focus")
		return
	}
	focus := models.Focus{Id: id}
	models.DB.Find(&focus)
	fmt.Println("focus = ", focus)
	c.HTML(http.StatusOK, "admin/focus/edit.html", gin.H{
		"focus": focus,
	})
}
func (con FocusController) DoEdit(c *gin.Context) {
	id, err1 := strconv.Atoi(c.PostForm("id"))
	title := c.PostForm("title")
	focusType, err2 := strconv.Atoi(c.PostForm("focus_type"))
	link := c.PostForm("link")
	sort, err3 := strconv.Atoi(c.PostForm("sort"))
	status, err4 := strconv.Atoi(c.PostForm("status"))

	if err1 != nil || err2 != nil || err4 != nil {
		con.Error(c, "非法请求", "/admin/focus")
	}
	if err3 != nil {
		con.Error(c, "请输入正确的排序值", "/admin/focus/edit?id="+strconv.Itoa(id))
	}
	//上传文件
	focusImg, _ := utils.UploadImg(c, "focus_img")

	focus := models.Focus{Id: id}
	models.DB.Find(&focus)
	focus.Title = title
	focus.FocusType = focusType
	focus.Link = link
	focus.Sort = sort
	focus.Status = status
	if focusImg != "" {
		focus.FocusImg = focusImg
	}
	err5 := models.DB.Save(&focus).Error
	if err5 != nil {
		con.Error(c, "修改数据失败请重新尝试", "/admin/focus/edit?id="+strconv.Itoa(id))
	} else {
		con.Success(c, "增加轮播图成功", "/admin/focus")
	}
	//c.HTML(http.StatusOK, "admin/focus/add.html", gin.H{})
}

func (con FocusController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		fmt.Println("ID 传参失败")
		con.Error(c, "ID 参数传递失败", "/admin/focus")
		return
	}
	focus := &models.Focus{Id: id}
	err = models.DB.Where("id = ?", id).Delete(focus).Error
	if err != nil {
		con.Error(c, "轮播图删除失败", "/admin/focus")
		return
	}
	con.Success(c, "轮播图删除成功", "/admin/focus")
}

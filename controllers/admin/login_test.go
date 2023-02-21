package admin

import (
	"demonwuwen/mallitying/models"
	"fmt"
	"testing"
)

func TestLogin(t *testing.T) {
	userinf := models.Manager{}
	username := "admin"
	//password := "123456"
	models.DB.Where("username = ?", username).Find(&userinf)

	fmt.Println(userinf)

}

package factory

import (
	"fmt"
	"github.com/huaanhuang/design-pattern-go/factory/models"
	"testing"
)

func TestCreateUser(t *testing.T) {
	admin := CreateUser(AdminUser)(1, "zhangsan").(*models.Admin)
	fmt.Println(admin)

	user := CreateUser(FrontUser)(2, "lisi").(*models.User)
	fmt.Println(user)
}

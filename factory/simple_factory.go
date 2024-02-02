package factory

import "github.com/huaanhuang/design-pattern-go/factory/models"

const (
	FrontUser UserType = iota
	AdminUser
)

type UserType int8

func CreateUser(t UserType) models.CreateUserFunc {
	switch t {
	case AdminUser:
		return models.NewAdmin()
	case FrontUser:
		return models.NewUser()
	default:
		return models.NewUser()
	}
}

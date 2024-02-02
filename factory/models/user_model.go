package models

type CreateUserFunc func(id int, name string) interface{}

type User struct {
	ID   int
	Name string
}

func NewUser() CreateUserFunc {
	return func(id int, name string) interface{} {
		return &User{
			ID:   id,
			Name: name,
		}
	}
}

type Admin struct {
	ID   int
	Name string
	Role string
}

func NewAdmin() CreateUserFunc {
	return func(id int, name string) interface{} {
		return &Admin{
			ID:   id,
			Name: name,
			Role: "admin",
		}
	}
}

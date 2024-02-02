package models

type IProduct interface {
	GetInfo() string
}

type Book struct {
}

func (obj *Book) GetInfo() string {
	return "书籍"
}

type Panties struct {
}

func (obj *Panties) GetInfo() string {
	return "火炮儿"
}

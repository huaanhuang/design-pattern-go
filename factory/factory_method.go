package factory

import "github.com/huaanhuang/design-pattern-go/factory/models"

const (
	BookType ProductType = iota
	PantiesType
)

type ProductType int8

// IProductFactory 总工厂
type IProductFactory interface {
	CreateProduct(t ProductType) models.IProduct
}

// KnowledgeFactory 知识类工厂
type KnowledgeFactory struct {
}

func (obj *KnowledgeFactory) CreateProduct(t ProductType) models.IProduct {
	switch t {
	case BookType:
		return &models.Book{}
	default:
		return &models.Book{}
	}
}

// GeneralMerchandiseFactory 百货类工厂
type GeneralMerchandiseFactory struct {
}

func (obj *GeneralMerchandiseFactory) CreateProduct(t ProductType) models.IProduct {
	switch t {
	case PantiesType:
		return &models.Panties{}
	default:
		return &models.Panties{}
	}
}

package factory

import (
	"fmt"
	"testing"
)

func TestFactoryFactory(t *testing.T) {
	// 测试知识类工厂
	fmt.Println(new(KnowledgeFactory).CreateProduct(BookType).GetInfo())

	// 测试百货类工厂
	fmt.Println(new(GeneralMerchandiseFactory).CreateProduct(PantiesType).GetInfo())
}

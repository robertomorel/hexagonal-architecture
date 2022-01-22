package dto

import "github.com/codeedu/go-hexagonal/application"

// Objeto anêmico DTO
type Product struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Status string  `json:"status"`
}

// Função para criar um novo objeto
func NewProduct() *Product {
	return &Product{}
}

/*
	O Bind será algo onde podemos pegar os dados desse DTO e grudar com os dados da entidade product
*/
func (p *Product) Bind(product *application.Product) (*application.Product, error) {
	if p.ID != "" {
		product.ID = p.ID
	}
	product.Name = p.Name
	product.Price = p.Price
	product.Status = p.Status
	_, err := product.IsValid()
	if err != nil {
		return &application.Product{}, err
	}
	return product, nil
}

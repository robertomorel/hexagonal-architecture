// Entidade Produto
package application

import (
	"errors"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

// Interface dos métodos da entidade Produto
type ProductInterface interface {
	// Método | Retorno
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetID() string
	GetName() string
	GetStatus() string
	GetPrice() float64
}

// Interface do Service, que irá interagir com o DB
type ProductServiceInterface interface {
	// Método | Retorno
	Get(id string) (ProductInterface, error)
	Create(name string, price float64) (ProductInterface, error)
	Enable(product ProductInterface) (ProductInterface, error)
	Disable(product ProductInterface) (ProductInterface, error)
}

// ---- Interface da Leitura e gravação -------
/*
	Aplicação do Interface Segregation Principle [SOLID]
*/
type ProductReader interface {
	Get(id string) (ProductInterface, error)
}

type ProductWriter interface {
	Save(product ProductInterface) (ProductInterface, error)
}

// Composição para unir as duas interfaces
type ProductPersistenceInterface interface {
	ProductReader
	ProductWriter
}

// ----------------------------------------------

const (
	DISABLED = "disabled"
	ENABLED  = "enabled"
)

// Struct para definir os campos da classe com as tags para validação
type Product struct {
	ID     string  `valid:"uuidv4"`
	Name   string  `valid:"required"`
	Price  float64 `valid:"float,optional"`
	Status string  `valid:"required"`
}

// Esta função estará criando uma nova struct com a convenção de preenchimento de alguns campos
// Estará retornando um ponteiro para Product (endereço de memória)
func NewProduct() *Product {
	product := Product{
		ID:     uuid.NewV4().String(),
		Status: DISABLED,
	}
	// O "&" deve ser usado obrigatoriamente com retornos de ponteiros para enviar a localização na memória deste elemento
	return &product
}

// Garante que nada na entidade esteja errado
func (p *Product) IsValid() (bool, error) {
	if p.Status == "" {
		p.Status = DISABLED
	}

	if p.Status != ENABLED && p.Status != DISABLED {
		return false, errors.New("the status must be enabled or disabled")
	}

	if p.Price < 0 {
		return false, errors.New("the price must be greater or equal zero")
	}

	// O "govalidator" valida o status e valor de cada uma das propriedades segundo o que foi definido na struct
	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (p *Product) Enable() error {
	// Para ativar o produto o p.Price > 0
	if p.Price > 0 {
		p.Status = ENABLED
		return nil
	}
	return errors.New("the price must be greater than zero to enable the product")
}

func (p *Product) Disable() error {
	if p.Price == 0 {
		p.Status = DISABLED
		return nil
	}
	return errors.New("the price must be zero in order to have the product disabled")
}

func (p *Product) GetID() string {
	return p.ID
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetStatus() string {
	return p.Status
}

func (p *Product) GetPrice() float64 {
	return p.Price
}

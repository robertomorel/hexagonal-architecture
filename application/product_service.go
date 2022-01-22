package application

/*
	Este service vai fazer exatamente o que as interfaces implementadas na entidade Product manda
	É um middleware entre a entidade e as ações no DB, por exemplo
*/

// Iniciando o struct para este service
type ProductService struct {

	// Aplicando o Dependency Inversion Principle [SOLID]
	// Vou passar o tipo de persistência que o service deve usar
	// O Service não precisa saber qual o DB é usado ou se iremos persistir localmente
	Persistence ProductPersistenceInterface
}

func NewProductService(persistence ProductPersistenceInterface) *ProductService {
	return &ProductService{Persistence: persistence}
}

// Implementando o Get da interface
func (s *ProductService) Get(id string) (ProductInterface, error) {
	product, err := s.Persistence.Get(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

// Implementando o Create da interface
func (s *ProductService) Create(name string, price float64) (ProductInterface, error) {
	product := NewProduct()
	product.Name = name
	product.Price = price
	_, err := product.IsValid()
	if err != nil {
		return &Product{}, err
	}
	result, err := s.Persistence.Save(product)
	if err != nil {
		return &Product{}, err
	}
	return result, nil
}

func (s *ProductService) Enable(product ProductInterface) (ProductInterface, error) {
	err := product.Enable()
	if err != nil {
		return &Product{}, err
	}
	result, err := s.Persistence.Save(product)
	if err != nil {
		return &Product{}, err
	}
	return result, nil
}

func (s *ProductService) Disable(product ProductInterface) (ProductInterface, error) {
	err := product.Disable()
	if err != nil {
		return &Product{}, err
	}
	result, err := s.Persistence.Save(product)
	if err != nil {
		return &Product{}, err
	}
	return result, nil
}

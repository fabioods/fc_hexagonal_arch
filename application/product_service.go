package application

type ProductService struct {
	Persistence ProductPersistenceInterface
}

func NewProductService(persistence ProductPersistenceInterface) *ProductService {
	return &ProductService{
		Persistence: persistence,
	}
}

func (s *ProductService) Get(id string) (ProductInterface, error) {
	product, err := s.Persistence.Get(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *ProductService) Create(name string, price float64) (ProductInterface, error) {
	product := NewProduct()
	product.Name = name
	product.Price = price

	if ok, err := product.IsValid(); !ok {
		return nil, err
	}

	return s.Persistence.Save(product)
}

func (s *ProductService) Enable(product ProductInterface) (ProductInterface, error) {
	err := product.Enable()
	if err != nil {
		return nil, err
	}

	return s.Persistence.Save(product)
}

func (s *ProductService) Disable(product ProductInterface) (ProductInterface, error) {
	err := product.Disable()
	if err != nil {
		return nil, err
	}

	return s.Persistence.Save(product)
}

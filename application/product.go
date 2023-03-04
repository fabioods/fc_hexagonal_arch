package application

import (
	"errors"
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetID() string
	GetName() string
	GetStatus() string
	GetPrice() float64
	UpdateName(name string) error
	UpdatePrice(price float64) error
}

type ProductServiceInterface interface {
	Get(id string) (ProductInterface, error)
	Create(name string, price float64) (ProductInterface, error)
	Enable(product ProductInterface) (ProductInterface, error)
	Disable(product ProductInterface) (ProductInterface, error)
}

type ProductReader interface {
	Get(id string) (ProductInterface, error)
}

type ProductWriter interface {
	Save(product ProductInterface) (ProductInterface, error)
}

type ProductPersistenceInterface interface {
	ProductReader
	ProductWriter
}

const (
	DISABLED = "disabled"
	ENABLED  = "enabled"
)

type Product struct {
	ID     string  `json:"id" valid:"uuidv4"`
	Name   string  `json:"name" valid:"required"`
	Status string  `json:"status" valid:"required"`
	Price  float64 `json:"price" valid:"float,optional"`
}

func NewProduct() *Product {
	return &Product{
		ID:     uuid.NewV4().String(),
		Status: ENABLED,
	}
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func (p *Product) Enable() error {
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
	return errors.New("the price must be lower or equal to zero to disable the product")
}

func (p *Product) IsValid() (bool, error) {
	if p.Status == "" {
		p.Status = DISABLED
	}

	if p.Status != ENABLED && p.Status != DISABLED {
		return false, errors.New("status is invalid")
	}

	if p.Price < 0 {
		return false, errors.New("the price must be greater than zero")
	}

	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return false, err
	}

	return true, nil
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

func (p *Product) UpdateName(name string) error {
	if name == "" {
		return errors.New("name is required")
	}
	p.Name = name
	return nil
}

func (p *Product) UpdatePrice(price float64) error {
	if price < 0 {
		return errors.New("the price must be greater than zero")
	}
	p.Price = price
	return nil
}

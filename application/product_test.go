package application_test

import (
	"testing"

	"github.com/fabioods/fc_hexagonal_arch/application"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Product 1"
	product.Price = 10
	product.Status = application.DISABLED

	err := product.Enable()
	require.Nil(t, err)
	require.Equal(t, application.ENABLED, product.Status)
}

func TestProduct_EnableWithError(t *testing.T) {
	product := application.Product{}
	product.Name = "Product 1"
	product.Price = 0
	product.Status = application.DISABLED

	err := product.Enable()
	require.NotNil(t, err)
	require.Equal(t, application.DISABLED, product.Status)
	require.Equal(t, "the price must be greater than zero to enable the product", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.Name = "Product 1"
	product.Price = 0
	product.Status = application.ENABLED

	err := product.Disable()
	require.Nil(t, err)
	require.Equal(t, application.DISABLED, product.Status)
}

func TestProduct_DisableWithEror(t *testing.T) {
	product := application.Product{}
	product.Name = "Product 1"
	product.Price = 10
	product.Status = application.ENABLED

	err := product.Disable()
	require.NotNil(t, err)
	require.Equal(t, application.ENABLED, product.Status)
	require.Equal(t, "the price must be greater than zero to disable the product", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.Name = "Product 1"
	product.Price = 10
	product.Status = application.DISABLED
	product.ID = uuid.NewV4().String()

	_, err := product.IsValid()
	require.Nil(t, err)
}

func TestProduct_IsValidWithInvalidStatus(t *testing.T) {
	product := application.Product{}
	product.Name = "Product 1"
	product.Price = 10
	product.Status = "invalid"
	product.ID = uuid.NewV4().String()

	_, err := product.IsValid()
	require.NotNil(t, err)
	require.Equal(t, "status is invalid", err.Error())
}

func TestProduct_IsValidWithPriceLessThanZero(t *testing.T) {
	product := application.Product{}
	product.Name = "Product 1"
	product.Price = -10
	product.Status = application.DISABLED
	product.ID = uuid.NewV4().String()

	_, err := product.IsValid()
	require.NotNil(t, err)
	require.Equal(t, "the price must be greater than zero", err.Error())
}

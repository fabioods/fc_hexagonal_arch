package cli

import (
	"fmt"
	"github.com/fabioods/fc_hexagonal_arch/application"
)

func Run(service application.ProductServiceInterface, action string, productID string, productName string, productPrice float64) (string, error) {
	result := ""

	switch action {
	case "create":
		product, err := service.Create(productName, productPrice)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product %s created with success!", product.GetID())
	case "enable":
		product, err := service.Get(productID)
		if err != nil {
			return result, err
		}
		res, err := service.Enable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product %s enabled with success!", res.GetID())
	case "disable":
		product, err := service.Get(productID)
		if err != nil {
			return result, err
		}
		res, err := service.Disable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product %s disabled with success!", res.GetID())
	default:
		product, err := service.Get(productID)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Action %s not found, Product with ID %v", action, product.GetID())
	}
	return result, nil
}

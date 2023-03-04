package application_test

import (
	"github.com/fabioods/fc_hexagonal_arch/application"
	mock_application "github.com/fabioods/fc_hexagonal_arch/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProductService_Get(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	product := mock_application.NewMockProductInterface(controller)
	persistence := mock_application.NewMockProductPersistenceInterface(controller)
	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.NewProductService(persistence)
	_, err := service.Get("123")
	require.Nil(t, err)
}

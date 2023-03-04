package cli_test

import (
	"github.com/fabioods/fc_hexagonal_arch/adapters/cli"
	mock_application "github.com/fabioods/fc_hexagonal_arch/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRun(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	productName := "Product 1"
	productPrice := 10.0
	productStatus := "enabled"
	productID := "1"

	productMock := mock_application.NewMockProductInterface(controller)
	productMock.EXPECT().GetID().Return(productID).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()

	productServiceMock := mock_application.NewMockProductServiceInterface(controller)
	productServiceMock.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	productServiceMock.EXPECT().Get(productID).Return(productMock, nil).AnyTimes()
	productServiceMock.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	productServiceMock.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	resultExpected := "Product 1 created with success!"
	result, err := cli.Run(productServiceMock, "create", "", productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = "Product 1 enabled with success!"
	result, err = cli.Run(productServiceMock, "enable", productID, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = "Product 1 disabled with success!"
	result, err = cli.Run(productServiceMock, "disable", productID, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = "Action notfound not found, Product with ID 1"
	result, err = cli.Run(productServiceMock, "notfound", productID, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

}

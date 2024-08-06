package application_test

import (
	"testing"

	mock_application "github.com/EvyOliveira/fullCycle/application/mocks"
	application "github.com/EvyOliveira/fullCycle/application/service"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestProductService_Get(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	product := mock_application.NewMockProductInterface(controller)
	persistence := mock_application.NewMockProductPersistanceInterface(controller)
	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()
	service := application.ProductService{
		Persistence: persistence,
	}

	result, err := service.Get("abc")
	require.Nil(t, err)
	require.Equal(t, product, result)
}

func TestProductService_Create(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	product := mock_application.NewMockProductInterface(controller)
	persistence := mock_application.NewMockProductPersistanceInterface(controller)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Persistence: persistence,
	}
	result, err := service.Create("Product 1", 10)
	require.Nil(t, err)
	require.Equal(t, product, result)

}

func TestProductService_EnableDisable(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	product := mock_application.NewMockProductInterface(controller)
	product.EXPECT().Enable().Return(nil)
	product.EXPECT().Disable().Return(nil)
	persistence := mock_application.NewMockProductPersistanceInterface(controller)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Persistence: persistence,
	}
	result, err := service.Enable(product)
	require.Nil(t, err)
	require.Equal(t, product, result)

	result, err = service.Disable(product)
	require.Nil(t, err)
	require.Equal(t, product, result)
}

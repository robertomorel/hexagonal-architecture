package application_test

/*
	Este arquivo de testes irá trabalhar com testes unitários.
	Então, para não precisarmos implementar ou usulizar qualquer controller de persistência, iremos ubilizar mocks e stubs.
	- Stub é uma camada que foi preparadas para manipular e trabalhar com os mocks
	- Bibliotecas: mockgen (dockerfile) e gomock

	Par usar o mockgen:
		> docker-compose up -d
		> docker-compose ps (cheking)
        > docker exec -it appproduct bash
		// Gerar os arquivos de mock na pasta "destination", aplicando todas as interfaces na pasta source
		> mockgen -destination=application/mocks/application.go -source=application/product.go application
*/

import (
	"testing"

	"github.com/codeedu/go-hexagonal/application"
	mock_application "github.com/codeedu/go-hexagonal/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestProductService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)

	// O defer faz com que após todos os comandos abaixo terminarem de eecutar, ele roda e finaliza o controller
	defer ctrl.Finish()
	// Criando Product fake
	product := mock_application.NewMockProductInterface(ctrl)
	// Criando Persistence fake
	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	// Mockando o método GET para retornar uma instância de Product sempre que for chamado, independente do valor
	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()
	service := application.ProductService{Persistence: persistence}

	result, err := service.Get("abc")
	require.Nil(t, err)
	require.Equal(t, product, result)
}

func TestProductService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	product := mock_application.NewMockProductInterface(ctrl)
	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()
	service := application.ProductService{Persistence: persistence}

	result, err := service.Create("Product 1", 10)
	require.Nil(t, err)
	require.Equal(t, product, result)
}

func TestProductService_EnableDisable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	product := mock_application.NewMockProductInterface(ctrl)
	product.EXPECT().Enable().Return(nil)
	product.EXPECT().Disable().Return(nil)

	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()
	service := application.ProductService{Persistence: persistence}

	result, err := service.Enable(product)
	require.Nil(t, err)
	require.Equal(t, product, result)

	result, err = service.Disable(product)
	require.Nil(t, err)
	require.Equal(t, product, result)
}

package application

import (
	entity "github.com/EvyOliveira/fullCycle/application/entity"
)

type ProductService struct {
	Persistence entity.ProductPersistanceInterface
}

func (s *ProductService) Get(id string) (entity.ProductInterface, error) {
	product, err := s.Persistence.Get(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *ProductService) Create(name string, price float64) (entity.ProductInterface, error) {
	product := entity.NewProduct()
	product.Name = name
	product.Price = price

	_, err := product.IsValid()
	if err != nil {
		return &entity.Product{}, err
	}

	result, err := s.Persistence.Save(product)
	if err != nil {
		return &entity.Product{}, err
	}
	return result, nil
}

func (s *ProductService) Enable(product entity.ProductInterface) (entity.ProductInterface, error) {
	err := product.Enable()
	if err != nil {
		return &entity.Product{}, err
	}

	result, err := s.Persistence.Save(product)
	if err != nil {
		return &entity.Product{}, err
	}
	return result, nil
}

func (s *ProductService) Disable(product entity.ProductInterface) (entity.ProductInterface, error) {
	err := product.Disable()
	if err != nil {
		return &entity.Product{}, err
	}

	result, err := s.Persistence.Save(product)
	if err != nil {
		return &entity.Product{}, err
	}
	return result, nil
}

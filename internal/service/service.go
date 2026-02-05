package service

import (
	"yosem/internal/repository"
)

type Service struct {
	Storage *repository.Storage
}

func NewService(s *repository.Storage) *Service {
	return &Service{
		Storage: s,
	}
}

// goods

// get goods service
func (s *Service) Goods() error {
	return nil
}

// add goods on storage
func (s *Service) AddGoods() error {
	return nil
}

// update goods on storage
func (s *Service) UpdateGoods() error {
	return nil
}

// delete goods from storage
func (s *Service) DeleteGoods() error {
	return nil
}

// users

// get users from storage
func (s *Service) Users() error {
	return nil
}

// orders
func (s *Service) Orders() error {
	return nil
}

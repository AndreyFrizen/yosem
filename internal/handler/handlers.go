package handler

import (
	"yosem/internal/service"
)

type Handler struct {
	Service *service.Service
}

func NewService(s *service.Service) *Handler {
	return &Handler{
		Service: s,
	}
}

// goods

// get goods service
func (h *Handler) Goods() error {
	return nil
}

// add goods on storage
func (h *Handler) AddGoods() error {
	return nil
}

// update goods on storage
func (h *Handler) UpdateGoods() error {
	return nil
}

// delete goods from storage
func (h *Handler) DeleteGoods() error {
	return nil
}

// users

// get users from storage
func (h *Handler) Users() error {
	return nil
}

// orders
func (h *Handler) Orders() error {
	return nil
}

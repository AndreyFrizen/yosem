package handler

import (
	"net/http"
	"yosem/internal/models"
	"yosem/internal/service"

	"github.com/labstack/echo/v5"
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
func (h *Handler) Goods(c *echo.Context) error {
	if err := h.Service.Goods(); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, nil)
}

// add goods on storage
func (h *Handler) AddGoods(c *echo.Context) error {
	model := new(models.Goods)
	if err := c.Bind(model); err != nil {
		return err
	}

	if err := h.Service.AddGoods(model); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, model)
}

// update goods on storage
func (h *Handler) UpdateGoods(c *echo.Context) error {
	model := new(models.Goods)
	if err := c.Bind(model); err != nil {
		return err
	}

	if err := h.Service.AddGoods(model); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, model)
}

// delete goods from storage
func (h *Handler) DeleteGoods(c *echo.Context) error {
	id := c.Param("id")
	if err := h.Service.DeleteGoods(id); err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
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

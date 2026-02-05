package repository

import "github.com/jackc/pgx/v5/pgxpool"

type Storage struct {
	Pool *pgxpool.Pool
}

func NewStorage(pool *pgxpool.Pool) *Storage {
	return &Storage{
		Pool: pool,
	}
}

func (s *Storage) Close() {
	s.Pool.Close()
}

// goods

// get goods from storage
func (s *Storage) Goods() error {
	return nil
}

// add goods on storage
func (s *Storage) AddGoods() error {
	return nil
}

// update goods on storage
func (s *Storage) UpdateGoods() error {
	return nil
}

// delete goods from storage
func (s *Storage) DeleteGoods() error {
	return nil
}

// users

// get users from storage
func (s *Storage) Users() error {
	return nil
}

// orders
func (s *Storage) Orders() error {
	return nil
}

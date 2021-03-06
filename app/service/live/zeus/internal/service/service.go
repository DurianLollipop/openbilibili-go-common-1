package service

import (
	"context"

	"go-common/app/service/live/zeus/internal/conf"
	"go-common/app/service/live/zeus/internal/dao"
)

// Service struct
type Service struct {
	c   *conf.Config
	dao *dao.Dao
}

// New init
func New(c *conf.Config) (s *Service) {
	s = &Service{
		c:   c,
		dao: dao.New(c),
	}
	return s
}

// Ping Service
func (s *Service) Ping(ctx context.Context) (err error) {
	//return s.dao.Ping(ctx)
	return nil
}

// Close Service
func (s *Service) Close() {
	s.dao.Close()
}

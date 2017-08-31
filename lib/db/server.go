package db

import (
	"github.com/go-redis/redis"
)

type Server struct {
	*redis.Client
}

func NewClient() (s *Server) {
	s = &Server{}
	s.Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return s
}

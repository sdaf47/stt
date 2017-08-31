package db

import (
	"github.com/go-redis/redis"
	"reflect"
	"fmt"
)

type Server struct {
	*redis.Client
	ReleasesList []int
	Releases map[int]Release
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

func (s *Server) AddEntity(e Entity) (int) {
	val := reflect.ValueOf(e).Elem()

	nameEntity := val.Type().Name()
	nameIterator := nameEntity + "-iterator"

	id := s.LLen(nameIterator).Val()
	fmt.Println("id", id)

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i).String()
		typeField := val.Type().Field(i)
		tag := typeField.Tag

		key := fmt.Sprintf("%s:%d:%s", nameEntity, id, tag.Get("redis"))

		fmt.Println(key, valueField)

		s.Set(key, valueField, 0)
	}

	s.RPush(nameIterator, id)

	return int(id)
}

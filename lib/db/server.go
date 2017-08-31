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

func (s *Server) GetEntity(m Entity, id int) (interface{}) {
	valRef := reflect.ValueOf(m).Elem()
	typeRef := reflect.TypeOf(m).Elem()

	rv := reflect.New(typeRef)

	nameEntity := valRef.Type().Name()

	for i := 0; i < valRef.NumField(); i++ {
		typeField := valRef.Type().Field(i)
		tag := typeField.Tag

		key := fmt.Sprintf("%s:%d:%s", nameEntity, id, tag.Get("redis"))

		redValue := s.Get(key)

		switch typeField.Type.String() {
		case "string":
			rv.Elem().Field(i).SetString(redValue.Val())
		case "int":
			val, _ := redValue.Int64()
			rv.Elem().Field(i).SetInt(val)
		}

		fmt.Println(key)
	}

	reflect.Indirect(rv)
	return rv.Interface()
}

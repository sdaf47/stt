package db

import (
	"time"
)

type Entity interface {

}

type Step struct {
	Id     int `redis:"id"`
	Action string `redis:"action"`
	Result string `redis:"result"`
}

type Case struct {
	Id     int `redis:"id"`
	Name   string `redis:"name"`
	Desc   string `redis:"desc"`
	Status bool `redis:"status"`
	Steps  []int `redis:"steps"`
}

type Release struct {
	Id    int `redis:"id"`
	Name  string `redis:"name"`
	Desc  string `redis:"desc"`
	Date  time.Duration `redis:"date"`
	Cases []int `redis:"cases"`
}

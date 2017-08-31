package db

import (
	"time"
)

type Entity interface {

}

type Step struct {
	Action string `redis:"action"`
	Result string `redis:"result"`
}

type Case struct {
	Name   string `redis:"name"`
	Desc   string `redis:"desc"`
	Status bool `redis:"status"`
	Steps  []int `redis:"steps"`
}

type Release struct {
	Name  string `redis:"name"`
	Desc  string `redis:"desc"`
	Date  time.Duration `redis:"date"`
	Cases []int `redis:"cases"`
}

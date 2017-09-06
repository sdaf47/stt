package db

import (
	"time"
)

type Entity interface{}

type Step struct {
	Action string
	Result string
}

type Case struct {
	Name   string
	Desc   string
	Status bool
	Steps  []Step
}

type Release struct {
	Name  string
	Desc  string
	Date  time.Time
	Cases []Case
}

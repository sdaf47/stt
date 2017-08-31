package db

import "time"

type Entity interface {

}

type Step struct {
	Id int `json:"id"`
	Action string `json:"action"`
	Result string `json:"result"`
}

type Case struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
	Status string `json:"status"`
	Steps []Step `json:"steps"`
}

type Release struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
	Date time.Duration `json:"date"`
	Cases []Case `json:"cases"`
}

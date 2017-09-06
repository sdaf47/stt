package db

import (
	"gopkg.in/mgo.v2-unstable"
)

type Server struct {
	ReleasesList []int
	Releases     map[int]Release
}

func NewClient(HostPort string) (s *mgo.Session, err error) {
	s, err = mgo.Dial(HostPort)
	if err != nil {
		return nil, err
	}

	s.SetMode(mgo.Monotonic, true)

	return s, nil

}

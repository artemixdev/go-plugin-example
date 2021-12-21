package storage

import "errors"

type Record struct {
	ID   uint
	Name string
}

type RecordStorage interface {
	FindOne(id uint) (*Record, error)
	FindMany() ([]*Record, error)
}

var ErrRecordNotFound = errors.New("record not found")

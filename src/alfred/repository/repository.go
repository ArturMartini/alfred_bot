package repository

import (
	"alfred/canonical"
	"alfred/repository/badgerdb"
)

var instance Repository

type Repository interface {
	Update(config canonical.Configuration) error
}


func New() Repository {
	if instance == nil {
		instance = badgerdb.Initialize()
	}
	return instance
}

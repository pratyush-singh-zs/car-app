package datastore

import "github.com/Pk05999/car-app/entities"

type Car interface {
	Get(id int) ([]entities.Car, error)
	Create(entities.Car) (entities.Car, error)
}

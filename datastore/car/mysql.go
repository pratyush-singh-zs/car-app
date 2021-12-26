package car

import (
	"database/sql"

	"github.com/pratyush-singh-zs/car-app/entities"
)

type CarStore struct {
	db *sql.DB
}

func New(db *sql.DB) CarStore {
	return CarStore{db: db}
}

func (a CarStore) Get(id int) ([]entities.Car, error) {
	var (
		rows *sql.Rows
		err  error
	)
	if id != 0 {
		rows, err = a.db.Query("SELECT * FROM car where id = ?", id)
	} else {
		rows, err = a.db.Query("SELECT * FROM car")
	}

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var car []entities.Car

	for rows.Next() {
		var a entities.Car
		_ = rows.Scan(&a.ID, &a.Name, &a.Price)
		car = append(car, a)
	}
	return car, nil
}

func (a CarStore) Create(car entities.Car) (entities.Car, error) {
	res, err := a.db.Exec("INSERT INTO car (name, price) VALUES(?, ?)", car.Name, car.Price)
	if err != nil {
		return entities.Car{}, err
	}
	id, _ := res.LastInsertId()
	car.ID = int(id)
	return car, nil
}

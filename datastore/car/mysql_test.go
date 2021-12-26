package car

import (
	"database/sql"
	"os"
	"reflect"
	"testing"

	"github.com/pratyush-singh-zs/car-app/driver"
	"github.com/pratyush-singh-zs/car-app/entities"
)

func insilizeMySQL(t *testing.T) *sql.DB {
	conf := driver.MySQLConfig{
		Host:     os.Getenv("SQL_HOST"),
		User:     os.Getenv("SQL_USER"),
		Password: os.Getenv("SQL_PASSWORD"),
		Port:     os.Getenv("SQL_PORT"),
		Db:       os.Getenv("SQL_DB"),
	}
	var err error
	db, err := driver.ConnectToMySQL(conf)
	if err != nil {
		t.Errorf("could not cnnect to sql, err:%v", err)
	}
	return db
}

func TestDatastore(t *testing.T) {
	db := insilizeMySQL(t)
	a := New(db)
	testCarStore_Get(t, a)
	testCarStore_Create(t, a)
}

func testCarStore_Create(t *testing.T, db CarStore) {
	testcases := []struct {
		req      entities.Car
		response entities.Car
	}{
		{entities.Car{Name: "Audi", Price: 1000}, entities.Car{3, "Audi", 1}},
		{entities.Car{Name: "Lambo", Price: 2000}, entities.Car{4, "Lambo", 2}},
	}
	for i, v := range testcases {
		resp, _ := db.Create(v.req)
		if !reflect.DeepEqual(resp, v.response) {
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, resp, v.response)
		}
	}
}

func testCarStore_Get(t *testing.T, db CarStore) {
	testcases := []struct {
		id   int
		resp []entities.Car
	}{
		{0, []entities.Car{{1, "amb", 1055}, {2, "nexon", 250}}},
		{1, []entities.Car{{1, "luna", 105}}},
	}
	for i, v := range testcases {
		resp, _ := db.Get(v.id)

		if !reflect.DeepEqual(resp, v.resp) {
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, resp, v.resp)
		}
	}
}

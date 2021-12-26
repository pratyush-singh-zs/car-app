package car

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/pratyush-singh-zs/car-app/entities"
)

func TestCarHandler_Handler(t *testing.T) {
	testcases := []struct {
		method             string
		expectedStatusCode int
	}{
		{"GET", http.StatusOK},
		{"POST", http.StatusOK},
		{"DELETE", http.StatusMethodNotAllowed},
	}
	for _, v := range testcases {
		req := httptest.NewRequest(v.method, "/car", nil)
		w := httptest.NewRecorder()

		a := New(mockDatastore{})
		a.Hander(w, req)

		if w.Code != v.expectedStatusCode {
			t.Errorf("Expected %v\tGot %v", v.expectedStatusCode, w.Code)
		}
	}
}

func TestCarGet(t *testing.T) {
	testcases := []struct {
		id       string
		response []byte
	}{
		{"1", []byte("could not retrieve car")},
		{"1a", []byte("invalid parameter id")},
		{"2", []byte(`[{"ID":2,"Name":"Audi","Price":1000}]`)},
		{"0", []byte(`[{"ID":1,"Name":"Lambo","Price":7623},{"ID":2,"Name":"Nexon","Price":8978}]`)},
	}

	for i, v := range testcases {
		req := httptest.NewRequest("GET", "/car?id="+v.id, nil)
		w := httptest.NewRecorder()

		a := New(mockDatastore{})

		a.get(w, req)

		if !reflect.DeepEqual(w.Body, bytes.NewBuffer(v.response)) {
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, w.Body.String(), string(v.response))
		}
	}
}

func TestCarPost(t *testing.T) {
	testcases := []struct {
		reqBody  []byte
		respBody []byte
	}{
		{[]byte(`{"Name":"Fiat","Price":1287}`), []byte(`could not create car`)},
		{[]byte(`{"Name":"Alha","Price":1078}`), []byte(`{"ID":12,"Name":"Maggie","Age":10}`)},
		{[]byte(`{"Name":"Maggie","Price":"10"}`), []byte(`invalid body`)},
	}
	for i, v := range testcases {
		req := httptest.NewRequest("GET", "/animal", bytes.NewReader(v.reqBody))
		w := httptest.NewRecorder()

		a := New(mockDatastore{})

		a.create(w, req)

		if !reflect.DeepEqual(w.Body, bytes.NewBuffer(v.respBody)) {
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, w.Body.String(), string(v.respBody))
		}
	}
}

type mockDatastore struct{}

func (m mockDatastore) Get(id int) ([]entities.Car, error) {
	if id == 1 {
		return nil, errors.New("db error")
	} else if id == 2 {
		return []entities.Car{{2, "Audi", 898}}, nil
	}
	return []entities.Car{{1, "kead", 2376}, {2, "Bolearo", 8787}}, nil
}

func (m mockDatastore) Create(car entities.Car) (entities.Car, error) {
	if car.Price == 9 {
		return entities.Car{}, errors.New("db error")
	}

	return entities.Car{9, "Pratyush", 876}, nil
}

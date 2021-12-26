package car

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/pratyush-singh-zs/car-app/datastore"
	"github.com/pratyush-singh-zs/car-app/entities"
)

type carHandler struct {
	datastore datastore.Car
}

func New(car datastore.Car) carHandler {
	return carHandler{datastore: car}
}

func (a carHandler) Hander(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		a.get(w, r)
	case http.MethodPost:
		a.create(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (a carHandler) get(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	i, err := strconv.Atoi(id)
	if err != nil {
		_, _ = w.Write([]byte("invalid parameter id"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	resp, err := a.datastore.Get(i)
	if err != nil {
		_, _ = w.Write([]byte("could not retrive car"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	body, _ := json.Marshal(resp)
	_, _ = w.Write(body)
}

func (a carHandler) create(w http.ResponseWriter, r *http.Request) {
	var car entities.Car
	body, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal(body, &car)
	if err != nil {
		fmt.Println(err)
		_, _ = w.Write([]byte("invalid body"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = a.datastore.Create(car)
	if err != nil {
		_, _ = w.Write([]byte("could not create car"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// body, _ = json.Marshal(resp)
	_, _ = w.Write([]byte("success"))
}

func (a carHandler) update(w http.ResponseWriter, r *http.Request) {

}

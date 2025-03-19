package handler

import (
	"Server/product"
	"encoding/json"
	"fmt"
	"net/http"
)

type Intial struct{}

func Initializer() *Intial {
	return &Intial{}
}

func (c *Intial) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		c.getrequest(rw, r)
	case http.MethodPost:
		c.postrequest(rw, r)
	case http.MethodDelete:
		c.deleterequest(rw, r)
	case http.MethodPut:
		c.putrequest(rw, r)
	}
}

func (c *Intial) getrequest(rw http.ResponseWriter, r *http.Request) {
	data, err := product.Getfunc()
	if err != nil {
		fmt.Printf("Error while Getting the data")
	}
	if len(data) == 0 {
		http.Error(rw, "No Recorde found", http.StatusNotFound)
	}
	marshdata, err := json.Marshal(data)
	if err != nil {
		http.Error(rw, "Error while Marshaling the data ", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(marshdata)
}

func (c *Intial) postrequest(rw http.ResponseWriter, r *http.Request) {
	err := product.Addfunc(r.Body)
	if err != nil {
		http.Error(rw, "Invalid request payload", http.StatusBadRequest)
		return
	}
	rw.WriteHeader(http.StatusCreated)
	rw.Write([]byte(`{"message": "Record added successfully"}`))
}
func (c *Intial) deleterequest(rw http.ResponseWriter, r *http.Request) {
	product.Deletefunc(r.Body)
}
func (c *Intial) putrequest(rw http.ResponseWriter, r *http.Request) {
	product.Replacefunc(r.Body)
}

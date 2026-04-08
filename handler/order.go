package handler

import (
    "fmt"
    "net/http"
)

type Order struct {}

func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Order created")
}

func (o *Order) List(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "All Order details")
}

func (o *Order) GetById(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Order details")
}

func (o *Order) Update(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Order updated")
}

func (o *Order) Delete(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Order deleted")
}

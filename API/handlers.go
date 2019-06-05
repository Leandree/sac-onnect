package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Index Handler
func Index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
    _, err := w.Write([]byte("Welcome!"))
    if err != nil {
        log.Println(err.Error())
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

//SaveUsers create a user
func SaveUsers(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
    body := req.Body
    usr := NewUser()
    err := json.NewDecoder(body).Decode(usr)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer body.Close()
    err = usr.Save()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}

//SaveUsers create a user
func SaveWeight(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
    body := req.Body
    wght := NewWeight()
    err := json.NewDecoder(body).Decode(wght)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer body.Close()
    err = wght.Save()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}

// UpdateUser update an existing user
func UpdateItem(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
    //ID := params.ByName("idTag")
    var i Item
    err := json.NewDecoder(req.Body).Decode(&i)
    defer req.Body.Close()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    if i.IdTag == "" {
        http.Error(w, "Please set ID TAG in User information", http.StatusBadRequest)
        return
    }
    err = NewItem().Update(i)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}
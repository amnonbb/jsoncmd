package main

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"io/ioutil"
	"fmt"
)

func (a *App) getJson(w http.ResponseWriter, r *http.Request) {
	var s status
	vars := mux.Vars(r)
	key := vars["id"]
	value := r.FormValue("spc")

	err := s.getExec(key, value)

	if err != nil {
		s.Status = "error"
	} else {
		s.Status = "ok"
	}

	respondWithJSON(w, http.StatusOK, s)
}

func (a *App) postJson(w http.ResponseWriter, r *http.Request) {
	var s status
	vars := mux.Vars(r)
	key := vars["id"]

	var postj map[string]interface{}

	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &postj)

	err := s.postExec(key, postj)

	defer r.Body.Close()

	if err != nil {
		s.Status = "error"
	} else {
		s.Status = "ok"
	}

	respondWithJSON(w, http.StatusOK, s)
}

func (a *App) putJson(w http.ResponseWriter, r *http.Request) {
	var s status
	vars := mux.Vars(r)
	key := vars["id"]

	b, _ := ioutil.ReadAll(r.Body)

	err := s.putExec(key, string(b))

	defer r.Body.Close()

	if err != nil {
		s.Status = "error"
	} else {
		s.Status = "ok"
	}

	respondWithJSON(w, http.StatusOK, s)
}
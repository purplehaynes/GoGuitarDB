package handlers

import (
	"GuitarDB/entities"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

type Service interface {
	CreateNewGuitar(axe entities.Guitar) error
	ReadAll() (*entities.GTStruct, error)
	GetByGuitarId(id string) (*entities.Guitar, error)
	DeleteGuitarId(id string) error
	UpdateGuitarInfo(id string, axe entities.Guitar) error
}

type GuitarHandler struct {
	Svc Service
}

func NewGuitarHandler(s Service) GuitarHandler {
	return GuitarHandler {
		Svc: s,
	}
}

func (gh GuitarHandler) CreateGuitar(w http.ResponseWriter, r *http.Request) {

	gt := entities.Guitar{}

	err := json.NewDecoder(r.Body).Decode(&gt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = gh.Svc.CreateNewGuitar(gt)
	if err != nil {
		switch err.Error() {
		case "guitar already exists in database":
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		// case "invalid rating":
		// 	http.Error(w, err.Error(), http.StatusNotAcceptable)
			// return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

}

func (gh GuitarHandler) ReadAll(w http.ResponseWriter, r *http.Request) {

	readDB, err := gh.Svc.ReadAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	readRequest, err := json.MarshalIndent(readDB, "", "	")
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(readRequest)
}

func (gh GuitarHandler) GetByGuitarId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["Id"]

	getID, err := gh.Svc.GetByGuitarId(id)
	if err != nil {
		switch err.Error() {
		case "guitar not found":
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
	}

	guitarInfo, err := json.MarshalIndent(getID, "", "	")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(guitarInfo)
}

func (gh GuitarHandler) DeleteGuitarId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["Id"]

	err := gh.Svc.DeleteGuitarId(id)
	if err != nil {
		switch err.Error() {
		case "guitar does not exist":
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (gh GuitarHandler) UpdateGuitarInfo(w http.ResponseWriter, r *http.Request) {
	gt := entities.Guitar{}
	vars := mux.Vars(r)
	id := vars["Id"]

	err := json.NewDecoder(r.Body).Decode(&gt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = gh.Svc.UpdateGuitarInfo(id, gt)
	if err != nil {
		switch err.Error() {
		case "id is mismatched":
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
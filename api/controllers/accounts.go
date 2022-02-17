package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/Kerlense/form3/api/models"
	"github.com/Kerlense/form3/api/reply"
	
)

func (s *DBServer) CreateAccount(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		reply.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	account := models.Account{}
	err = json.Unmarshal(body, &account)
	if err != nil {
		reply.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	
	accountCreated, err := account.SaveAccount(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		reply.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Lacation", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, accountCreated.ID))
	reply.JSON(w, http.StatusCreated, accountCreated)
}



func (s *DBServer) GetAccount(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		reply.ERROR(w, http.StatusBadRequest, err)
		return
	}
	account := models.Account{}

	accountReceived, err := account.FindAccountbyID(s.DB, pid)
	if err != nil {
		reply.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	reply.JSON(w, http.StatusOK, postReceived)
}


func (s *DBServer) DeleteAccount(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Is a valid post id given to us?
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		reply.ERROR(w, http.StatusBadRequest, err)
		return
	}

	// Check if the post exist
	account := models.Account{}
	err = server.DB.Debug().Model(models.Account{}).Where("id = ?", pid).Take(&account).Error
	if err != nil {
		reply.ERROR(w, http.StatusNotFound, errors.New("Unauthorized"))
		return
	}

	_, err = account.DeleteAccount(s.DB, pid, uid)
	if err != nil {
		reply.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", pid))
	responses.JSON(w, http.StatusNoContent, "")
}
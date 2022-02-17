package controllers

import (
	"net/http"

	"github.com/Kerlense/form3/api/reply"
)

func (s *Server) Home(w http.ResponseWriter, r *http.Request) {
	reply.JSON(w, http.StatusOK, "Welcome")

}

package controllers

import "github.com/Kerlense/form3/api/middleware"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middleware.SetMiddlewareJSON(s.Home)).Methods("GET")

	//Account routes
	s.Router.HandleFunc("/accounts", middleware.SetMiddlewareJSON(s.GetAccount)).Methods("GET")
	s.Router.HandleFunc("/accounts", middleware.SetMiddlewareJSON(s.CreateAccount)).Methods("POST")
	s.Router.HandleFunc("/accounts/{account_id}", middleware.SetMiddlewareAuthentication(s.DeleteAccount)).Methods("DELETE")
	s.Router.HandleFunc("/accounts/{account_id}", middleware.SetMiddlewareJSON(s.GetAccount)).Methods("GET")

}

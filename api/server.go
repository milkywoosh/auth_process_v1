package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"oraluke.com/conn-ora1/services"
)

// NOTE: token use JWT or PASETO
type Server struct {
	route   *httprouter.Router
	service *services.Services
	token   string
}

// NOTE: pendefinisian untuk response harus tersendiri, berbeda dengan yg di domain
type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

func NewServer(services *services.Services) (*Server, error) {

	// need token maker here JWT or Paseto

	s := &Server{
		route:   httprouter.New(),
		service: services,
	}

	s.setupRoutes()
	return s, nil
}

// TEST write other response like base64 or other
func (s *Server) writeJSON(w http.ResponseWriter, statusCode int, data any) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func (s *Server) writeError(w http.ResponseWriter, statusCode int, message string) {
	s.writeJSON(w, statusCode, Response{
		Success: false,
		Error:   message,
	})
}

// TEST REQ TYPE
type TestReqBody struct {
	Param1 string `json:"param1"`
	Param2 string `json:"param2"`
}

func (s *Server) setupRoutes() {
	// define each routes here, must satisfy interface httprouter.Handle
	s.route.GET("/health-check", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		// how to parse body request from http.Request?
		/*
			says i have
			{
				"username": string
				"password": string
			}
		*/

		var ReqBody TestReqBody
		// read input from Body
		errDecode := json.NewDecoder(r.Body).Decode(&ReqBody)
		if errDecode != nil {

			s.writeError(w, http.StatusBadRequest, errDecode.Error())
			return
		}

		s.writeJSON(w, http.StatusOK, Response{
			Success: true,
			Message: "health check okay",
			Data:    ReqBody,
		})
	})

	s.stockTransferRoutes() // each of stock_transfer route

}

func (s *Server) stockTransferRoutes() {
	s.route.POST("/stock-transfer/create", nil)
	s.route.GET("/stock-transfer/header", nil)
	s.route.PUT("/stock-transfer/submit", nil)
	s.route.PUT("/stock-transfer/reject", nil)
	s.route.PUT("/stock-transfer/approve", nil)
	s.route.PUT("/stock-transfer/cancel", nil)
}

func (s *Server) Start(addr string) error {
	return http.ListenAndServe(addr, s.route)
}

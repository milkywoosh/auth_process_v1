package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"oraluke.com/conn-ora1/util"
)

type reqLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// note: oprek context
func (s *Server) Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	w.Header().Set("Content-Type", "application/json") // must be always first at handler
	var req reqLogin

	// read input from Body
	defer r.Body.Close()
	errDecode := json.NewDecoder(r.Body).Decode(&req)
	if errDecode != nil {
		s.writeError(w, http.StatusBadRequest, errDecode.Error())
		return
	}

	userInfo, err := s.service.GetInfoUser(r.Context(), req.Username)

	if err != nil {
		s.writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = util.ComparePassword(userInfo.Password, req.Password)
	if err != nil {
		s.writeError(w, http.StatusNotAcceptable, err.Error())
		return
	}
	// w.WriteHeader(http.StatusAccepted)
	// json.NewEncoder(w).Encode("error")

	s.writeJSON(w, http.StatusOK, Response{
		Success: true,
		Message: "Success",
		Data:    userInfo,
	})

}

type reqRegistration struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (s *Server) UserRegistration(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json") // must be always first at handler

	var registration reqRegistration

	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&registration)
	if err != nil {
		s.writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	hashedPW, err := util.HashPassword(registration.Password)
	if err != nil {
		s.writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = s.service.UserRegistration(r.Context(), registration.Username, hashedPW)
	if err != nil {
		s.writeError(w, http.StatusExpectationFailed, err.Error())
		return
	}

	s.writeJSON(w, http.StatusOK, Response{
		Success: true,
		Message: "success create new user",
		Data:    nil,
	})

}

package services

import dbora "oraluke.com/conn-ora1/db"

type Services struct {
	*AuthenticationService
}

func NewServices(db *dbora.OraDBInstance) *Services {
	return &Services{
		AuthenticationService: NewAuthentication(db),
	}
}

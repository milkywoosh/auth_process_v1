package services

import dbora "oraluke.com/conn-ora1/db"

type Services struct {
}

func NewServices(db *dbora.OraDBInstance) *Services {
	return &Services{}
}

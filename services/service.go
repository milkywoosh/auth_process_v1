package services

import dbora "oraluke.com/conn-ora1/db"

type Services struct {
	*StockTransferSrv
	*StockReceiveSrv
}

func NewServices(db *dbora.OraDBInstance) *Services {
	return &Services{
		StockTransferSrv: NewStockTransferService(db),
		StockReceiveSrv:  NewStockReceiveService(db),
	}
}

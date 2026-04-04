package dbora

import (
	"database/sql"
)

type oraTransaction struct {
	tx *sql.Tx
	// locationRepo       domain.LocationRepository // define in db layer
	// itemRepo           domain.ItemRepository     // define in db layer
	// itemDomainSrv      *domain.ItemDomainService
	// warehouseDomainSrv *domain.WarehouseDomainService
}

/*
func newOraTransaction(tx *sql.Tx) *oraTransaction {

	newLocationRepo := NewDBLocationRepository(tx)
	newItemRepo := NewDBItemRepository(tx)
	newItemTransferRepo := NewDBItemTransferRepository(tx)
	newItemDomainSrv := domain.NewItemDomainService(newItemRepo)
	newWarehouseDomainSrv := domain.NewWarehouseDomainSrv(newItemDomainSrv, newLocationRepo, newItemTransferRepo)

	return &oraTransaction{
		tx:                 tx,
		locationRepo:       newLocationRepo,
		itemRepo:           newItemRepo,
		itemDomainSrv:      newItemDomainSrv,
		warehouseDomainSrv: newWarehouseDomainSrv,
	}
}

func (ot *oraTransaction) LocationRepo() domain.LocationRepository {
	return ot.locationRepo
}

func (ot *oraTransaction) ItemRepo() domain.ItemRepository {
	return ot.itemRepo
}

func (ot *oraTransaction) ItemDomainSrv() *domain.ItemDomainService {
	return ot.itemDomainSrv
}

func (ot *oraTransaction) WarehouseDomainSrv() *domain.WarehouseDomainService {
	return ot.warehouseDomainSrv
}

func (ot *oraTransaction) Commit() error {
	return ot.tx.Commit()
}

func (ot *oraTransaction) Rollback() error {
	return ot.tx.Rollback()
}

*/

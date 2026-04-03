package dbora

import (
	"database/sql"

	"oraluke.com/conn-ora1/domain"
)

type oraQuery struct {
	db                 *sql.DB
	locationRepo       domain.LocationRepository // define in db layer
	itemRepo           domain.ItemRepository     // define in db layer
	itemDomainSrv      *domain.ItemDomainService
	itemTransfer       domain.ItemTransfer
	warehouseDomainSrv *domain.WarehouseDomainService
}

func newOraQuery(db *sql.DB) *oraQuery {

	newLocationRepo := NewDBLocationRepository(db)
	newItemRepo := NewDBItemRepository(db)
	newItemTransferRepo := NewDBItemTransferRepository(db)
	newItemDomainSrv := domain.NewItemDomainService(newItemRepo)
	newWarehouseDomainSrv := domain.NewWarehouseDomainSrv(newItemDomainSrv, newLocationRepo, newItemTransferRepo)

	return &oraQuery{
		db:                 db,
		locationRepo:       newLocationRepo,
		itemRepo:           newItemRepo,
		itemTransfer:       newItemTransferRepo,
		itemDomainSrv:      newItemDomainSrv,
		warehouseDomainSrv: newWarehouseDomainSrv,
	}
}

func (ot *oraQuery) LocationRepo() domain.LocationRepository {
	return ot.locationRepo
}

func (ot *oraQuery) ItemRepo() domain.ItemRepository {
	return ot.itemRepo
}

func (ot *oraQuery) ItemDomainSrv() *domain.ItemDomainService {
	return ot.itemDomainSrv
}
func (ot *oraQuery) ItemTransfer() domain.ItemTransfer {
	return ot.itemTransfer
}

func (ot *oraQuery) WarehouseDomainSrv() *domain.WarehouseDomainService {
	return ot.warehouseDomainSrv
}

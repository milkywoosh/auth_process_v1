package services

import (
	"context"
	"errors"
	"fmt"
	"log"

	dbora "oraluke.com/conn-ora1/db"
	"oraluke.com/conn-ora1/domain"
	"oraluke.com/conn-ora1/util"
)

/*
	==================================================

	NEVER PUT LOGIC BUSSINESS <IF> IN THIS LAYER

	orchestrating aplication Logic Level

	orchestrating concrete implementation from db/infra function

	remember, bussiness RULES should be defined inside domain layer

	==================================================
*/

type StockTransferSrv struct {
	uow domain.UnitOfWork
	q   domain.Queries
}

func NewStockTransferService(dbOraInstance *dbora.OraDBInstance) *StockTransferSrv {
	return &StockTransferSrv{
		uow: dbOraInstance,
		q:   dbOraInstance,
	}
}

// NOTE: duplication with
func (t *StockTransferSrv) Send(ctx context.Context, currentWarehouseCode string, toLocationCode string, items []*domain.Item) error {

	log.Printf("start send ..............")

	tx, err := t.uow.BeginStockTransfer(ctx)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	// get info current WH currentWarehouseInfo
	_, errWhInfo := tx.WarehouseDomainSrv().CurrentWarehouseInfo(ctx, currentWarehouseCode)
	if errWhInfo != nil {
		return errWhInfo
	}

	var allQRCode []string = make([]string, len(items))

	for _, val := range items {

		allQRCode = append(allQRCode, val.QRCode)
	}

	errValidate := tx.WarehouseDomainSrv().ValidateSendItem(ctx, items)
	if errValidate != nil {
		return errValidate
	}

	log.Printf("pass here!")

	for _, val := range items {
		errSetAllocated := tx.WarehouseDomainSrv().ItemDomainSrv.SetAllocated(ctx, val.QRCode)
		if errSetAllocated != nil {
			return errSetAllocated
		}
	}

	// return fmt.Errorf("cuss test err trigger rollback")

	if errCommit := tx.Commit(); errCommit != nil {
		return errCommit
	}
	return nil
}

func (t *StockTransferSrv) GenerateTransactionNumber() string {

	randInt_ := util.RandomInt(100, 1000)
	randStr_ := util.RandomStringAllUpperCase(10)

	generatedString := fmt.Sprintf("WH-WH-%s%d", randStr_, randInt_)

	return generatedString
}

func (t *StockTransferSrv) Create(ctx context.Context, currentWarehouseCode string, toWarehouseCode string) error {

	// SHOUD MOVE THIS VALIDATION TO DOMAIN LAYER
	if currentWarehouseCode == toWarehouseCode {
		errMsg := "error proses stock transfer tidak boleh memiliki asal dan tujuan yang sama"
		return errors.New(errMsg)
	}

	if currentWarehouseCode == "" || toWarehouseCode == "" {
		errMsg := "error proses stock transfer tidak boleh memiliki asal dan tujuan yang kosong"
		return errors.New(errMsg)

	}

	generatedNoTrans := t.GenerateTransactionNumber()

	tx, err := t.uow.BeginStockTransfer(ctx)
	if err != nil {
		log.Printf("err begin st: %v", err)
		return err
	}

	defer tx.Rollback()

	locInfoCurrWH, errCurrWH := tx.WarehouseDomainSrv().LocationRepo.FetchByLocationCode(ctx, currentWarehouseCode)
	if errCurrWH != nil {
		return errCurrWH
	}

	locInfoToWH, errToWH := tx.WarehouseDomainSrv().LocationRepo.FetchByLocationCode(ctx, toWarehouseCode)
	if errToWH != nil {
		return errToWH
	}

	st := domain.CreateNewStockTransferParams{
		StNumber:      generatedNoTrans,
		TipeSTId:      6,
		CurrWarehouse: locInfoCurrWH.Id,
		ToWarehouse:   locInfoToWH.Id,
		CreatedBy:     0,
	}

	_, erCreate := tx.WarehouseDomainSrv().ItemTransfer.CreateStockTransfer(ctx, st)
	// log.Printf("check err CreateStockTransfer: %s", erCreate.Error())
	if erCreate != nil {
		log.Printf("err erCreate st: %v", erCreate)
		return erCreate
	}

	// log.Printf("b4 commit no stock transfers: %s", generatedNoTrans)
	return tx.Commit()
}

func (t *StockTransferSrv) SetBackDraft(ctx context.Context, identifier string) error {
	tx, err := t.uow.BeginStockTransfer(ctx)
	if err != nil {
		return err
	}

	defer tx.Rollback()
	// TEST
	exists, errExists := tx.WarehouseDomainSrv().ItemTransfer.TransactionExisted(ctx, identifier)
	if errExists != nil {
		return err
	}
	if !exists {
		return errors.New("Gagal update ke status DRAFT")
	}
	return tx.Commit()
}

func (t *StockTransferSrv) Header(ctx context.Context, identifier string) ([]domain.StockTransferInfo, error) {
	db, err := t.q.StockTransferQueries(ctx)
	if err != nil {
		return []domain.StockTransferInfo{}, err
	}

	datas, errDatas := db.StockTransferHeader(ctx, identifier)
	if errDatas != nil {
		return []domain.StockTransferInfo{}, errDatas
	}

	return datas, nil

}
func (t *StockTransferSrv) Details(ctx context.Context, identifier string) ([]domain.StockTransferDetail, error) {
	db, err := t.q.StockTransferQueries(ctx)
	if err != nil {
		return []domain.StockTransferDetail{}, err
	}
	return db.StockTransferDetails(ctx, identifier)

}

func (t *StockTransferSrv) InsertDetails(ctx context.Context, noTrans string, identifier []string) error {
	return nil
	//
}

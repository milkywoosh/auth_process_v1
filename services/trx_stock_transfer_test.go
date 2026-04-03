package services

import (
	"context"
	"database/sql"
	"log"
	"testing"

	dbora "oraluke.com/conn-ora1/db"
	"oraluke.com/conn-ora1/domain"
)

/*
	==================================================
	orchestrating aplication Logic Level

	orchestrating concrete implementation from db/infra function

	==================================================
*/

func TestStockTransfer1(t *testing.T) {

	ctxBg := context.Background()

	// implement UnitOrWork and Queries interface
	dbOraInstance := dbora.NewOraDBInstance(SetupTestDBOra)

	// accept interface Unit Of Work
	TestSTAppSrv := NewStockTransferService(dbOraInstance)

	/*
		Id              string
		QRCode          string
		ProductId       string
		LocationId      sql.NullString
		InventoryStatus string
		MaterialStatus  string
		StockTransferId sql.NullString
		UpdatedAt       *time.Time

	*/

	allQRCodes := []string{
		"2AYJ8PRAYD",
		"HF8TNFU3SM",
		"HDRZ1GSI31",
		"GSGMINI748",
		"T3P081G4TEB9",
		"6565890",
		"6627353",
		"6696256",
		"6694529",
		"6694554",
		"TE30A3ZC9SFW",
		"TC248ZMPZ3",
	}

	allItems := []*domain.Item{}

	for _, QRCode := range allQRCodes {

		var itemDest domain.Item

		row := SetupTestDBOra.QueryRowContext(ctxBg,
			`SELECT 
				ID,
				QR_CODE,
				PRODUCT_ID,
				LOCATION_ID,
				INVENTORY_STATUS,
				MATERIAL_STATUS,
				STOCK_TRANSFER_ID,
				UPDATED_AT 
			FROM ITEMS_INVENTORY
			WHERE QR_CODE = :QR_CODE
				`,
			sql.Named("QR_CODE", QRCode),
		)

		err := row.Scan(
			&itemDest.Id,
			&itemDest.QRCode,
			&itemDest.ProductId,
			&itemDest.LocationId,
			&itemDest.InventoryStatus,
			&itemDest.MaterialStatus,
			&itemDest.StockTransferId,
			&itemDest.UpdatedAt,
		)
		if err != nil {
			t.Fatalf("err scan fetch Item: %s", err.Error())
		}

		log.Printf("each fetch Item: %v", itemDest)

		allItems = append(allItems, &itemDest)
	}

	errST := TestSTAppSrv.Send(ctxBg, "A3206", "A3713", allItems)
	if errST != nil {
		t.Fatalf("err Stock Transfer with Locked QRCode %s", errST.Error())
	} else {

		log.Printf("all item passed to transfers")
	}

}

func TestCreate(t *testing.T) {
	// db, err := SetupTestDB(t)
	// if err != nil {
	// 	t.Fatalf("err db conn: %s", err.Error())
	// }

	defer func() {
		err := recover()
		if err != nil {
			log.Printf("Unexpected Err: %v", err)
			log.Fatalf("Unexpected Err: %v", err)
		}
	}()

	ctxBg := context.Background()

	uow := dbora.NewOraDBInstance(SetupTestDBOra)

	// accept interface Unit Of Work
	TestSTAppSrv := NewStockTransferService(uow)

	errCreate := TestSTAppSrv.Create(ctxBg, "6732", "4723")
	if errCreate != nil {
		t.Errorf("error create new stock transfer %s", errCreate.Error())
	}
}

func TestCreateConcurrent(t *testing.T) {
	testCases := []struct {
		name    string
		fromLoc string
		toLoc   string
		wantErr bool
	}{
		{name: "valid transfer", fromLoc: "6732", toLoc: "4723", wantErr: false}, // must be false
		{name: "valid transfer001", fromLoc: "6732", toLoc: "4723", wantErr: false},
		{name: "same location", fromLoc: "6732", toLoc: "6732", wantErr: true},
		{name: "invalid from", fromLoc: "", toLoc: "4723", wantErr: true},
		{name: "invalid to", fromLoc: "6732", toLoc: "", wantErr: true},
	}

	for _, tc := range testCases {
		// tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel() // marks this subtest to run concurrently
			// db, err := SetupTestDB(t)
			// if err != nil {
			// 	t.Fatalf("err db conn: %s", err.Error())
			// }

			dbInstance := dbora.NewOraDBInstance(SetupTestDBOra)
			svc := NewStockTransferService(dbInstance)

			errCreate := svc.Create(context.Background(), tc.fromLoc, tc.toLoc)
			if (errCreate != nil) != tc.wantErr {
				t.Errorf("Create() error = %v, wantErr = %v", errCreate, tc.wantErr)
			}
		})
	}
}

func TestSetBackDraft(t *testing.T) {

	testCases := []struct {
		name          string
		noTransaction string
		wantErr       bool
	}{
		{name: "set back draft ok", noTransaction: "WH-WH_20260325-163108-833-242457", wantErr: false},
		{name: "set back draft ok1", noTransaction: "WH-TECH_20260325-162909-907-34639WW", wantErr: false},
		{name: "set back draft err not found", noTransaction: "WH-WH_20260325-163108-833-242457xx", wantErr: true},
	}

	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {

			ctxBg := context.Background()

			dbInstance := dbora.NewOraDBInstance(SetupTestDBOra)
			if dbInstance == nil {
				t.Fatal("dbInstance is nil — check SetupTestDBOra config")
				return
			}
			svc := NewStockTransferService(dbInstance)
			if svc == nil {
				t.Fatal("svc is nil — NewStockTransferService returned nil")
				return
			}

			err := svc.SetBackDraft(ctxBg, tc.noTransaction)
			var receivedErr bool = (err != nil)
			log.Printf("receivedErr: %t", receivedErr)
			if receivedErr != tc.wantErr {
				// t.Logf("err condition: %s. err info: %s", tc.name, err.Error())
				// t.Logf("err condition: %s. err info: %s", tc.name, err.Error())
				t.Errorf("err condition: %s. err info: %v", tc.name, err) // use %v, not err.Error()

			} else {
				// t.Logf("cuss: %s", tc.name)
			}

		})

	}

}

package dbora

import (
	"context"
	"database/sql"
	"errors"
	"log"

	_ "github.com/godror/godror"
	"oraluke.com/conn-ora1/domain"
)

type DBTX interface {
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) // return rows
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row        // return atmost 1 rows cannot >1
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) // exec query without NO return any row
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
}

type Conn struct {
	DB           *sql.DB
	DriverExt    string
	DBConnString string
}

func NewConn(driverExt, connString string) (*Conn, error) {

	db, err := sql.Open(driverExt, connString)
	if err != nil {
		return nil, err
	}

	errPing := db.Ping()
	if errPing != nil {
		log.Printf("error Ping DB: %s\n", errPing.Error())
		return nil, errPing
	}

	return &Conn{
		DB:           db,
		DriverExt:    driverExt,
		DBConnString: connString,
	}, nil
}

type Queries struct {
	db DBTX
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

// implement UnitOfWork, Queries Interface
type OraDBInstance struct {
	db *sql.DB
}

func NewOraDBInstance(db *sql.DB) *OraDBInstance {
	return &OraDBInstance{
		db: db,
	}
}

func (or *OraDBInstance) beginQuery() (*oraQuery, error) {
	if or.db == nil {
		return nil, errors.New("DB Instance doesnt exists")
	}
	return newOraQuery(or.db), nil
}

/*


// implement UnitOfWork Interface
func (or *OraDBInstance) begin(ctx context.Context) (*oraTransaction, error) {

	tx, err := or.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: false})
	if err != nil {
		errMsg := fmt.Sprintf("err begin transaction ora: %s", err.Error())
		return nil, errors.New(errMsg)
	}
	return newOraTransaction(tx), nil

}


// implement UnitOfWork Interface
func (or *OraDBInstance) BeginStockTransfer(ctx context.Context) (domain.WarehouseSrv, error) {
	return or.begin(ctx)
}

// work it later
func (or *OraDBInstance) BeginStockReceive(ctx context.Context) (domain.WarehouseSrv, error) {
	return or.begin(ctx)
}

func (or *OraDBInstance) StockTransferQueries(ctx context.Context) (domain.ItemTransfer, error) {
	oraQuery, err := or.beginQuery()
	if err != nil {
		return nil, err
	}

	return oraQuery.itemTransfer, nil
}
*/

func (or *OraDBInstance) AuthQueries(ctx context.Context) (domain.Authentication, error) {
	oraQuery, err := or.beginQuery()
	if err != nil {
		return nil, err
	}

	return oraQuery.AuthRepo(), nil
}

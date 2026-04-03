package services

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	dbora "oraluke.com/conn-ora1/db"
)

var SetupTestDBOra *sql.DB

// NOTE : Hook for starting all the test, and shared only in this package
func TestMain(m *testing.M) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	SetupTestDB1()

	os.Exit(m.Run())
}

func SetupTestDB1() {
	log.Printf("TestMain pkg services")
	os.Setenv("PATH", os.Getenv("PATH")+";C:\\oracle\\instantclient_21_6\\windows")

	// variable dipake nanti
	conn, err := dbora.NewConn("godror", os.Getenv("ORA_CONNSTRING"))
	// log.Printf("log err: %s", err.Error())
	if err != nil {
		log.Fatalf("error init DB ====>: %s\n", err.Error())
	}

	SetupTestDBOra = conn.DB

}

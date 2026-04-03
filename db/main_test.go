package dbora

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
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
	log.Printf("TestMain pkg db")
	os.Setenv("PATH", os.Getenv("PATH")+";C:\\oracle\\instantclient_21_6\\windows")

	// variable dipake nanti
	conn, err := NewConn("godror", os.Getenv("ORA_CONNSTRING"))
	if err != nil {
		log.Fatalf("error init DB ====>: %s\n", err.Error())
	}

	log.Println("TEST PRINT HERE!")

	SetupTestDBOra = conn.DB

}

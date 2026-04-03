package dbora

import (
	"log"
	"os"
	"testing"
)

func TestDBNegativeCase1(t *testing.T) {

	log.Printf("TestDBNegativeCase1 Testing")
	// os.Setenv("PATH", os.Getenv("PATH")+";C:\\oracle\\instantclient_21_6\\windows")

	// variable dipake nanti
	_, err := NewConn("godror", os.Getenv("ORA_CONNSTRINGG"))
	// log.Printf("log err: %s", err.Error())
	if err == nil {
		// log.Fatalf("error init DB ====>: %s\n", err.Error())
		t.Errorf("errPing must NOT nil, and should give error, because credential is NOT valid")
	}

}

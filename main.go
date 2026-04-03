package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"oraluke.com/conn-ora1/api"
	dbora "oraluke.com/conn-ora1/db"
	"oraluke.com/conn-ora1/services"
)

var Conn_ *dbora.Conn
var err error
var MainConn *sql.DB

// var interrupsSignal = []os.Signal{
// 	os.Interrupt,
// 	syscall.SIGINT,
// 	syscall.SIGTERM,
// }

// very first initiation
func init() {

	errEnv := godotenv.Load("./.env")
	if errEnv != nil {
		log.Fatalf("err loading env: %s", errEnv.Error())
	}

	os.Setenv("PATH", os.Getenv("PATH")+";C:\\oracle\\instantclient_21_6\\windows")

	// variable dipake nanti
	Conn_, err = dbora.NewConn("godror", os.Getenv("ORA_CONNSTRING"))
	// log.Printf("log err: %s", err.Error())
	if err != nil {
		log.Fatalf("error init DB ====>: %s\n", err.Error())
	}

}

func main() {
	fmt.Println("go oracle test")

	dbOraInstance := dbora.NewOraDBInstance(Conn_.DB)
	newServices := services.NewServices(dbOraInstance)
	serverREST, err := api.NewServer(newServices)
	if err != nil {
		log.Fatalf("err info server: %v", err)
	}

	errStart := serverREST.Start(":8081")
	if errStart != nil {
		log.Fatalf("err start server: %v", errStart)
	}
	// init all repositories
	// init all domain service

	// ctx, stop := signal.NotifyContext(context.Background(), interrupsSignal...)
	// defer stop()

	// trial os.Signal

	// storeTransaction := db.NewStoreTransaction(Conn_.DB)

	// router.GET("/", handlerhttp.Index01)
	// log.Fatal(http.ListenAndServe(":8081", router))
}

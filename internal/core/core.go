package core

import (
    "net/http"
    "log" 
    "fmt"
    "github.com/TonyFeng628/http-server/internal/dbm"
)


type Core struct {
	dbm *dbm.DBManager
}

func NewCore() *Core {
	return &Core{
		dbm: dbm.NewDBManager("cot","]cotnetwork}","tianhua.cotnetwork.com:3306","test_tony"),
	}
}

func (core *Core) Init() {
   fmt.Println("serve files")
    http.HandleFunc("/version",core.version)
    http.HandleFunc("/create",core.createHandler)
    http.HandleFunc("/read",core.readHandler)
    http.HandleFunc("/update",core.updateHandler)
    http.HandleFunc("/delete",core.updateHandler)

}

func (core *Core) Serve(){
	  log.Fatal(http.ListenAndServe(":8080", nil))
}

func (core *Core) WriteToDB(input []byte) {

}


package core

import(
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/TonyFeng628/http-server/internal/dbm"
	"io/ioutil"
	"log"
)

func (core *Core)version(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "v0.0.1!\n")
}

func (core *Core)writeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "v0.0.1!\n")
}

func (core *Core)createHandler(w http.ResponseWriter,r *http.Request) {
	var input dbm.ContactInfo
	reqBody, err := ioutil.ReadAll(r.Body)
    if err != nil {
        log.Fatal(err)
    }
	err = json.Unmarshal(reqBody,&input)
	if err != nil {
		w.Write([]byte("error")) 
		fmt.Println(err)
		return
	}
	core.dbm.Write(input)
	w.Write([]byte("success"))
}
func (core *Core)readHandler(w http.ResponseWriter,r *http.Request){
	var input dbm.ContactInfo
	reqBody, err := ioutil.ReadAll(r.Body)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%s", reqBody)

	if err == nil {
		json.Unmarshal(reqBody,&input)
	}
	rets := core.dbm.Read(input)
	b, _ := json.Marshal(rets)
	w.Write(b)
}
func (core *Core)updateHandler(w http.ResponseWriter,r *http.Request){
	var input dbm.ContactInfo
	reqBody, err := ioutil.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Update not found", 404)
        return
    }
   err = json.Unmarshal(reqBody,&input)
   if err != nil {
   		log.Println(err)
        http.Error(w, "Update not found", 404)
        return
    }

	err = core.dbm.Update(input)
	if err != nil {
        http.Error(w, "Update not found", 404)
        return
    }

	w.Write([]byte("success"))
}
func (core *Core)deleteHandler(w http.ResponseWriter,r *http.Request){
	var input dbm.ContactInfo
	reqBody, err := ioutil.ReadAll(r.Body)
    if err != nil {
        log.Fatalln(err)
    }
    json.Unmarshal(reqBody,&input) 
	core.dbm.Delete(input)
	b, _ := json.Marshal(input)
	w.Write(b)
}


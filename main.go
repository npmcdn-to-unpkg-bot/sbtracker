package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/julienschmidt/httprouter"
	"log"
)

// Make this configurable via CLI or something
const databaseName = "sandbox_tracker.db"
const serverName = "localhost:8080" 


type Sandbox struct {
	Id int
	Url string
	Owner string
	Branch string
}

func (s *Sandbox) save() error {
	filename := fmt.Sprintf("%02d.txt", s.Id)
	return ioutil.WriteFile(filename, []byte(s.Owner), 0600)
}

func loadAll(db *sql.DB) ([]*Sandbox) {
	rows, err := db.Query("SELECT * FROM sandboxes")
	if err != nil {
		panic(err)
		return nil
	}

	var sandboxes [100]*Sandbox
	i := 0

	for rows.Next() {
		var id int
		var url string
		var owner string
		var branch string
		
		err = rows.Scan(&id, &url, &owner, &branch)
		if err != nil {
			panic(err)
			return nil
		}

		sandboxes[i] = &Sandbox{
			Id: id,
			Url: url,
			Owner: owner,
			Branch: branch}
		i = i+1
	}

	return sandboxes[0:i-1];
}

func dbOpen() (*sql.DB) {
	var db *sql.DB
	db, err := sql.Open("sqlite3", databaseName)
	if err != nil {
		panic(err)
	}

	return db
}

func List(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("loading sandboxes")
	db := dbOpen()
	defer db.Close()
	sandboxes := loadAll(db)
	json, _ := json.Marshal(sandboxes)

	w.Header().Set("Access-Control-Allow-Origin", serverName)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(json))
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t, _ := template.ParseFiles("web/index.html")
	db := dbOpen()
	defer db.Close()
	sandboxes := loadAll(db)

	err := t.Execute(w, sandboxes)
	if err != nil {
		panic(err);
	}
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.ServeFiles("/assets/*filepath", http.Dir("web"))
	router.GET("/sandbox/list.json", List)

	log.Fatal(http.ListenAndServe(":8080", router))
}

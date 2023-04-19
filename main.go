package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/gorilla/mux"

	"time"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func InsertToDB(key string, url string) string {

	// DBConnect

	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/indexer")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println("Unable to ping with DB")
	}

	// Retrieved from DB

	stmt, err := db.Prepare("INSERT INTO urls (url_key,url) VALUES(? , ? )")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(key, url)
	if err != nil {
		log.Fatal(err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Database log :: Operation = INSERT , ID = %d, affected = %d \n", lastId, rowCnt)
	return key
}

func RetriveUrlFromDB(key string) string {

	// DBConnect

	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/indexer")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println("Unable to ping with DB")
	}

	var name string
	err = db.QueryRow("select url from urls where url_key = ?", key).Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	return name
}

func ShortenUrl(w http.ResponseWriter, r *http.Request) {

	const string_lenght = 8
	randomString := RandStringRunes(string_lenght)

	Shrinkreq := ShrinkRequest{}

	err := json.NewDecoder(r.Body).Decode(&Shrinkreq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	InsertToDB(randomString, Shrinkreq.Url)

	Shrinkres := ShrinkResponse{}
	Shrinkres.Key = randomString
	Shrinkres.Url = Shrinkreq.Url

	shrinkJson, err := json.Marshal(Shrinkres)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(shrinkJson)

}

func ResolveUrl(w http.ResponseWriter, r *http.Request) {
	Resolreq := ResolveRequest{}
	err := json.NewDecoder(r.Body).Decode(&Resolreq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	url := RetriveUrlFromDB(Resolreq.Key)
	Resolresp := ResolveResponse{}
	Resolresp.Key = Resolreq.Key
	Resolresp.Url = url
	fmt.Printf("Resolving the url %s%s ==> %s \n ", localUrlPrefix, Resolresp.Key, Resolresp.Url)

	resolveJson, err := json.Marshal(Resolresp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resolveJson)

}

type ResolveRequest struct {
	Key string `json:"Key"`
}
type ResolveResponse struct {
	Key string `json:"Key"`
	Url string `json:"Url"`
}
type ShrinkRequest struct {
	Url string `json:"Url"`
}
type ShrinkResponse struct {
	Url string `json:"Url"`
	Key string `json:"Key"`
}

const localUrlPrefix = "http://shrink.io/"

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/shrink/", ShortenUrl).Methods("POST")
	r.HandleFunc("/resolve/", ResolveUrl).Methods("GET")
	http.ListenAndServe(":9908", r)
}

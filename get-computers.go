package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type comp struct {
	ComputerName string `json:"Computer_Name"`
	Username     string `json:"Username"`
}

type reports struct {
	Computers []comp `json:"computer_reports"`
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: get-computers <sqlite db>")
		return
	}
	db, err := sql.Open("sqlite3", os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	log.Println("Getting new list of computers...")
	req, err := http.NewRequest("GET", "https://jssi-web-005-p.cisco.com:8443/mdm/JSSResource/computerreports/id/724", nil)
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{}
	req.SetBasicAuth("dlp_casper_api.gen", "J&#45kklmn")
	req.Header.Set("Accept", "application/json")
	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Parsing response...")
	var r reports
	if err = json.Unmarshal(data, &r); err != nil {
		log.Fatal(err)
	}
	//log.Println("Adding entries to database...")
	//tx, err := db.Begin()
	//if err != nil {
	//	log.Fatal(err)
	//}
       f, err := os.Create("computers")
       if err != nil {
           log.Fatal(err)
       }
       defer f.Close()
for _, comp := range r.Computers {
                //_, err := stmt.Exec(comp.Username, comp.ComputerName)
                log.Println(comp.Username + "," + comp.ComputerName)
                f.WriteString(comp.ComputerName +"\n")
        }
        f.Sync()

	/*stmt, err := tx.Prepare("insert into computers(username, computer_name) values (?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	for _, comp := range r.Computers {
		_, err := stmt.Exec(comp.Username, comp.ComputerName)
		log.Println(comp.Username + "," + comp.ComputerName)
		if err != nil {
			log.Println(err)
		}
	}
	if err = tx.Commit(); err != nil {
		log.Fatal(err)
	}*/
}

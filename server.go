package main

import(
  "net/http"
  "html/template"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  _ "fmt"
   "log"
)

func main() {
    http.HandleFunc("/yatani", yatani)
    http.ListenAndServe(":8000", nil)
}

func yatani(w http.ResponseWriter, r *http.Request) {
    db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/seweb")
    		if err != nil {
			panic(err.Error())
		}
    defer db.Close()  

    rows, err := db.Query("SELECT body FROM articles WHERE id=1")
                if err != nil {
                        panic(err.Error())
                }
    defer rows.Close()

    var body string
    rows.Next()
    rows.Scan(&body)
// log.Print(body)
    t , err := template.ParseFiles("yatani.tpl")
                if err != nil {
                        panic(err.Error())
                }
    t.Execute(w, body)
}

// create database seweb;
// create table articles (id int, body varchar(255));
// insert into articles values(1,'矢谷のページ');


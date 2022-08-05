package main

import(
  "net/http"
  _ "github.com/go-sql-driver/mysql"
  _ "fmt"
  _ "log"
)

func yataniHandler (w http.ResponseWriter, r *http.Request) {
    rows, err := DB.Query("SELECT body FROM articles WHERE id=1")
                if err != nil {
                        panic(err.Error())
                }
    defer rows.Close()

    var body string
    rows.Next()
    err = rows.Scan(&body)
                if err != nil {
                        panic(err.Error())
                }
    err = HTMLTemplates.ExecuteTemplate(w, "yatani.tpl", body)
                if err != nil {
                        panic(err.Error())
                }
}




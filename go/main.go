package main

import(
  "net/http"
  "html/template"
  "database/sql"
  "os"
  _ "github.com/go-sql-driver/mysql"
  _ "log"
)

var DB *sql.DB
var HTMLTemplates *template.Template

func init() {
}

func main() {
  // DB接続
    DBConnect()
  // Template読み込み
    HTMLTemplates = template.Must(template.ParseGlob("templates/*tpl"))
  // WEBサーバ起動
    startWebServer()
}

func DBConnect() {
  // DB接続用文字列作成
    connectString := os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + 
                    os.Getenv("DB_HOST") + ":" + os.Getenv("DOCKER_DB_PORT") + ")/" + os.Getenv("DB_DATABASE")
                    var err error
  // 接続
    DB, err = sql.Open("mysql", connectString)
                if err != nil {
                        panic(err.Error())
                }
}

func startWebServer(){
    http.HandleFunc("/yichikawa", yichikawaHandler)
    http.HandleFunc("/myamagata", myamagataHandler)
    http.HandleFunc("/tozono", tozonoHandler)
    http.HandleFunc("/syatani", syataniHandler)
    http.HandleFunc("/skakeda", skakedaHandler)
    http.ListenAndServe(":8080", nil)
}

func Must(num int, err error) int {
                if err != nil {
                        panic(err.Error())
                }
  return num
}


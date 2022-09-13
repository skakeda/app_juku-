package main

import(
  "net/http"
  "html/template"
  "database/sql"
  "strconv"
  "time"
  "os"
  _ "github.com/go-sql-driver/mysql"
  _ "fmt"
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
    connectString := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + 
                    os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME")
    var err error
  // 接続
    DB, err = sql.Open("mysql", connectString)
                if err != nil {
                        panic(err.Error())
                }
  // 環境変数の値を設定
    if os.Getenv("DB_MAXCONN") != "" {
        DB.SetMaxOpenConns(Must(strconv.Atoi(os.Getenv("DB_MAXCONN"))))
    }
    if os.Getenv("DB_MAXIDLECONN") != "" {
        DB.SetMaxIdleConns(Must(strconv.Atoi(os.Getenv("DB_MAXIDLECONN"))))
    }
    if os.Getenv("DB_MAXLIFEMINUT") != "" {
        DB.SetConnMaxLifetime(time.Duration(Must(strconv.Atoi(os.Getenv("DB_MAXLIFEMINUTE")))) *time.Minute)
    }
    // if os.Getenv("DB_MAXIDLETIME") != "" {
    //   DB.SetConnMaxIdleTime(time.Duration(Must(strconv.Atoi(os.Getenv("DB_MAXIDLETIME")))) *time.Minute)
    // }
}

func startWebServer(){
    http.HandleFunc("/yatani", yataniHandler)
    http.ListenAndServe(":8080", nil)
}

func Must(num int, err error) int {
                if err != nil {
                        panic(err.Error())
                }
  return num
}

// create database seweb;
// create table articles (id int, body varchar(255));
// insert into articles values(1,'矢谷のページ');



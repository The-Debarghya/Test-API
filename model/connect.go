package model

import (
  "database/sql"
  "log"
  "fmt"
)
var con *sql.DB

func Connect() *sql.DB  {
  db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/mysql")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("Connection established with the database!!")
  con = db
  return db

}

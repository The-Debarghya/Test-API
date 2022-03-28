package main
import (
  "net/http"
  "controller"
  "model"
  "log"
  _"github.com/go-sql-driver/mysql"
  "fmt"
)

func main() {
  mux := controller.Register()
  db := model.Connect()
  defer db.Close()
  fmt.Println("Server Started...")
  log.Fatal(http.ListenAndServe(":3000", mux))
}

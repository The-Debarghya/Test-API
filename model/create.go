package model

func CreateTodo() error {
  insertQ, err := con.Query("INSERT TODO VALUES(?, ?)", "Debarghya", "First Server")
  defer insertQ.Close()
  if err != nil {
    return err
  }
  return nil

}

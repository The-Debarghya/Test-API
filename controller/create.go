package controller

import (
  "net/http"
  "model"
)

func create() http.HandlerFunc  {
  return func(w http.ResponseWriter, r *http.Request){
    if r.Method == http.MethodGet{
      //take data from user and save it
      if err := model.CreateTodo(); err != nil{
        w.Write([]byte("Some error occured midway!!"))
        return
      }
    }
  }

}

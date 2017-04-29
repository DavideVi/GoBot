package main

import (
  "fmt"
  "io"
  "net/http"
  "io/ioutil"
  "encoding/json"
)
// "fmt"
// "github.com/gorilla/mux"

type UserMessage struct {
  User    string  `json:"user"`
  Message string  `json:"message"`
}

func ProcessMessage(w http.ResponseWriter, r *http.Request) {

  // Holds the body of the message in JSON
  var message UserMessage

  // Reading the body
  body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
  if err != nil {
    panic(err)
  }
  if err := r.Body.Close(); err != nil {
    panic(err)
  }
  if err := json.Unmarshal(body, &message); err != nil {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(422) // unprocessable entity
    if err := json.NewEncoder(w).Encode(err); err != nil {
      panic(err)
    }
  }

  fmt.Println(message)

  // Going through command and checking if the regex matches
  // for i := range commands.Commands {
  //   fmt.Println(commands)
  //
  // }

}


//
// func TodoIndex(w http.ResponseWriter, r *http.Request) {
//     // todos := Todos{
//     //     Todo{Name: "Write presentation"},
//     //     Todo{Name: "Host meetup"},
//     // }
//     //
//     // if err := json.NewEncoder(w).Encode(todos); err != nil {
//     //     panic(err)
//     // }
// }
//
// func TodoShow(w http.ResponseWriter, r *http.Request) {
//     // vars := mux.Vars(r)
//     // todoId := vars["todoId"]
//     // fmt.Fprintln(w, "Todo show:", todoId)
// }

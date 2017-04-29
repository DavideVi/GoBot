package main

import (
    "os"
    "log"
    "net/http"
)

func main() {

  // TODO: Make path to commands.json an environment variable
  LoadCommands(os.Getenv("GOPATH") + "/bin/commands.json")

  // Creating router and listening
  // TODO: Make port an environment variable
  router := NewRouter()
  log.Fatal(http.ListenAndServe(":8080", router))
}

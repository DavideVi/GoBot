package main

import (
    "io/ioutil"
    "os"
    "fmt"
    "encoding/json"
)

type Command struct {
  Regex     string  `json:"regex"`
  Script    string  `json:"script"`
  Executor  string  `json:"executor"`
}

type Commands struct {
  Commands  []Command  `json:"commands"`
}

var commands Commands

func LoadCommands(filepath string) {

  // Commands file must exist
  if _, err := os.Stat(filepath); os.IsNotExist(err) {
    fmt.Printf("The file '%v' cannot be found. \nApplication will now exit\n", filepath)
    os.Exit(1)
  }

  dat, err := ioutil.ReadFile(filepath)
  if err != nil {
    panic(err)
  }
  fmt.Print(string(dat))

  // Turning it into JSON
  json.Unmarshal(dat, &commands)
  // fmt.Printf("Results: %v\n", &commands)
  //
  // for i := range commands.Commands {
  //       fmt.Println("\nRegex: " + commands.Commands[i].Regex)
  //       fmt.Println("Script: " + commands.Commands[i].Script)
  //       fmt.Println("Executor: " + commands.Commands[i].Executor + "\n")
  // }

}

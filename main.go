package main

import (
  "fmt"
  "os"
  "gopkg.in/alecthomas/kingpin.v2"
  "./keychain"
)

var (
  name = kingpin.Arg("name", "Name of note").Required().String()
)

func main() {
  kingpin.Parse()

  key := *name

  note, err := keychain.GetNote(key)

  if (err != nil) {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }

  fmt.Println(note)
}
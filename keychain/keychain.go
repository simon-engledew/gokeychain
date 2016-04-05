package keychain

import (
  "encoding/hex"
  "github.com/DHowett/go-plist"
  "os/exec"
  "strings"
  "bytes"
  "errors"
)

type Note struct {
    Text string `plist:"NOTE"`
}

func FirstString(args ...string) (string) {
  for _, item := range(args) {
    if (item != "") {
      return item
    }
  }

  return ""
}

func FindGenericPassword(key string) ([]byte, error) {
  command := exec.Command("security", "find-generic-password", "-l", key, "-w")

  out, err := command.Output()

  if err != nil {
    return nil, errors.New(FirstString(string(err.(*exec.ExitError).Stderr), "Access denied"))
  }

  out, err = hex.DecodeString(strings.TrimSuffix(string(out), "\n"))

  if err != nil {
    return nil, err
  }

  return out, nil
}

func GetNote(key string) (string, error) {
  out, err := FindGenericPassword(key)

  if err != nil {
    return "", err
  }

  var note Note

  decoder := plist.NewDecoder(bytes.NewReader(out))

  err = decoder.Decode(&note)

  if err != nil {
    return "", err
  }

  return note.Text, nil
}
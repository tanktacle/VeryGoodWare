// audio_stegano - a tool to hide files in WAV audio files writen in golang
// this script was an absolute nightmare to do thank you very much
package main

import (
  "io/ioutil"
  "strings"
  "log"
  "fmt"
  "os"
)

func secretMessageDummy(mess string, wav []byte) []byte {

  // Append dummy data to fill out the rest of the bytes
  var messLen int
  messLen = len(wav) - len(mess)
  mess = mess + strings.Repeat("#", int(messLen/8)-(len(mess)*8)  )
  return []byte(mess)
}

func main() {

  var secretMessage string
  secretMessage = "Just a random string"

  // Read audio file and convert it to bytearray
  wavArray, err := ioutil.ReadFile("sample.wav")
  if err != nil {
    log.Fatal(err)
  }
  // Convert the secretMessage to its binary representation
  secretMessageBytes := secretMessageDummy(secretMessage, wavArray)

  fmt.Printf("SecretMsgByteLength: %d, wavArray length: %d\n",len(secretMessageBytes),len(wavArray))

  if len(secretMessageBytes) > len(wavArray) {
    fmt.Println("The secret message is too long to be encyrypted into the supplied wav file!")
    os.Exit(-1)
  }

  var wavIndex int = 0

  for _,bte := range secretMessageBytes {

    secretMsgBits := bits([]byte{bte}) // convert 1 character of secretMessage to an array of 0's and 1's
    for _,sbit := range secretMsgBits {

      wavArray[wavIndex] = (wavArray[wavIndex] & byte(254+sbit))
      wavIndex++

      if wavIndex == len(wavArray) {
        fmt.Println("Overflow! Wrapping over")
        wavIndex = 0
      }
    }
  }
  // Replace the LSB of each byte of the audio bit by bit from the text bit array
  //for i,bit := range []byte(secretMessage) {
    // Perform ANDing with the LSB of the wav file and store it in a new array
    // Perform ORing with the resulting array with the secret message
    //wavArray[i] = (wavArray[i] & 254) | bit

  //}

  // Write the byte array to a new wav file
  ioutil.WriteFile("hidden.wav", wavArray, 0644)
}

func bits(bs []byte) []int {
    r := make([]int, len(bs)*8)
    for i, b := range bs {
        for j := 0; j < 8; j++ {
            r[i*8+j] = int(b >> uint(7-j) & 0x01)
        }
    }
    return r
}

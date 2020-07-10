// audio_stegano - a tool to hide files in WAV audio files writen in golang
// this script was an absolute nightmare to do thank you very much

package main

import (
  "io/ioutil"
  "log"
  "fmt"
  "strconv"
)

func main() {

  secretWav, err := ioutil.ReadFile("hidden.wav")
  if err != nil {
    log.Fatal(err)
  }

  var byteResult []byte

  var b []byte
  for _,bte := range secretWav {
    //secretMsgBits := bits([]byte{bte})
    b = append(b,getLSB(bte))

    if len(b) == 8 {
      byteStr := ""
      for _,n := range b {
        byteStr += fmt.Sprintf("%d",n)
      }

      intResult := bin2int(byteStr)

      //fmt.Printf("%s => %d\n",byteStr,intResult)
      byteResult = append(byteResult,byte(intResult))

      b = []byte{}
    }

    //for _,sbit := range secretMsgBits {

      // Extract the LSB of each byte

    //}
    // Convert byte array back to string

  }
  ioutil.WriteFile("unhidden.txt", byteResult, 0644)
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

func getLSB(b byte) byte {
        if b%2 == 0 {
                return 0
        }
        return 1
}


func bin2int(binStr string) int {
  result, _ := strconv.ParseInt(binStr, 2, 64)
  return int(result)
}

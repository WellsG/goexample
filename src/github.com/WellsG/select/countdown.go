package main

import (
  "os"
  "fmt"
  "time"
)

func main(){
  abort := make(chan struct{})
  go func(){
    os.Stdin.Read(make([]byte, 1))  // read a single byte
    abort <- struct{}{}
  }()

  fmt.Println("Commencing countdown,  Press return to abort.")
  
  tick := time.Tick(1*time.Second)
  for countdown := 10; countdown > 0; countdown-- {
    fmt.Println(countdown)
    select {
      case <-tick:  //time.After(10 * time.Second):
       // Do Nothing
      case <-abort:
       fmt.Println("Launch aborted")
       return
    }
  }

  launch()
}

func launch(){
  fmt.Println("Launch .")
}
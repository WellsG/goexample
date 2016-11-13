package main

import (
  "sync"
  "fmt"
  "time"
)

var (
  mu  sync.Mutex // guards balance
  balance int
)

func Deposit(amount int) {
  mu.Lock()
  balance += amount
  mu.Unlock()
}

func Balance() int {
  mu.Lock()
  b := balance
  mu.Unlock()
  return b
}

func main(){
  go Deposit(200)
  go Balance()
  go Deposit(300)
  go Balance()

  time.Sleep(2 * time.Second)

  fmt.Println(Balance())
}

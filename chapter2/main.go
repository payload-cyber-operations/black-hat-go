package main


import (
  "fmt"
//  "net"
  "sync"
//  "sort"
)


func workers(ports chan int, wg *sync.WaitGroup){
  for p := range ports{
    fmt.Println(p)
    wg.Done()
  }
}

func main(){
  var wg sync.WaitGroup
  ports := make(chan int, 100)
  for x:= 0; x < cap(ports); x++{
    go workers(ports, &wg)
  }

  for x:=0; x<=1024;x++{
    wg.Add(1)
    ports <- x
  }

  wg.Wait()
  close(ports)
}

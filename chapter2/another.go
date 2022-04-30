package main


// this is only a simple tcp scanner
// nothing interesting

import (
  "fmt"
  "net"
  "sync"
)


func conncurrentScanner(host string){
}

func main(){

  var wg sync.WaitGroup
  for i:=0; i<= 1024;i++{
    
    wg.Add(1)
    go func(j int){
      defer wg.Done()
      host := fmt.Sprintf("scanme.nmap.org:%d", j)
      conn, err := net.Dial("tcp", host)
      if err != nil {
        // connection refused or port closed
        return
      }

      conn.Close()
      fmt.Printf("\t\t Port %d open\n", j)
 
    }(i)
  }
  wg.Wait()
}



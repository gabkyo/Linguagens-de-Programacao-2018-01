package main

import(
  "fmt"
)

func main(){

  v:= New(64,"teste",triang,'A',2.01,5,0,4)
  fmt.Printf("Proprietario %v e id %v.\n",v.Owner(),v.Id())
}

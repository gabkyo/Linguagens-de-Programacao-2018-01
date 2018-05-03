package main

import(
  "fmt"
)

func main(){
  var v Imoveis = NewT(64,"teste",triang,'A',2.01,5,0,4)
  c,e:=v.Preco()
  fmt.Printf("%v %v\n",c,e)
}

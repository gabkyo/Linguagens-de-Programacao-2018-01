package main

import(
"fmt"
)

func main(){
	v:=NewBasicList()
	e:=ler_catalogo(v)
	fmt.Printf("%v",e)
	v.Print()
}

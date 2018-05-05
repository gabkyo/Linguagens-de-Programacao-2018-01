package main

import(
//"fmt"
)

func main(){
	v:=NewBasicList()
	e:=ler_catalogo(v)
	check(e)
	e=ler_atual(v)
	check(e)
	e=ler_espec(v)
	check(e)
}

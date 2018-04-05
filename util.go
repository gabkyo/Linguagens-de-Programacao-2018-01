package main

import (
		"bufio"
    "fmt"
    "io"
    "io/ioutil"
    "os"
		"string"
)
//Strings que possivelmente vou usar toda hora  como nomes de arquivo
const catalogo string = "catalogo.txt"
const atual string = "atual.txt"
const espec string = "espec.txt"
const saida string = "saída.txt"
const result string = "result.txt"
// e tipos de imoveis para comparacao e declaracao
const triang string = "triang"
const retang string = "retang"
const trapez string = "trapez"
const casa string = "casa"
const apto string = "apto"


type Imoveis interface{
	Preco() float32
	Id() uint64
	Owner() string
	Categoria() string
}
//////////////////////////////////

type Lote struct{
	id uint64
	owner string
	categoria string
}
func (l *Lote) Id() uint64{
	return l.id
}
func (l *Lote) Owner() string{
	return owner
}
func (l *Lote) Categoria() {
	return l.categoria
}
/////////////////////////////////

type Terreno struct{
	*Lote //embbebed tem que ser inicailizado
	solo int32
	precom2 float32
	base1, base2, altura float32
}

func InicializarTerreno(t *Terreno){
	t.Lote:=&Lote
}

func (t *Terreno) Area() float32{
	if(t.categoria == triang ){
		return t.base1*t.altura/2
	}else if(t.categoria == retang){
		return t.base1*t.altura
	}else if(t.categoria == trapez){
		return t.altura*(t.base1+t.base2)/2
	}else return 0
}

func (t *Terreno) Preco()float32{
}
////////////////////////////////


////////////////////////////

type Residencia struct{
	*Lote
	nQuartos int //numero de quartos
	nVagas int //vagas de garagem
}

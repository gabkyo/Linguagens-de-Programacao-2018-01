package main

import (
	"errors"
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
	Preco() (float32,error)
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
	return l.owner
}
func (l *Lote) Categoria() string{
	return l.categoria
}
/////////////////////////////////

type Terreno struct{
	*Lote //embbebed tem que ser inicailizado
	solo int32
	preco_m2 float32 //preco por m^2
	base1, base2, altura float32
}

func New(id uint64, owner string,categoria string,solo int32, preco_m2 float32, base1 float32, base2 float32, altura float32) *Terreno{
	Terreno := &Terreno{
		Lote : &Lote{
			id : id,
			owner : owner,
			categoria : categoria,
		},
		solo : solo,
		preco_m2 : preco_m2,
		base1 : base1,
		base2 : base2,
		altura : altura,
	}
	return Terreno
}

func (t *Terreno) Area() (float32,error){
	if(t.categoria == triang ){
		return t.base1*t.altura/2,nil
	}else if(t.categoria == retang){
		return t.base1*t.altura,nil
	}else if(t.categoria == trapez){
		return t.altura*(t.base1+t.base2)/2,nil
	}
	return 0,errors.New("Tipo invalido para "+string(t.id)+" ao calcular Area.\n")
}

func (t *Terreno) Preco()(float32,error){
	v:=t.Area() * t.preco_m2
	if(t.solo=='A'){
		return v*0.9,nil
	}else if(t.solo=='G'){
		return v*1.3,nil
	}else if(t.solo == 'R'){
		return v*1.1,nil
	}
	return 0,errors.New("Solo invalido para "+string(t.id)+" ao calcular Preco.\n")
}
////////////////////////////////

type Casa struct{
	*Lote
	n_Quartos uint //numero de quartos
	n_Vagas uint //vagas de garagem
	n_pavimentos uint //n de pavimentos
	area_Pavimento float32
	preco_m2_Pavimento float32
	area_Livre float32
	preco_m2_Area_Livre float32
	preco_m2_Area_Construida float32
}

func(c *Casa) Preco()(float32,error){
	return c.preco_m2_Area_Construida * c.area_Pavimento * c.n_pavimentos + c.preco_m2_Area_Livre* c.area_Livre,nil
}

////////////////////////////

type Apartamento struct{
	*Lote
	n_Quartos uint //numero de quartos
	n_Vagas uint //vagas de garagem
	andar uint
	area_Construida float32
	preco_m2_Area_Construida float32
	lazer int32 //S ou N
	n_Andares uint
}

func (a *Apartamento) Preco()(float32,error){
	v:=a.preco_m2_Area_Construida *a.area_Construida * (0.9 + float32(a.andar/a.n_Andares))
	if(a.lazer=='S'){
		return v*1.15,nil
	}else if(a.lazer=='N'){
		return v,nil
	}
	return v,errors.New("Lazer invalido para "+string(a.id)+" ao calcular Preco.\n")
}

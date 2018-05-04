package main

import (
"bufio"
"os"
"strconv"
"errors"
)

//scanner scan le a linha text passa pra string

func addCasa(scanner *bufio.Scanner,lista *BasicList,owner string,id uint64) error{
    var temp Imoveis
    var n_Quartos, n_Vagas, n_pavimentos uint
    var u uint64
    var area_Pavimento, preco_m2_Pavimento, area_Livre, preco_m2_Area_Livre float32
    var f float64
    var s string
    var err error

    scanner.Scan() //le uma linha
    s=scanner.Text() //passa linha para string
    u,err = strconv.ParseUint(s,10,32) //string para uint64
    check(err) //se tiver erro, panico
    n_Quartos=uint(u)//uint64 para uint

    scanner.Scan()
    s=scanner.Text()
    u,err = strconv.ParseUint(s,10,32)
    check(err)
    n_Vagas=uint(u)

    scanner.Scan()
    s=scanner.Text()
    u,err = strconv.ParseUint(s,10,32)
    check(err)
    n_pavimentos=uint(u)

    scanner.Scan()
    s=scanner.Text()
    f,err = strconv.ParseFloat(s,32)
    check(err)
    area_Pavimento=float32(f)

    scanner.Scan()
    s=scanner.Text()
    f,err = strconv.ParseFloat(s,32)
    check(err)
    preco_m2_Pavimento=float32(f)

    scanner.Scan()
    s=scanner.Text()
    f,err = strconv.ParseFloat(s,32)
    check(err)
    area_Livre=float32(f)

    scanner.Scan()
    s=scanner.Text()
    f,err = strconv.ParseFloat(s,32)
    check(err)
    preco_m2_Area_Livre=float32(f)

    temp=NewC(id,owner,n_Quartos,n_Vagas,n_pavimentos,area_Pavimento,preco_m2_Pavimento,area_Livre,preco_m2_Area_Livre)
    lista.Insert(temp)
    return nil
}

func addTriang(scanner *bufio.Scanner,lista *BasicList,owner string,id uint64) error{
   var temp Imoveis
   var f float64
   var s string
   var err error
   var solo int32
   var preco_m2 ,base1, altura float32

   scanner.Scan()
   s=scanner.Text()
   solo = int32(s[0])

   scanner.Scan()
   s=scanner.Text()
   f,err = strconv.ParseFloat(s,32)
   check(err)
   preco_m2=float32(f)

   scanner.Scan()
   s=scanner.Text()
   f,err = strconv.ParseFloat(s,32)
   check(err)
   base1=float32(f)

   scanner.Scan()
   s=scanner.Text()
   f,err = strconv.ParseFloat(s,32)
   check(err)
   altura=float32(f)

   temp=NewT(id,owner,triang,solo,preco_m2,base1,0,altura)
   lista.Insert(temp)
   return nil
}

func addRetang(scanner *bufio.Scanner,lista *BasicList,owner string,id uint64) error{
   var temp Imoveis
   var f float64
   var s string
   var err error
   var solo int32
   var preco_m2 ,base1, altura float32

   scanner.Scan()
   s=scanner.Text()
   solo = int32(s[0])

   scanner.Scan()
   s=scanner.Text()
   f,err = strconv.ParseFloat(s,32)
   check(err)
   preco_m2=float32(f)

   scanner.Scan()
   s=scanner.Text()
   f,err = strconv.ParseFloat(s,32)
   check(err)
   base1=float32(f)

   scanner.Scan()
   s=scanner.Text()
   f,err = strconv.ParseFloat(s,32)
   check(err)
   altura=float32(f)

   temp=NewT(id,owner,retang,solo,preco_m2,base1,0,altura)
   lista.Insert(temp)
   return nil
}

func addTrapez(scanner *bufio.Scanner,lista *BasicList,owner string,id uint64) error{
 var temp Imoveis
 var f float64
 var s string
 var solo int32
 var preco_m2 ,base1, base2, altura float32 
 var err error

 scanner.Scan()
 s=scanner.Text()
 solo = int32(s[0])

 scanner.Scan()
 s=scanner.Text()
 f,err = strconv.ParseFloat(s,32)
 check(err)
 preco_m2=float32(f)

 scanner.Scan()
 s=scanner.Text()
 f,err = strconv.ParseFloat(s,32)
 check(err)
 base1=float32(f)

 scanner.Scan()
 s=scanner.Text()
 f,err = strconv.ParseFloat(s,32)
 check(err)
 base2=float32(f)

 scanner.Scan()
 s=scanner.Text()
 f,err = strconv.ParseFloat(s,32)
 check(err)
 altura=float32(f)

 temp=NewT(id,owner,triang,solo,preco_m2,base1,base2,altura)
 lista.Insert(temp)
 return nil
}

func addApto(scanner *bufio.Scanner,lista *BasicList,owner string,id uint64) error{
   var temp Imoveis
   var u uint64
   var f float64
   var s string
   var n_Quartos, n_Vagas, andar, n_Andares uint 
   var area_Construida, preco_m2_Area_Construida float32
    var lazer int32 //S ou N
    var err error

    scanner.Scan()
    s=scanner.Text()
    u,err = strconv.ParseUint(s,10,32) // bloco
    if err!=nil {                              // bloco 
        return err                             // bloco                       
    }
    n_Quartos=uint(u)
                                              // bloco     
    scanner.Scan()
    s=scanner.Text()
    u,err = strconv.ParseUint(s,10,32)
    check(err)
    n_Vagas=uint(f)

    scanner.Scan()
    s=scanner.Text()
    u,err = strconv.ParseUint(s,10,32)
    check(err)
    andar=uint(f)

    scanner.Scan()
    s=scanner.Text()
    f,err = strconv.ParseFloat(s,32)
    check(err)
    area_Construida=float32(f)

    scanner.Scan()
    s=scanner.Text()
    f,err = strconv.ParseFloat(s,32)
    check(err)
    preco_m2_Area_Construida=float32(f)

    scanner.Scan()
    s=scanner.Text()
    lazer= int32(s[0])

    scanner.Scan()
    s=scanner.Text()
    u,err = strconv.ParseUint(s,10,32)
    check(err)
    n_Andares=uint(u)

    temp=NewA(id,owner,n_Quartos,n_Vagas,andar,area_Construida,preco_m2_Area_Construida,lazer,n_Andares)
    lista.Insert(temp)
    return nil
}

func ler_catalogo(lista *BasicList) error{
    file, err := os.Open(catalogo)
    var s, owner, categoria string
    var id uint64
    if err != nil {
        return err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        categoria=scanner.Text() 
        scanner.Scan()
        s=scanner.Text()
        id,err = strconv.ParseUint(s,10,32)
        check(err)
        scanner.Scan()
        owner=scanner.Text()
        switch categoria{
        case casa:
            addCasa(scanner,lista,owner,id)
        case triang:
            addTriang(scanner,lista,owner,id)
        case retang:
            addRetang(scanner,lista,owner,id)
        case trapez:
            addTrapez(scanner,lista,owner,id)
        case apto:
            addApto(scanner,lista,owner,id)
            default: 
            return errors.New("Categoia inválida "+categoria+".\n")
        }
        scanner.Scan() //pular linha em branco
    }

    if err := scanner.Err(); err != nil {
        return err
    }
    return nil
}

func ler_atual(lista *BasicList) error{
    file, err := os.Open(catalogo)
    var s, owner, categoria string
    var id uint64
    var temp Imoveis
    var u uint64
    var f float64 
    if err != nil {
        return err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        categoria=scanner.Text() 
        
        switch categoria{
        case "i":
            scanner.Scan()
            s=scanner.Text()
            id,err = strconv.ParseUint(s,10,32)
            check(err)
            scanner.Scan()
            owner=scanner.Text()

        case triang:

        case retang:

        case trapez:

        case apto:

            default: 
            return errors.New("Categoia inválida "+categoria+".\n")
        }
        scanner.Scan() //pular linha em branco
    }

    if err := scanner.Err(); err != nil {
        return err
    }
    return nil
}
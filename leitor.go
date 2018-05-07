package main

import (
"bufio"
"os"
"strconv"
"errors"
"fmt"
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


 temp=NewT(id,owner,trapez,solo,preco_m2,base1,base2,altura)
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
            return errors.New("Categoria inválida "+categoria+".\n")
        }
        scanner.Scan() //pular linha em branco
    }

    if err := scanner.Err(); err != nil {
        return err
    }
    return nil
}

func ler_atual(lista *BasicList) error{
    file, err := os.Open(atual)
    var s, owner, categoria, tipo string
    var id uint64

    if err != nil {
        return err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        tipo=scanner.Text()

        switch tipo{
        case "i":
            scanner.Scan()
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
                return errors.New("Categoria inválida "+categoria+".\n")
            }

        case "e":
            scanner.Scan()
            s=scanner.Text()
            id,err = strconv.ParseUint(s,10,32)
            check(err)

            err=lista.Remove(id)
            check(err)

        case "a":
            scanner.Scan()
            categoria=scanner.Text()
            scanner.Scan()
            s=scanner.Text()
            id,err = strconv.ParseUint(s,10,32)
            check(err)

            scanner.Scan()
            owner=scanner.Text()
            err=lista.Remove(id)
            check(err)

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
                return errors.New("Categoria inválida "+categoria+".\n")
            }
        default:
            return errors.New("Comando inválido "+categoria+".\n")
        }
        scanner.Scan() //pular linha em branco
    }

    if err := scanner.Err(); err != nil {
        return err
    }
    return nil
}


func saida_print(lista *BasicList, p_imoveis_caros uint64,p_menores_argilosos uint64, i uint64, j uint64, k uint64,area_limite float64,preco_limite float64){
    file,err:= os.Create(saida)
    limite:= uint(float32(lista.Length())* float32(p_imoveis_caros)/100) //malabarismo matematico
    contador:=lista.Length()
    marcador:=0
    currentNode := lista.head //nodo atual, lembrar de resetar posicao
    check(err)
    defer file.Close()
    w:=bufio.NewWriter(file)
    var soma uint64 //soma de i j k
    soma=0

    for contador >0 && currentNode !=nil{ //printando os terrenos de tras para frente pois estao do maior preco para menor
        if contador< limite {
            _,err=fmt.Fprintf(w,", ")
            check(err)
        }
        if contador <= limite {
            s:=strconv.FormatUint(currentNode.key.Id(),10)
            _,err=fmt.Fprintf(w,s)
            check(err)
            marcador++
            if i==uint64(marcador) {
                soma+=currentNode.key.Id()
            }
        }
        currentNode=currentNode.Next()
        contador-=1

    }

    _,err=fmt.Fprintf(w,"\n")
    check(err)
    contador=lista.Length()
    currentNode = lista.head
    var argilosos []uint64 // lista de terrenos argilosos
    for contador >0 && currentNode !=nil{
        if currentNode.key.Info() == "G"{
            argilosos=append(argilosos, currentNode.key.Id())
        }
        currentNode=currentNode.Next()
        contador-=1
    }

    contador=1
    for int(contador) < len(argilosos) { //sort
        a,_:=lista.Search(argilosos[contador-1])
        b,_:=lista.Search(argilosos[contador])
        aA,_:=a.key.Area()
        bA,_:=b.key.Area()
        if  aA<bA  { //se a area do anterior menor que do proximo ou seja (em ordem de area decrescente)
            argilosos[contador-1],argilosos[contador] = argilosos[contador],argilosos[contador-1]
            contador=0
        }
        contador+=1
    }

    limite= uint(float32(len(argilosos))* (1-float32(p_menores_argilosos)/100)) //malabarismo matematico 2
    contador=0
    currentNode = lista.head
    for contador <uint(len(argilosos)) && currentNode !=nil{
        if contador > limite+1 { //pq comeca print em limite+1 entao ',' a partir de limite +2
            _,err=fmt.Fprintf(w,", ")
            check(err)
        }
        if contador > limite {
            s:=strconv.FormatUint(argilosos[contador],10)
            _,err=fmt.Fprintf(w,s)
            check(err)
        }
        if(uint64(contador)+1==j){
            soma+=currentNode.key.Id()
        }
        currentNode=currentNode.Next()
        contador+=1
    }

    _,err=fmt.Fprintf(w,"\n")
    check(err)
    currentNode = lista.head
    var casas []uint64 //casas dentro dos parametros
    for currentNode!=nil { //cada casa com preco< precolimite e area pavimento > area limite
        if currentNode.key.Categoria()==casa{
            var area_Construida,preco float32
            area_Construida,err=currentNode.key.Area()
            check(err)
            preco,err=currentNode.key.Preco()
            check(err)
            //fmt.Printf("%d : %v %v / %v %v / %v\n",currentNode.key.Id(),area_Construida,area_limite,preco, preco_limite,currentNode.key.Info())
            if float64(area_Construida) > area_limite && float64(preco) < preco_limite{
                casas=append(casas,currentNode.key.Id())
            }
        }
        currentNode=currentNode.Next()
    }
    currentNode = lista.head
    contador=1
    for int(contador) < len(casas) { //sort
        a,_:=lista.Search(casas[contador-1])
        b,_:=lista.Search(casas[contador])
        s:=a.key.Info()
        var qa,qb uint64
        qa,err=strconv.ParseUint(s,10,64)
        check(err)
        s=b.key.Info()
        qb,err=strconv.ParseUint(s,10,64)
        check(err)
        if  qa<qb  { //se nquartos do anterior menor que do proximo ou seja (em ordem de area decrescente)
            casas[contador-1],casas[contador] = casas[contador],casas[contador-1]
            contador=0
        }else if qa==qb{
          if a.key.Id()<b.key.Id() {
            casas[contador-1],casas[contador] = casas[contador],casas[contador-1]
          }
        }
        contador+=1
    }
    contador=0
    for contador <uint(len(casas)){
        if contador>0 {
            fmt.Fprintf(w,", ")
        }
        fmt.Fprintf(w,"%d",casas[contador])
        if uint64(contador)+1==k {
            soma+=casas[contador]
        }
        contador++
    }
    w.Flush()
    file.Close()
    file,err=os.Create(result)
    check(err)
    w=bufio.NewWriter(file)
    fmt.Fprintf(w,"%d",soma) //print da soma
    w.Flush()
    file.Close()
}

func ler_espec(lista *BasicList) error{
    file, err := os.Open(espec)
    var p_imoveis_caros, p_menores_argilosos, i, j, k uint64
    var area_limite, preco_limite float64
    var s string

    if err != nil {
        return err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Scan()
    s=scanner.Text()
    p_imoveis_caros,err=strconv.ParseUint(s,10,64)
    check(err)

    scanner.Scan()
    s=scanner.Text()
    p_menores_argilosos,err=strconv.ParseUint(s,10,64)
    check(err)

    scanner.Scan()
    s=scanner.Text()
    area_limite,err=strconv.ParseFloat(s,64)
    check(err)

    scanner.Scan()
    s=scanner.Text()
    preco_limite,err=strconv.ParseFloat(s,64)
    check(err)

    scanner.Scan()
    s=scanner.Text()
    i,err=strconv.ParseUint(s,10,64)
    check(err)

    scanner.Scan()
    s=scanner.Text()
    j,err=strconv.ParseUint(s,10,64)
    check(err)

    scanner.Scan()
    s=scanner.Text()
    k,err=strconv.ParseUint(s,10,64)
    check(err)


    if err := scanner.Err(); err != nil {
        return err
    }
    saida_print(lista,p_imoveis_caros, p_menores_argilosos, i, j, k,area_limite, preco_limite)
    return nil
}

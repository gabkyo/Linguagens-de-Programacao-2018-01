from _class import *
from _list import *
#nao pude nomear esse arquivo _io.py


def ler_linha(file): #pra facilitar minha vida ,cada lerlinha retorna a proxima linha lida do arquivo
	linha=file.readline()
	linha=linha.rstrip()
	return linha


def addCasa(file,lista): #file tipo arquivo lista tipo Lista, cria e insere uma casa parametros: id,owner,quartos,vagas,pavimentos,area_pavimento,pm2pavimento,area_livre,pm2livre
	linha=ler_linha(file)
	id=int(linha)
	linha=ler_linha(file)
	owner=linha
	linha=ler_linha(file)
	quartos=int(linha)
	linha=ler_linha(file)
	vagas=int(linha)
	linha=ler_linha(file)
	pavimentos=int(linha)
	linha=ler_linha(file)
	area_pavimento=float(linha)
	linha=ler_linha(file)
	pm2pavimento=float(linha)
	linha=ler_linha(file)
	area_livre=float(linha)
	linha=ler_linha(file)
	pm2livre=float(linha)
	casa=Casa(id,owner,quartos,vagas,pavimentos,area_pavimento,pm2pavimento,area_livre,pm2livre)
	lista.Insert(casa)

def addApto(file,lista): #file tipo arquivo lista tipo Lista, cria e insere um apto parametros: id,owner,quartos,vagas,andar,area_construida,pm2construida,lazer,andares
	linha=ler_linha(file)
	id=int(linha)
	linha=ler_linha(file)
	owner=linha
	linha=ler_linha(file)
	quartos=int(linha)
	linha=ler_linha(file)
	vagas=int(linha)
	linha=ler_linha(file)
	andar=int(linha)
	linha=ler_linha(file)
	area_construida=float(linha)
	linha=ler_linha(file)
	pm2construida=float(linha)
	linha=ler_linha(file)
	lazer=linha
	linha=ler_linha(file)
	andares=int(linha)
	apto=Apto(id,owner,quartos,vagas,andar,area_construida,pm2construida,lazer,andares)
	lista.Insert(apto)

def addTerreno(file,lista,categoria): #file tipo arquivo lista tipo Lista, cria e insere um terreno parametros: id, owner, categoria, solo, precom2, base1 , base2, altura
	linha=ler_linha(file)
	id=int(linha)
	linha=ler_linha(file)
	owner=linha
	linha=ler_linha(file)
	solo=linha
	linha=ler_linha(file)
	precom2=float(linha)
	linha=ler_linha(file)
	base1=float(linha)
	linha=ler_linha(file)
	if categoria=="trapez":
		base2=float(linha)
		linha=ler_linha(file)
	else:
		base2=0
	altura=float(linha)
	terreno=Terreno(id, owner, categoria, solo, precom2, base1, base2, altura)
	lista.Insert(terreno)


def ler_catalogo(lista): #processa catalogo txt
	file=open("catalogo.txt",'r')
	linha=ler_linha(file) #categoria
	while linha:
		if linha=="apto":
			addApto(file,lista)
		elif linha=="casa":
			addCasa(file,lista)
		elif linha=="triang" or  linha=="trapez" or  linha=="retang": #se nao for casa ou apto, terreno ja que nao ha erros no input
			addTerreno(file,lista,linha)
		else: print("Erro:",linha)
		ler_linha(file) #pular linha vazia
		linha=ler_linha(file) #proxima linha
	file.close()

def ler_atual(lista):
	file=open("atual.txt",'r')
	linha=ler_linha(file)
	while linha:
		if linha=='i':
			linha=ler_linha(file)
			if linha=="apto":
				addApto(file,lista)
			elif linha=="casa":
				addCasa(file,lista)
			elif linha=="triang" or  linha=="trapez" or  linha=="retang": #se nao for casa ou apto, terreno ja que nao ha erros no input
				addTerreno(file,lista,linha)
			else: print("Erro:",linha)
		elif linha=='a':
			posicao=file.tell() #posicao antes de comecar a ler
			ler_linha(file) #pular categoria
			linha=ler_linha(file) #pegar o id
			id=int(linha)
			lista.Remove(id) #remover elemento antigo
			file.seek(posicao) #voltar pra antes de ler categoria
			linha=ler_linha(file)
			if linha=="apto":
				addApto(file,lista)
			elif linha=="casa":
				addCasa(file,lista)
			elif linha=="triang" or  linha=="trapez" or  linha=="retang": #se nao for casa ou apto, terreno ja que nao ha erros no input
				addTerreno(file,lista,linha)
			else: print("Erro:",linha)
		elif linha=='e': #pegar o id e remover
			linha=ler_linha(file)
			id=int(linha)
			lista.Remove(id)
		else: print("Erro:",linha)
		ler_linha(file) #pular linha vazia
		linha=ler_linha(file)
	file.close()

def criarArgilosos(lista,percentArgilosos): #fazer uma lista de argilosos

	atual=lista.head
	temp=LList()
	while atual!=None: #lista temporaria com todos os argilosos
		if atual.key.Info()=="G":
			temp.Insert_Argiloso(atual.key)
		atual=atual.next
		
	atual=temp.head
	contador=temp.Length()
	argilosos=LList()
	numeroArgilosos = int(percentArgilosos*temp.Length()/100) #numero do total que vai ser usado
	while atual!=None: #lista final so com o numero de argilosos necessario
		if contador<=numeroArgilosos:
			argilosos.Insert_Argiloso(atual.key)
		contador-=1
		atual=atual.next
	return argilosos

def criarCasas(lista, area_limite, preco_limite): #lista de casas que seguem criterios
	atual=lista.head
	casas=LList()
	while atual!=None:
		if atual.key.categoria=="casa" and atual.key.Preco() < preco_limite and atual.key.Area() > area_limite:
			casas.Insert_Casa(atual.key)
		atual=atual.next

	return casas

def criarImoveis(lista,percentImoveis): #n imoveis mais caros
	numeroImoveis = int(percentImoveis*lista.Length()/100)
	contador=lista.Length()
	atual=lista.head
	imoveis=LList()
	while atual!=None:
		if contador<=numeroImoveis:
			imoveis.Insert(atual.key)
		contador-=1
		atual=atual.next
	return imoveis
			


def ler_espec(lista): #le espec e gera as saidas
	soma=0 #soma de i j k esimos
	file=open("espec.txt",'r')
	linha=ler_linha(file) 

	percentImoveis=int(linha)	#% dos imoveis
	linha=ler_linha(file) 
	percentArgilosos=int(linha)	#% de imoveis argilosos
	linha=ler_linha(file) 
	area_limite=float(linha)	#area limite para 3 item
	linha=ler_linha(file) 
	preco_limite=float(linha)	#preco limite para 3 item
	linha=ler_linha(file) 
	i=int(linha)	#iesimo imovel mais caro
	linha=ler_linha(file) 
	j=int(linha)	#jesimo argiloso
	linha=ler_linha(file) 
	k=int(linha)	
	file.close()

	file=open("saida.txt",'w') #escrevendo 
	imoveis=criarImoveis(lista,percentImoveis) 
	argilosos=criarArgilosos(lista,percentArgilosos)
	casas=criarCasas(lista,area_limite,preco_limite)

	atual=imoveis.head
	primeiro=True
	contador=1
	while atual !=None:
		if contador==i:
			soma+=atual.key.Id()
		elif contador==0:
			break
		if primeiro:
			primeiro=False
		else:
			file.write(", ")
		file.write(str(atual.key.Id()))
		atual=atual.next
		contador+=1
	file.write("\n")

	atual=argilosos.head
	primeiro=True
	contador=1
	while atual !=None:
		if contador==j:
			soma+=atual.key.Id()
		if primeiro:
			primeiro=False
		else:
			file.write(", ")
		file.write(str(atual.key.Id()))
		atual=atual.next
		contador+=1
	file.write("\n")

	atual=casas.head
	primeiro=True
	contador=1
	while atual !=None:
		if contador==k:
			soma+=atual.key.Id()
		if primeiro:
			primeiro=False
		else:
			file.write(", ")
		file.write(str(atual.key.Id()))
		atual=atual.next
		contador+=1
	file.write("\n")
	file.close()

	file=open("result.txt","w")
	file.write(str(soma))



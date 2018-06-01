require "class"
require "lista"

-- 90% desse arquivo e :
--linha=lerlinha(file)
--parametro=conversao(lerlinha)

function ler_linha(file) --tirar newline
	linha=file:read()
	if linha ~=nil then
		linha=string.gsub(linha, "\n", "")
	end
	return linha
end

function addCasa(file,lista)
	linha=ler_linha(file)
	id=tonumber(linha)
	linha=ler_linha(file)
	owner=linha
	linha=ler_linha(file)
	quartos=tonumber(linha)
	linha=ler_linha(file)
	vagas=tonumber(linha)
	linha=ler_linha(file)
	pavimentos=tonumber(linha)
	linha=ler_linha(file)
	area_pavimento=tonumber(linha)
	linha=ler_linha(file)
	pm2pavimento=tonumber(linha)
	linha=ler_linha(file)
	area_livre=tonumber(linha)
	linha=ler_linha(file)
	pm2livre=tonumber(linha)
	casa=Casa.new(id,owner,quartos,vagas,pavimentos,area_pavimento,pm2pavimento,area_livre,pm2livre)
	lista:Insert(casa)
end

function addApto(file,lista) --file tipo arquivo lista tipo Lista, cria e insere um apto parametros: id,owner,quartos,vagas,andar,area_construida,pm2construida,lazer,andares
	linha=ler_linha(file)
	id=tonumber(linha)
	linha=ler_linha(file)
	owner=linha
	linha=ler_linha(file)
	quartos=tonumber(linha)
	linha=ler_linha(file)
	vagas=tonumber(linha)
	linha=ler_linha(file)
	andar=tonumber(linha)
	linha=ler_linha(file)
	area_construida=tonumber(linha)
	linha=ler_linha(file)
	pm2construida=tonumber(linha)
	linha=ler_linha(file)
	lazer=linha
	linha=ler_linha(file)
	andares=tonumber(linha)
	apto=Apto.new(id,owner,quartos,vagas,andar,area_construida,pm2construida,lazer,andares)
	lista:Insert(apto)
end

function addTerreno(file,lista,categoria) --file tipo arquivo lista tipo Lista, cria e insere um terreno parametros: id, owner, categoria, solo, precom2, base1 , base2, altura
	linha=ler_linha(file)
	id=tonumber(linha)
	linha=ler_linha(file)
	owner=linha
	linha=ler_linha(file)
	solo=linha
	linha=ler_linha(file)
	precom2=tonumber(linha)
	linha=ler_linha(file)
	base1=tonumber(linha)
	linha=ler_linha(file)
	if categoria=="trapez" then
		base2=tonumber(linha)
		linha=ler_linha(file)
	else
		base2=0
	end
	altura=tonumber(linha)
	terreno=Terreno.new(id, owner, categoria, solo, precom2, base1, base2, altura)
	lista:Insert(terreno)
end

function ler_catalogo(lista) --processa catalogo txt
	file=io.open("catalogo.txt",'r')
	local linha=ler_linha(file) --categoria
	while linha ~= nil do
		if linha=="apto" then
			addApto(file,lista)
		elseif linha=="casa" then
			addCasa(file,lista)
		elseif linha=="triang" or  linha=="trapez" or  linha=="retang" then --se nao for casa ou apto, terreno ja que nao ha erros no input
			addTerreno(file,lista,linha)
		else print("Erro:",linha)
		end
		ler_linha(file) --pular linha vazia
		linha=ler_linha(file) --proxima linha
	end
	file:close()
end

function ler_atual(lista)
	file=io.open("atual.txt",'r')
	linha=ler_linha(file)
	while linha~=nil do
		if linha=='i' then
			linha=ler_linha(file)
			if linha=="apto" then
				addApto(file,lista)
			elseif linha=="casa" then
				addCasa(file,lista)
			elseif linha=="triang" or  linha=="trapez" or  linha=="retang" then --se nao for casa ou apto, terreno ja que nao ha erros no input
				addTerreno(file,lista,linha)
			else print("Erro:",linha)
			end
		elseif linha == 'a' then --seek nao funciona direito entao blocos e blocos
			categoria=ler_linha(file)
			linha=ler_linha(file)
			id=tonumber(linha)
			lista:Remove(id)
			linha=ler_linha(file)
			owner=linha
			if categoria=="apto" then
				linha=ler_linha(file)
				quartos=tonumber(linha)
				linha=ler_linha(file)
				vagas=tonumber(linha)
				linha=ler_linha(file)
				andar=tonumber(linha)
				linha=ler_linha(file)
				area_construida=tonumber(linha)
				linha=ler_linha(file)
				pm2construida=tonumber(linha)
				linha=ler_linha(file)
				lazer=linha
				linha=ler_linha(file)
				andares=tonumber(linha)
				apto=Apto.new(id,owner,quartos,vagas,andar,area_construida,pm2construida,lazer,andares)
				lista:Insert(apto)
			elseif categoria=="casa" then
				linha=ler_linha(file)
				quartos=tonumber(linha)
				linha=ler_linha(file)
				vagas=tonumber(linha)
				linha=ler_linha(file)
				pavimentos=tonumber(linha)
				linha=ler_linha(file)
				area_pavimento=tonumber(linha)
				linha=ler_linha(file)
				pm2pavimento=tonumber(linha)
				linha=ler_linha(file)
				area_livre=tonumber(linha)
				linha=ler_linha(file)
				pm2livre=tonumber(linha)
				casa=Casa.new(id,owner,quartos,vagas,pavimentos,area_pavimento,pm2pavimento,area_livre,pm2livre)
				lista:Insert(casa)
			elseif categoria=="triang" or  categoria=="trapez" or  categoria=="retang" then --se nao for casa ou apto, terreno ja que nao ha erros no input
				linha=ler_linha(file)
				solo=linha
				linha=ler_linha(file)
				precom2=tonumber(linha)
				linha=ler_linha(file)
				base1=tonumber(linha)
				linha=ler_linha(file)
				if categoria=="trapez" then
					base2=tonumber(linha)
					linha=ler_linha(file)
				else
					base2=0
				end
				altura=tonumber(linha)
				terreno=Terreno.new(id, owner, categoria, solo, precom2, base1, base2, altura)
				lista:Insert(terreno)
			else print("Erro categoria:",linha)
			end
		elseif linha == 'e' then
			linha=ler_linha(file)
			id=tonumber(linha)
			lista:Remove(id)
		else print("Erro",linha)
		end
		ler_linha(file) --pular linha vazia
		linha=""
		linha=ler_linha(file)
	end
	file:close()
end

function criarArgilosos(lista,percentArgilosos) --fazer uma lista de argilosos
	local atual=lista.head
	temp=LList.new()
	while atual~=nil do --lista temporaria com todos os argilosos
		if atual.key:Info()=="G" then
			temp:Insert_Argiloso(atual.key)
		end
		atual=atual.next
	end
	atual=temp.head
	contador=temp:Length()
	argilosos=LList.new()
	numeroArgilosos = math.floor(percentArgilosos*temp:Length()/100) --numero do total que vai ser usado
	while atual~=nil do --lista final so com o numero de argilosos necessario
		if contador<=numeroArgilosos then
			argilosos:Insert_Argiloso(atual.key)
		end
		contador=contador-1
		atual=atual.next
	end
	return argilosos
end

function criarCasas(lista, area_limite, preco_limite) --lista de casas que seguem criterios
	local atual=lista.head
	casas=LList.new()
	while atual~=nil do
		if atual.key.categoria=="casa" and atual.key:Preco() < preco_limite and atual.key:Area() > area_limite then
			casas:Insert_Casa(atual.key)
		end
		atual=atual.next
	end
	return casas
end

function criarImoveis(lista,percentImoveis) --n imoveis mais caros
	numeroImoveis = math.floor(percentImoveis*lista:Length()/100)
	contador=lista:Length()
	local atual=lista.head
	imoveis=LList.new()
	while atual~=nil do
		if contador<=numeroImoveis then
			imoveis:Insert(atual.key)
		end
		contador=contador-1
		atual=atual.next

	end
	return imoveis
end

function ler_espec(lista) --le espec e gera as saidas
	soma=0 --soma de i j k esimos
	file=io.open("espec.txt",'r')
	linha=ler_linha(file) 

	percentImoveis=tonumber(linha)	--% dos imoveis
	linha=ler_linha(file) 
	percentArgilosos=tonumber(linha)	--% de imoveis argilosos
	linha=ler_linha(file) 
	area_limite=tonumber(linha)	--area limite para 3 item
	linha=ler_linha(file) 
	preco_limite=tonumber(linha)	--preco limite para 3 item
	linha=ler_linha(file) 
	i=tonumber(linha)	--iesimo imovel mais caro
	linha=ler_linha(file) 
	j=tonumber(linha)	--jesimo argiloso
	linha=ler_linha(file) 
	k=tonumber(linha)	
	io.close(file)

	file=io.open("saida.txt",'w') --escrevendo 
	io.output(file)
	imoveis=criarImoveis(lista,percentImoveis) 
	argilosos=criarArgilosos(lista,percentArgilosos)
	casas=criarCasas(lista,area_limite,preco_limite)

	local atual=imoveis.head
	primeiro=true
	contador=1
	while atual ~=nil do
		if contador==i then
			soma=soma+atual.key:Id()
		end
		if primeiro then
			primeiro=false
		else
			io.write(", ")
		end
		io.write(atual.key:Id())
		atual=atual.next
		contador=contador+1
	end
	io.write("\n")

	atual=argilosos.head
	primeiro=true
	contador=1
	while atual ~=nil do
		if contador==j then
			soma=soma+atual.key:Id()
		end
		if primeiro then
			primeiro=false
		else 
			io.write(", ")
		end
		io.write(atual.key:Id())
		atual=atual.next
		contador=contador+1
	end
	io.write("\n")

	atual=casas.head
	primeiro=true
	contador=1
	while atual ~=nil do
		if contador==k then
			soma=soma+atual.key:Id()
		end
		if primeiro then
			primeiro=false
		else 
			io.write(", ")
		end
		io.write(atual.key:Id())
		atual=atual.next
		contador=contador+1
	end
	io.write("\n")
	io.close(file)

	file=io.open("result.txt","w")
	io.output(file)
	io.write(soma)
	io.close(file)
end
--- Node ---
Node={}
Node.__index=Node

function Node.new(Lote)
  local self=setmetatable({},Node)
  self.key=Lote
  self.next=nil
  return self
end

function  NodeKey()
  return self.key
end

function NodeNext()
  return self.next
end

function NodeHasNext()
  if self.next then
    return true
  else
    return false
  end
end

--- LList ---
LList={}
LList.__index= LList

	function LList.new()
		local self=setmetatable({},LList)
    		self.head=nil
    return self
  end

	function LList:Head()
		return self.head
  end

	function LList:Insert(key)--key é do tipo Lote ou subclasse
		if self.head==nil then
			self.head=Node.new(key)
		else
			atual=self.head --node atual
			antigo=nil --node antigo
			found=false
			novo=Node.new(key)
			while not found do
				p1=atual.key:Preco()
				p2=novo.key:Preco()
				if p1>p2 then --swap novo e atual crescente
					if antigo ~= nil then
						antigo.next=novo
						novo.next=atual

					else
						novo.next=atual
						self.head=novo
          			end
					found = true
				elseif p1==p2 then --desempate por id
					if atual.key:Id() > novo.key:Id() then
						if antigo ~= nil then
							antigo.next=novo
							novo.next=atual
						else
							novo.next=atual
							self.head=novo
            			end
						found = true
         			end
        		end
				if atual.next==nil then
					break
        		end
				antigo=atual
				atual=atual.next
			end
			if not found then
				atual.next=novo
     		end
    	end
  	end

	function LList:Insert_Argiloso(key) --nao verifica tipo so muda criterio de ordenacao
		if self.head==nil then
			self.head=Node.new(key)
		else
			atual=self.head --node atual
			antigo=nil --node antigo
			found=false
			novo=Node(key)
			print(novo.key.Id())
			while not found do
				p1=atual.key:Area()
				p2=novo.key:Area()
				if p1<p2 then--area decrescente
					if antigo ~= nil then
						antigo.next=novo
						novo.next=atual
					else
						novo.next=atual
						self.head=novo
					end
					found = true
				elseif p1==p2 then
					if atual.key:Id() < novo.key:Id() then
						if antigo ~= nil then
							antigo.next=novo
							novo.next=atual
						else
							novo.next=atual
							self.head=novo
						end
						found = true
					end
				end
				if atual.next==nil then
					break
				end
				antigo=atual
				atual=atual.next
			end
			if not found then
				atual.next=novo
			end
		end
	end

	function LList:Insert_Casa(key)--nao verifica tipo so muda criterio de ordenacao
		if self.head==nil then
			self.head=Node.new(key)
		else
			atual=self.head --node atual
			antigo=nil --node antigo
			found=false
			novo=Node(key)
			while not found do
				p1=atual.key:Info() --n de quartos
				p2=novo.key:Info()
				if p1<p2 then --decrscente
					if antigo ~= nil then
						antigo.next=novo
						novo.next=atual
					else
						novo.next=atual
						self.head=novo
					end
					found = true
				elseif p1==p2 then
					if atual.key:Id() < novo.key:Id() then
						if antigo ~= nil then
							antigo.next=novo
							novo.next=atual
						else
							novo.next=atual
							self.head=novo
						end
						found = true
					end
				end
				if atual.next==nil then
					break
				end
				antigo=atual
				atual=atual.next
			end
			if not found then
				atual.next=novo
			end
		end
	end

	function LList:Search(id) --1= true 0= false
		atual=self.head
		while 1 do
			if atual.key:Id()==id then
				return atual
			end

			if atual.next==nil then
				break
			end
			atual=atual.next
		end
		return nil
	end

	function LList:Remove(id) --1= removido 0= nao existe
		atual=self.head
		antigo=nil
		while 1 do
			if atual.key:Id()==id then
				if antigo ~= nil then
					antigo.next=atual.next
					return true
				else
					self.head=atual.next
					return true
				end
			end
			if atual.next==nil then
				break
			end
			antigo=atual
			atual=atual.next
		end
		return false
	end

	function LList:Printar() --para teste // MUDAR
		print("head->")
		atual=self.head
		while 1 do
			print("id",atual.key:Id(),"owner",atual.key:Owner(),"preco",atual.key:Preco(),"\n") -- mudar
			if atual.next==nil then
				break
			end
   			atual=atual.next
    	end
		print("Fim.")
  	end

	function LList:Length()
		l=0
		atual=self.head
		while atual ~=nil do
			l=l+1
			atual=atual.next
    	end
		return l
 	end
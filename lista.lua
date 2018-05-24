--- Node ---
Node={}
Node.__index=Node

function Node.new(Lote)
  local self=setmetatable({},Node)
  self.key=Lote
  self.next=nil
  return self
end

function  Node:Key():
  return self.key
end

function Node:Next():
  return self.next
end

function Node:HasNext():
  if self.next then
    return 1
  else
    return 0
  end
end

--- LList ---
LList={}
LList.__index=LList

	function LList.new():
		local self=setmetatable({},LList)
    self.head=nil
    return self
  end

	function Head():
		return self.head
  end

	function Insert(key):--key é do tipo Lote ou subclasse
		if self.head==nil then
			self.head=Node(key)
		else
			atual=self.head --node atual
			antigo=None --node antigo
			found=0
			novo=Node.new(key)
			while not found do
				p1=atual.key:Preco()
				p2=novo.key:Preco()
				if p1>p2 then --swap novo e atual crescente
					if antigo != None then
						antigo.next=novo
						novo.next=atual
					else
						novo.next=atual
						self.head=novo
          end
					found = 1
				elseif p1==p2: --desempate por id
					if atual.key:Id() > novo.key:Id() then
						if antigo != None then
							antigo.next=novo
							novo.next=atual
						else
							novo.next=atual
							self.head=novo
            end
						found = 1
          end
        end
				if atual.next==None then
					break
        end
				antigo=atual
				atual=atual.next
			if not found then
				atual.next=novo
      end
      end
    end
  end

	function Insert_Argiloso(key): --nao verifica tipo so muda criterio de ordenacao
		if self.head==None:
			self.head=Node(key)
		else:
			atual=self.head --node atual
			antigo=None --node antigo
			found=0
			novo=Node(key)
			while not found:
				p1=atual.key.Area()
				p2=novo.key.Area()
				if p1<p2: --area decrescente
					if antigo != None:
						antigo.next=novo
						novo.next=atual
					else:
						novo.next=atual
						self.head=novo
					found = 1
				elif p1==p2:
					if atual.key.Id() < novo.key.Id():
						if antigo != None:
							antigo.next=novo
							novo.next=atual
						else:
							novo.next=atual
							self.head=novo
						found = 1
				if atual.next==None:
					break
				antigo=atual
				atual=atual.next
			if not found:
				atual.next=novo

	function Insert_Casa(key):--nao verifica tipo so muda criterio de ordenacao
		if self.head==None:
			self.head=Node(key)
		else:
			atual=self.head --node atual
			antigo=None --node antigo
			found=0
			novo=Node(key)
			while not found:
				p1=atual.key.Info() --n de quartos
				p2=novo.key.Info()
				if p1<p2: --decrscente
					if antigo != None:
						antigo.next=novo
						novo.next=atual
					else:
						novo.next=atual
						self.head=novo
					found = 1
				elif p1==p2:
					if atual.key.Id() < novo.key.Id():
						if antigo != None:
							antigo.next=novo
							novo.next=atual
						else:
							novo.next=atual
							self.head=novo
						found = 1
				if atual.next==None:
					break
				antigo=atual
				atual=atual.next
			if not found:
				atual.next=novo

	function Search(id): --1= true 0= false
		atual=self.head
		while 1:
			if atual.key.Id()==id:
				return atual

			if atual.next==None:
				break

			atual=atual.next

		return None

	function Remove(id): --1= removido 0= nao existe
		atual=self.head
		antigo=None
		while 1:
			if atual.key.Id()==id:
				if antigo != None:
					antigo.next=atual.next
					return 1
				else:
					self.head=atual.next
					return 1
			if atual.next==None:
				break

			antigo=atual
			atual=atual.next
		return 0

	function Printar(): --para teste // MUDAR
		print("head->")
		atual=self.head
		while 1:
			print("id: ",atual.key.Id()," owner: ",atual.key.Owner()," preco: {:.2f}".format(atual.key.Preco()))
			if atual.next==None:
				break
      end
			atual=atual.next
    end
		print("Fim.")
  end

	function Length():
		l=0
		atual=self.head
		while atual !=nil:
			l+=1
			atual=atual.next
    end
		return l
  end

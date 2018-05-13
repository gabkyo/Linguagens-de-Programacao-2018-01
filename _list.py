class Node(object):
	def __init__(self, Imovel): #inicializa Node com um imovel
		self.key=Imovel
		self.next=None

	def Key(self):
		return self.key

	def Next(self):
		return self.next

	def HasNext(self): #se next diferente de None
		if self.next:
			return 1
		else:
			return 0


class LList(object): #Lista ou o inicio dos nodes
	def __init__(self):
		self.head=None #head é um node

	def Head(self):
		return self.head

	def Insert(self,key):#key é do tipo Lote ou subclasse
		if self.head==None: 
			self.head=Node(key)
		else:
			atual=self.head #node atual
			antigo=None #node antigo
			found=0
			novo=Node(key)
			while not found:
				p1=atual.key.Preco()
				p2=novo.key.Preco()
				if p1>p2:
					if antigo != None:
						antigo.next=novo
						novo.next=atual
					else:
						novo.next=atual
						self.head=novo
					found = 1
				elif p1==p2:
					if atual.key.Id() > novo.key.Id():
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

	def Search(self,id): #1= true 0= false
		atual=self.head
		while 1:
			if atual.key.Id()==id:
				return atual

			if atual.next==None:
				break

			atual=atual.next

		return None

	def Remove(self,id): #1= true 0= false
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

	def Printar(self): #para teste
		print("head->")
		atual=self.head
		while 1:
			print("id: ",atual.key.Id()," owner: ",atual.key.Owner()," preco: {:.2f}".format(atual.key.Preco()))
			if atual.next==None:
				break
			atual=atual.next
		print("Fim.")

	def Length(self):
		l=0
		atual=self.head
		while atual !=None:
			l+=1
			atual=atual.next

				
		
		
class Lote(object):
	"""docstring for ClassName"""

	def __init__(self, id,owner):
		self.id=id
		self.owner=owner
		self.categoria=''

	def Categoria(self): #retorna o tipo do objeto
		return self.categoria

	def Id(self): #auto explicatorio
		return self.Id

	def Owner(self): #auto explicatorio
		return self.owner

	def Preco(self): #auto explicatorio
		return 0

	def Area(self): #retorna area relevante para calculos ex area total dos pavimentos
		return 0

	def Info(): #retorna informacao necessaria como n de quartos ou tipo do solo
		return ""

class Casa(Lote):
	def __init__(self,id,owner,quartos,vagas,pavimentos,area_pavimento,pm2pavimento,area_livre,pm2livre):
		Lote.__init__(self,id,owner)
		self.categoria='casa'
		self.quartos=quartos
		self.vagas=vagas
		self.pavimentos=pavimentos
		self.area_pavimento=area_pavimento
		self.pm2pavimento=pm2pavimento
		self.area_livre=area_livre
		self.pm2livre=pm2livre

	def Preco(self):
		return self.pm2construida*self.area_pavimento*self.pavimentos+self.pm2livre*self.area_livre

	def Area(self):
		return self.area_pavimento*self.pavimentos

	def Info(self):
		return str(self.quartos)

class Apto(object):
	def __init__(self,id,owner,quartos,vagas,andar,area_construida,pm2construida,lazer,andares):
		Lote.__init__(self,id,owner)
		self.quartos=quartos
		self.vagas=vagas
		self.andar=andar
		self.area_construida=area_construida
		self.pm2construida=pm2construida
		self.lazer=lazer
		self.andares=andares

	def Preco(self):
		if self.lazer=='S':
			return pm2construida*area_construida*(0.9+andar/andares)*1.15
		else:
			return pm2construida*area_construida*(0.9+andar/andares)

	def Area(self):
		return self.area_construida

	def Info(self):
		return ""
		
class Terreno(object):
	def __init__(self, id, owner, categoria, solo, precom2, base1, base2, altura):
		Lote.__init__(self,id,owner)
		self.categoria=categoria
		self.solo=solo
		self.precom2=precom2
		self.base1=base1
		self.base2=base2
		self.altura=altura

	def Preco(self):
		if self.solo=='A':
			return self.precom2*self.Area()*0.9
		elif self.solo=='G':
			return self.precom2*self.Area()*1.3
		elif self.solo=='R':
			return self.precom2*self.Area()*1.1
		else:
			return 0

	def Area(self):
		if self.categoria=="triang":
			return self.base1*self.altura/2
		elif self.categoria=="retang":
			return self.base1*self.altura
		elif self.categoria=="trapez":
			return (self.base1+self.base2)*self.altura/2
		else:
			return 0;

	def Info(self):
		return self.solo
		
---Lote Antecessor comum---
Lote={}
Lote.__index=Lote

function Lote.new(id,owner,categoria)
	local self=setmetatable({},Lote)
	self.id=id
	self.owner=owner
	self.categoria=categoria
	return self
end

function Lote:Preco()
	return 0
end

function Lote:Id()
	return self.id
end

function Lote:Owner()
	return self.owner
end

function Lote:Area()
	return 0
end

function Lote:Info()
	return ""
end

--- Apto ---
Apto={};
Apto.__index=Apto
function Apto.new(id, owner, categoria, nQuartos, nVagas, andar, area_construida, pm2_construida, lazer, nAndares)
	local self=setmetatable({},Apto)
	self.id=id
	self.owner=owner
	self.categoria=categoria
	self.nQuartos=nQuartos
	self.nVagas=nVagas
	self.andar=andar
	self.area_construida=area_construida
	self.pm2_construida=pm2_construida
	self.lazer=lazer
	self.nAndares=nAndares
	return self
end

--Heranca
setmetatable(Apto,{__index = Lote})

--overrides
function Apto:Preco()
	local p=self.pm2_construida * self.area_construida * (0.9 + self.andar/self.nAndares)
	if self.lazer == 'S' then
		return p*1.15
	else
		return p
end

function Apto:Area()
	return self.area_construida
end

--- Terreno ---

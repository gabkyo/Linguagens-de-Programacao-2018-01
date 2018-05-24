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

function Apto.new(id, owner, nQuartos, nVagas, andar, area_construida, pm2_construida, lazer, nAndares)
  local self=setmetatable({},Apto)
  self.id=id
  self.owner=owner
  self.categoria="apto"
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
end

function Apto:Area()
  return self.area_construida
end

--- Terreno ---
Terreno={};
Terreno.__index=Terreno

function Terreno.new(id, owner, categoria, solo, precom2, base1, base2, altura)
  local self=setmetatable({},Terreno)
  self.id=id
  self.owner=owner
  self.categoria=categoria
  self.solo=solo
  self.precom2=precom2
  self.base1=base1
  self.base2=base2
  self.altura=altura
  return self
end

--Heranca
setmetatable(Terreno,{__index = Lote})

--overrides
function Terreno:Area()
  if categoria == "triang" then
    return self.base1 * self.altura/2
  elseif categoria =="retang" then
    return self.base1 * self.altura
  elseif categoria == "trapez" then
    return (self.base1 + self.base2) * self.altura /2
  else
    return 0
  end
end

function Terreno:Preco()
  if self.solo=='A' then
    return self.precom2*self.Area()*0.9
  elseif self.solo=='G' then
    return self.precom2*self.Area()*1.3
  elseif self.solo=='R' then
    return self.precom2*self.Area()*1.1
  else
    return 0
  end
end

function Terreno:Info()
    return self.solo
end

--- Casa ---
Casa={};
Casa.__index=Casa

function Casa.new(id,owner,quartos,vagas,pavimentos,area_pavimento,pm2pavimento,area_livre,pm2livre)
  local self=setmetatable({},Casa)
  self.id=id
  self.owner=owner
  self.categoria="casa"
  self.quartos=quartos
  self.vagas=vagas
  self.pavimentos=pavimentos
  self.area_pavimento=area_pavimento
  self.pm2pavimento=pm2pavimento
  self.area_livre=area_livre
  self.pm2livre=pm2livre
  return self
end

--Heranca
setmetatable(Casa,{__index = Lote})

--overrides
function Casa:Preco()
  return self.pm2pavimento*self.area_pavimento*self.pavimentos+self.pm2livre*self.area_livre
end

function Casa:Area() --area construida da casa nao inclui area livre
  return self.area_pavimento*self.pavimentos
end

function Casa:Info()
  return tostring(self.quartos)
end

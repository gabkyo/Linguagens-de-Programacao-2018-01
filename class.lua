Lote={
	id=0
	owner=nil
	categoria=nil
}
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

function Lote:



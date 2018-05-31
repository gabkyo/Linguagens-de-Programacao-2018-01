require "class"
require "lista"
--nao precisa funcao so escrever o q fazer
teste=LList.new()
a=Terreno.new(45,"Rob","triang",'A',10,4,0,3)
b=Terreno.new(23,"Schneider","triang",'A',10,4,0,3)
teste:Insert(a)
teste:Insert(b) --nao ta inserindo
teste:Printar()
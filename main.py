from _class import *
from _list import *

def main():
	x=Terreno(56,"Rob","triang","R",5,3,0,4)
	teste=LList()
	teste.Insert(x)
	x=Apto(45,"Bob",2,2,5,10,5,"S",10)
	teste.Insert(x)
	teste.Printar()

if __name__ == "__main__":
    main()
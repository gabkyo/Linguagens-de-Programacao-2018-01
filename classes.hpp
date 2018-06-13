#include <string>
#include <vector>

using namespace std;

class Lote{
public:
	Lote(int i, string o,string c){
		id=i;
		owner=o;
		categoria=c;
	}

	int Id(){
		return id;
	}

	string Owner(){
		return owner;
	}

	virtual double Preco(){
		return 0;
	}

	string Categoria(){
		return categoria;
	}

	virtual double Area(){
		return 0;
	}

	virtual int Quartos(){
		return 0;
	}

	virtual string Solo(){
		return "";
	}

protected:
	int id;
	string owner,categoria;
};

class Terreno:public Lote{
public:
	Terreno(int i, string o,string c,string s,double p,double b1,double b2,double a) : Lote(i,o,c){
		solo=s;
		precom2=p;
		base1=b1;
		base2=b2;
		altura=a;
	}

	double Area(){
		if (categoria=="triang"){
			return base1*altura/2;
		}else if (categoria=="retang"){
			return base1*altura;
		}else if (categoria=="trapez"){
			return (base1+base2)*altura/2;
		}else return 0;
	}

	double Preco(){
		if (solo=="A"){
			return precom2 * Area() * 0.9;
		}else if (solo=="G"){
			return precom2 * Area() * 1.3;
		}else if (solo=="R"){
			return precom2 * Area() * 1.1;
		}else return 0;
	}

	string Solo(){
		return solo;
	}

protected:
	string solo;
	double precom2,base1,base2,altura;
};

class Apto:public Lote{
public:
	Apto(int i, string o,int nq,int nv,int a,double ac,double pa,string l,int na) : Lote(i,o,"apto"){
		nQuartos=nq;
		nVagas=nv;
		andar=a;
		areac=ac;
		pm2ac=pa;
		lazer=l;
		nAndares=na;
	}

	double Area(){
		return areac;
	}

	double Preco(){
		if (lazer=="S") {
			return pm2ac*areac*(0.9+(double)andar/(double)nAndares)*1.15;
		}else return pm2ac*areac*(0.9+(double)andar/(double)nAndares);
	}

protected:
	int nQuartos,nVagas,andar,nAndares;
	string lazer;
	double areac,pm2ac;

};

class Casa:public Lote{
public:
	Casa(int i, string o,int q,int v,int p,double ap,double pm2p,double al, double pm2l) : Lote(i,o,"casa"){
		quartos=q;
		vagas=v;
		pavimentos=p;
		areap=ap;
		pm2pav=pm2p;
		areal=al;
		pm2livre=pm2l;
	}

	int Quartos(){
		return quartos;
	}

	double Area(){
		return areap*pavimentos;
	}

	double Preco(){
		return pm2pav*areap*pavimentos+pm2livre*areal;
	}

protected:
	int quartos,vagas,pavimentos;
	double areap,pm2pav,areal,pm2livre;
};

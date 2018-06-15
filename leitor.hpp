#include "list.hpp"
#include <fstream>
#include <algorithm>

string trim(string a){
	a.erase(remove(a.begin(),a.end(),'\n'),a.end());
	return a;
}

void ler_catalogo(LList *lista){
	Lote *novo;
	ifstream catalogo ("catalogo.txt");
	string linha, owner,categoria;
	int id;
	if (catalogo.is_open()) {
		while (getline(catalogo,linha) && !catalogo.eof()) {
			categoria=trim(linha);
			getline(catalogo,linha);
			sscanf(linha.c_str(),"%d",&id);
			cout<<id<<endl;
			getline(catalogo,linha);
			owner=trim(linha);
			getline(catalogo,linha);
			if (categoria=="apto") {
				int nQuartos,nVagas,andar,nAndares;
				double areac,pm2ac;
				string lazer;
				sscanf(linha.c_str(),"%d",&nQuartos);
				getline(catalogo,linha);
				sscanf(linha.c_str(),"%d",&nVagas);
				getline(catalogo,linha);
				sscanf(linha.c_str(),"%d",&andar);
				getline(catalogo,linha);
				sscanf(linha.c_str(),"%lf",&areac);
				getline(catalogo,linha);
				sscanf(linha.c_str(),"%lf",&pm2ac);
				getline(catalogo,linha);
				lazer=trim(linha);
				getline(catalogo,linha);
				sscanf(linha.c_str(),"%d",&nAndares);
				novo=new Apto(id,owner,nQuartos,nVagas,andar,areac,pm2ac,lazer,nAndares);
			}else if (categoria=="casa") {
				int nQuartos,nVagas,pavimentos;
				double areap,pm2pav,areal,pm2livre;
				sscanf(linha.c_str(),"%d",&nQuartos);
				cout<<nQuartos<<" ";
				getline(catalogo,linha);
				sscanf(linha.c_str(),"%d",&nVagas);
				cout<<nVagas<<" ";
				getline(catalogo,linha);
				sscanf(linha.c_str(),"%d",&pavimentos);
				cout<<pavimentos<<" ";
				getline(catalogo,linha);
				sscanf(linha.c_str(),"%lf",&areap);
				cout<<areap<<" ";
				getline(catalogo,linha);
				sscanf(linha.c_str(),"%lf",&pm2pav);
				cout<<pm2pav<<" ";
				getline(catalogo,linha);
				sscanf(linha.c_str(),"%lf",&areal);
				cout<<areal<<" ";
				getline(catalogo,linha);
				sscanf(linha.c_str(),"%lf",&pm2livre);
				cout<<pm2livre<<endl;
				novo=new Casa(id,owner,nQuartos,nVagas,pavimentos,areap,pm2pav,areal,pm2livre);
			}else{
				string solo;
				double pm2,b1,b2,altura;
				solo=trim(linha);
				getline(catalogo,linha);
				sscanf(linha.c_str(),"%lf",&pm2);
				getline(catalogo,linha);
				sscanf(linha.c_str(),"%lf",&b1);
				if (categoria=="trapez") {
					getline(catalogo,linha);
					sscanf(linha.c_str(),"%lf",&b2);
				}else{
					b2=0;
				}
				getline(catalogo,linha);
				sscanf(linha.c_str(),"%lf",&altura);
				novo=new Terreno(id,owner,categoria,solo,pm2,b1,b2,altura);
			}
			getline(catalogo,linha);
			cout<<linha<<"linha ";
			lista->Insert(novo);
		}
		catalogo.close();
	}
}

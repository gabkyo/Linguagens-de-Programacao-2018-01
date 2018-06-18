#include "list.hpp"
#include <fstream>
#include <algorithm>

Lote *new_apto(ifstream& arquivo,int id,string owner){
	Lote *novo;
	int nQuartos,nVagas,andar,nAndares;
	double areac,pm2ac;
	string lazer,linha;

	getline(arquivo,linha);
	sscanf(linha.c_str(),"%d",&nQuartos);
	getline(arquivo,linha);
	sscanf(linha.c_str(),"%d",&nVagas);
	getline(arquivo,linha);
	sscanf(linha.c_str(),"%d",&andar);
	getline(arquivo,linha);
	sscanf(linha.c_str(),"%lf",&areac);
	getline(arquivo,linha);
	sscanf(linha.c_str(),"%lf",&pm2ac);
	getline(arquivo,linha);
	lazer=linha;
	getline(arquivo,linha);
	sscanf(linha.c_str(),"%d",&nAndares);
	novo=new Apto(id,owner,nQuartos,nVagas,andar,areac,pm2ac,lazer,nAndares);
	return novo;
}

Lote *new_casa(ifstream& arquivo,int id,string owner){
	int nQuartos,nVagas,pavimentos;
	double areap,pm2pav,areal,pm2livre;
	Lote *novo;
	string linha;

	getline(arquivo,linha);
	sscanf(linha.c_str(),"%d",&nQuartos);
	getline(arquivo,linha);
	sscanf(linha.c_str(),"%d",&nVagas);
	getline(arquivo,linha);
	sscanf(linha.c_str(),"%d",&pavimentos);
	getline(arquivo,linha);
	sscanf(linha.c_str(),"%lf",&areap);
	getline(arquivo,linha);
	sscanf(linha.c_str(),"%lf",&pm2pav);
	getline(arquivo,linha);
	sscanf(linha.c_str(),"%lf",&areal);
	getline(arquivo,linha);
	sscanf(linha.c_str(),"%lf",&pm2livre);
	novo=new Casa(id,owner,nQuartos,nVagas,pavimentos,areap,pm2pav,areal,pm2livre);
	return novo;
}

Lote *new_terreno(ifstream& arquivo,int id,string owner,string categoria){
	Lote *novo;
	string solo,linha;
	double pm2,b1,b2,altura;

	getline(arquivo,linha);
	solo=linha;
	getline(arquivo,linha);
	sscanf(linha.c_str(),"%lf",&pm2);
	getline(arquivo,linha);
	sscanf(linha.c_str(),"%lf",&b1);
	if (categoria=="trapez") {
		getline(arquivo,linha);
		sscanf(linha.c_str(),"%lf",&b2);
	}else{
		b2=0;
	}
	getline(arquivo,linha);
	sscanf(linha.c_str(),"%lf",&altura);
	novo=new Terreno(id,owner,categoria,solo,pm2,b1,b2,altura);
	return novo;
}


void ler_catalogo(LList *lista){
	Lote *novo;
	ifstream catalogo ("catalogo.txt");
	string linha, owner,categoria;
	int id;

	if (catalogo.is_open()) {
		while (getline(catalogo,linha)) {
			categoria=linha;
			getline(catalogo,linha);
			sscanf(linha.c_str(),"%d",&id);
			getline(catalogo,linha);
			owner=linha;
			if (categoria=="apto") {
				novo=new_apto(catalogo,id,owner);
			}else if (categoria=="casa") {
				novo=new_casa(catalogo,id,owner);
			}else{
				novo=new_terreno(catalogo,id,owner,categoria);
			}
			getline(catalogo,linha);
			lista->Insert(novo);
		}
		catalogo.close();
	}
}

void ler_atual(LList *lista){
	Lote *novo;
	ifstream atual ("atual.txt");
	string linha, owner,categoria;
	int id;

	if (atual.is_open()){
		while(getline(atual,linha)){
			if (linha=="i"){
				getline(atual,linha);
				categoria=linha;
				getline(atual,linha);
				sscanf(linha.c_str(),"%d",&id);
				getline(atual,linha);
				owner=linha;
				if (categoria=="apto") {
					novo=new_apto(atual,id,owner);
				}else if (categoria=="casa") {
					novo=new_casa(atual,id,owner);
				}else{
					novo=new_terreno(atual,id,owner,categoria);
				}
				lista->Insert(novo);
			}else if (linha=="a"){
				getline(atual,linha);
				categoria=linha;
				getline(atual,linha);
				sscanf(linha.c_str(),"%d",&id);
				getline(atual,linha);
				owner=linha;
				lista->Remove(id);
				if (categoria=="apto") {
					novo=new_apto(atual,id,owner);
				}else if (categoria=="casa") {
					novo=new_casa(atual,id,owner);
				}else{
					novo=new_terreno(atual,id,owner,categoria);
				}
				lista->Insert(novo);
			}else{
				getline(atual,linha);
				sscanf(linha.c_str(),"%d",&id);
				lista->Remove(id);
			}
			getline(atual,linha);
		}
	}
	atual.close();
}

LList *Argilosos(LList *lista,int pargilosos){
	node *atual=lista->head;
	LList *temp=new LList(),*argilosos=new LList();
	int total=lista->Tamanho() * pargilosos /100, contador;


	while(atual!=NULL){
		if (atual->key->Solo()=="G"){
			temp->Insert_Argiloso(atual->key);
		}
		atual=atual->next;

	}
	atual=temp->head;
	contador=temp->Tamanho();
	while(atual!=NULL){
		if (contador<=total){
			argilosos->Insert_Argiloso(atual->key);
		}
		contador--;
		atual=atual->next;
	}
	return argilosos
}

LList *Casas(LList *lista,double plim,double arealim){
	node *atual=lista->head;
	LList *casas=new LList();
	while(atual!=NULL){
		if (atual->key->Categoria()=="casa" && atual->key->Preco() < plim && atual->key->Area() > arealim){
			casas->Insert_Casa(atual->key);
		}
		atual=atual->next;
	}
	return casas;
}

LList *Imoveis(LList *lista, int pimoveis){
	int total=pimoveis * lista->Tamanho() /100,contador=lista->Tamanho();
	node *atual=lista->head;
	LList *imoveis=new LList();
	while(atual!=NULL){
		if (contador<=total){
			imoveis->Insert(atual->key);
		}
		contador--;
		atual=atual->next;
	}
	return imoveis; 
}

void ler_espec(LList *lista){
	ifstream espec ("espec.txt");
	int pimoveis,pargilosos,i,j,k,contador,soma=0;
	string linha;
	double arealim,plim;
	LList *imoveis,*casas,*argilosos;

	getline(espec,linha);
	sscanf(linha.c_str(),"%d",&pimoveis);
	getline(espec,linha);
	sscanf(linha.c_str(),"%d",&pargilosos);
	getline(espec,linha);
	sscanf(linha.c_str(),"%lf",&arealim);
	getline(espec,linha);
	sscanf(linha.c_str(),"%lf",&plim);
	getline(espec,linha);
	sscanf(linha.c_str(),"%d",&i);
	getline(espec,linha);
	sscanf(linha.c_str(),"%d",&j);
	getline(espec,linha);
	sscanf(linha.c_str(),"%d",&k);
	espec.close();

	ofstream saida ("saida.txt");
	imoveis=Imoveis(lista,pimoveis);
	casas=Casas(lista,plim,arealim);
	argilosos=Argilosos(lista,pargilosos);
	bool tail=false;

	atual=imoveis->head;
	contador=1;
	while(atual!=NULL){
		if(contador==i){
			soma+=atual->key->Id();
		}else if(contador==0){
			break;
		}
		if (tail){
			saida<<", ";
		}else tail=true;
		saida<<atual->key->Id();
		atual=atual->next;
		contador++;
	}
	saida<<endl;

	atual=argilosos->head;
	contador=1;
	while(atual!=NULL){
		if(contador==i){
			soma+=atual->key->Id();
		}else if(contador==0){
			break;
		}
		if (tail){
			saida<<", ";
		}else tail=true;
		saida<<atual->key->Id();
		atual=atual->next;
		contador++;
	}
	saida<<endl;

	atual=casas->head;
	contador=1;
	tail=false;
	while(atual!=NULL){
		if(contador==j){
			soma+=atual->key->Id();
		}else if(contador==0){
			break;
		}
		if (tail){
			saida<<", ";
		}else tail=true;
		saida<<atual->key->Id();
		atual=atual->next;
		contador++;
	}
	saida<<endl;

	saida.close();
	ofstream result ("result.txt");
	result<<soma<<endl;
	result.close()


}

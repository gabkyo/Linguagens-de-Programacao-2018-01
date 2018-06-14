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
     while (catalogo>>categoria) {
       catalogo>>id;
       catalogo>>owner;
       if (categoria=="apto") {
         int nQuartos,nVagas,andar,nAndares;
         double areac,pm2ac;
         string lazer;
         catalogo>>nQuartos;
         catalogo>>nVagas;
         catalogo>>andar;
         catalogo>>areac;
         catalogo>>pm2ac;
         catalogo>>lazer;
         catalogo>>nAndares;
         novo=new Apto(id,owner,nQuartos,nVagas,andar,areac,pm2ac,lazer,nAndares);
       }else if (categoria=="casa") {

       }else{
         string solo;
         double pm2,b1,b2=0,altura;
         catalogo>>solo;
         catalogo>>pm2;
         if (categoria=="trapez") {
           catalogo>>b1;
           catalogo>>b2;
           catalogo>>altura;
         }else{
           catalogo>>b1;
           catalogo>>altura;
         }
         novo=new Terreno(id,owner,categoria,solo,b1,b2,altura);
       }
       lista->Insert(novo);
     }
     catalogo.close();
   }
}

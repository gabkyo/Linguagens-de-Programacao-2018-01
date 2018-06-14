#include "list.hpp"
#include <fstream>
#include <algorithm>

string trim(string a){
  a.erase(remove(a.begin(),a.end(),'\n'),a.end());
  return a;
}

void ler_catalogo(LList *lista){
   ifstream catalogo ("catalogo.txt");
   string linha, owner,categoria;
   int id;
   if (catalogo.is_open()) {
     while (catalogo>>categoria) {
       catalogo>>id;
       catalogo>>owner;
       if (categoria=="apto") {

       }else if (categoria=="casa") {

       }else{
         string solo;
         double pm2,b1,b2,altura;
         catalogo>>solo;
         catalogo>>pm2;
         if (categoria=="trapez") {
           catalogo>>b1;
           catalogo>>b2;
           catalogo>>altura;
         }
       }
     }
     catalogo.close();
   }
}

#include <string>
#include <vector>
#include <iostream>
#include "list.hpp"

int main() {
  cout<<"ok"<<endl;
  string v="ROb";
  LList *lista=new LList();
  Lote *teste =new Casa(65,v,1,1,1,10.0,10.0,10.0,10.0), *t2=new Casa(10,v,1,1,1,10.0,10.0,10.0,10.0);
  lista->Insert(teste);
  lista->Insert(t2);
  lista->Print();
  return 0;
}

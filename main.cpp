#include <string>
#include <vector>
#include <iostream>
#include "list.hpp"

int main() {
  cout<<"ok"<<endl;
  string v="ROb";
  Lote *teste =new Casa(65,v,1,1,1,10.0,10.0,10.0,10.0);
  cout<<teste->Quartos()<<endl;
  return 0;
}

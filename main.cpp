#include <string>
#include <vector>
#include <iostream>
#include "leitor.hpp"

int main() {
  LList *lista=new LList();
  ler_catalogo(lista);
  cout<<"ok";
  ler_atual(lista);
  ler_espec(lista);
  //lista->Print();
  return 0;
}

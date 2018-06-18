#include <string>
#include <vector>
#include <iostream>
#include "leitor.hpp"

int main() {
  LList *lista=new LList();
  ler_catalogo(lista);
  ler_atual(lista);
  lista->Print();
  return 0;
}

#include <string>
#include <vector>
#include <ios>
#include "classes.hpp"

struct node{
	Lote *key;
	node *next;
};

node *NewNode(Lote *k){
	node *novo= new node();
	novo->key=k;
	novo->next=NULL;
	return novo;
}

class LList{
public:
	LList(){
		head=NULL;
	}

	void Insert(Lote *key){
		bool found=false;
		node *atual=head, *novo=NewNode(key),*antigo=NULL;
		if (head==NULL) {
			head=novo;
		}else{
			while (!found) {
				if (atual->key->Preco() > novo->key->Preco()) {
					if (antigo!=NULL) {
						antigo->next=novo;
						novo->next=atual;
					}else{
						novo->next=atual;
						head=novo;
					}
					found=true;
				}else if (atual->key->Preco() == novo->key->Preco()) {
					if (atual->key->Id() > novo->key->Id()) {
						if (antigo!=NULL) {
							antigo->next=novo;
							novo->next=atual;
						}else{
							novo->next=atual;
							head=novo;
						}
						found=true;
					}
				}
				if (atual->next==NULL) {
					break;
				}
				antigo=atual;
				atual=atual->next;
			}
			if (!found) {
				atual->next=novo;
				found=true;
			}
		}
	}

	void Insert_Argiloso(Lote *key){
		bool found=false;
		node *atual=head, *novo=NewNode(key),*antigo=NULL;
		if (head==NULL) {
			head=novo;
		}else{
			while (!found) {
				if (atual->key->Area() < novo->key->Area()) {
					if (antigo!=NULL) {
						antigo->next=novo;
						novo->next=atual;
					}else{
						novo->next=atual;
						head=novo;
					}
					found=true;
				}else if (atual->key->Area() == novo->key->Area()) {
					if (atual->key->Id() < novo->key->Id()) {
						if (antigo!=NULL) {
							antigo->next=novo;
							novo->next=atual;
						}else{
							novo->next=atual;
							head=novo;
						}
						found=true;
					}
				}
				if (atual->next==NULL) {
					break;
				}
				antigo=atual;
				atual=atual->next;
			}
			if (!found) {
				atual->next=novo;
				found=true;
			}
		}
	}

	void Insert_Casa(Lote *key){
		bool found=false;
		node *atual=head, *novo=NewNode(key),*antigo=NULL;
		if (head==NULL) {
			head=novo;
		}else{
			while (!found) {
				if (atual->key->Quartos() < novo->key->Quartos()) {
					if (antigo!=NULL) {
						antigo->next=novo;
						novo->next=atual;
					}else{
						novo->next=atual;
						head=novo;
					}
					found=true;
				}else if (atual->key->Quartos() == novo->key->Quartos()) {
					if (atual->key->Id() < novo->key->Id()) {
						if (antigo!=NULL) {
							antigo->next=novo;
							novo->next=atual;
						}else{
							novo->next=atual;
							head=novo;
						}
						found=true;
					}
				}
				if (atual->next==NULL) {
					break;
				}
				antigo=atual;
				atual=atual->next;
			}
			if (!found) {
				atual->next=novo;
				found=true;
			}
		}
	}

	void Remove(int id){
		node *atual=head, *antigo=NULL;
		bool found=false;
		if (head==NULL) {
			cout<<"Lista nula nada pra apagar."<<endl;
			return;
		}
		while(!found){
			if (atual->key->Id() == id) {
				if (antigo==NULL) {
					head=atual->next;
					delete atual->key;
					delete atual;
				}else {
					delete atual->key;
					antigo->next=atual->next;
					delete atual;
					found=true;
				}
			}
			if (!found && atual->next==NULL) {
				cout<<"Não encontrado: id "<<id<<endl;
				return;
			}
			antigo=atual;
			atual=atual->next;
		}
	}

	int Tamanho(){
		node *atual=head;
		int contador=0;
		while(atual!=NULL){
			++contador;
			atual=atual->next;
		}
		return contador;
	}



	void Print(){
		cout<<"head-> "<<endl;
		node *atual=head;
		cout.precision(2);
		while(atual!=NULL){
			cout<<"Id: "<<atual->key->Id()<<" owner: "<<atual->key->Owner()<<" preco: "<<fixed<<atual->key->Preco()<<endl;
			atual=atual->next;
		}
	}

	node *Head(){
		return head;
	}

private:
	node *head;
};

package main

import (
"errors"
"fmt"
)

type BasicNode struct {
	key Imoveis
	next *BasicNode
}

func NewBasicNode(key Imoveis) *BasicNode {
	return &BasicNode{key, nil}
}

func (node *BasicNode) Key() Imoveis {
	return node.key
}

func (node *BasicNode) Next() *BasicNode {
	return node.next
}

func (node *BasicNode) HasNext() bool {
	return node.next != nil
}

type BasicList struct {
	head *BasicNode
}

//NewBasicList : Init structure for basic Sorted Linked List.
func NewBasicList() *BasicList {
	return &BasicList{head: NewBasicNode(nil)}
}

func (b *BasicList) Head() *BasicNode {
	return b.head
}

func (b *BasicList) Insert(key Imoveis) { //sort por mais caro
	if b.head == nil || b.head.key==nil {
		b.head = NewBasicNode(key)
	} else {
		var currentNode *BasicNode //node atual
		currentNode = b.head
		var previousNode *BasicNode
		var found bool
		newNode := NewBasicNode(key)
		for {
			c,_:=currentNode.key.Preco()
			d,_:=key.Preco()
			if  c < d { //se o preco novo e maior que o preco do nodo atual
				if previousNode!=nil {
					previousNode.next=newNode
					newNode.next=currentNode
				}else{
					newNode.next=currentNode
					b.head=newNode
				}
				found =true
				break
			}else if c==d {
				if currentNode.key.Id()>newNode.key.Id() { //se preco igual e id do atual for maior que id novo
					if previousNode!=nil {
						previousNode.next=newNode
						newNode.next=currentNode
					}else{
						newNode.next=currentNode
						b.head=newNode
					}
					found =true
				}
			}

			if currentNode.Next() == nil {
				break
			}

			previousNode = currentNode
			currentNode = currentNode.Next()
		}

		if found == false {
			currentNode.next = newNode
		}
	}
}

func (b *BasicList) Search(id uint64) (*BasicNode,error) {
	currentNode := b.head
	for {
		if currentNode.key.Id() == id {
			return currentNode, nil
		}

		if currentNode.Next() == nil {
			break
		}
		currentNode = currentNode.Next()
	}
	return nil, errors.New("key not found")
}

func (b *BasicList) Remove(id uint64) error {
	currentNode := b.head
	var previousNote *BasicNode
	for {
		if currentNode.key.Id() == id {
			if previousNote !=nil {
				previousNote.next = currentNode.Next()
				return nil
			}else {
				b.head=currentNode.Next()
				return nil
			}
			
		}

		if currentNode.Next() == nil {
			break
		}
		previousNote = currentNode
		currentNode = currentNode.Next()
	}
	return errors.New("key "+string(id)+" not found",)
}

func (b *BasicList) Print() {
	fmt.Printf("head->")
	currentNode := b.head
	for {
		c,_:=currentNode.key.Preco();
		fmt.Printf("id: %v owner: %v preco: %.2f\n",currentNode.key.Id(),currentNode.key.Owner(),c)
		//fmt.Printf("[key:%d][val:%v]->", currentNode.key, currentNode.val)
		if currentNode.Next() == nil {
			break
		}
		currentNode = currentNode.Next()
	}
	fmt.Printf("fim")
}
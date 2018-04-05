package main

import (
	"os"
	"fmt"
	"errors"
)

func check(e error){
	if e!=nil {
		panic(e)
	}
}

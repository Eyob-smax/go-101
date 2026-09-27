package main

import (
	"strings"
)

func main(){
	var name = "Eyob"
	var splited = strings.Split(name,"Ey" )

	for i := 0; i < len(splited); i++ {
		print(splited[i], " ")
	}
}


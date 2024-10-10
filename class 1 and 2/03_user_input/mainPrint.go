package main

import (
	"bufio"
	"fmt"
	"os"
)

func mainPrint() {

	//var name string

	//fmt.Println("Enter Your name :")
	//fmt.Scan(&name)

	//fmt.Printf("Name is %s \n", name)

	reader := bufio.NewReader(os.Stdin)

	name, _ := reader.ReadString('\n')

	fmt.Printf("Hello , %s", name)

}

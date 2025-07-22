/*
	Packages is a group of functions, that can be used in other files.
	Importing a package allows you to use its functions and variables in your code.
*/
package main

import (
	"fmt"
	"os"
	"io"
	"log"
	"bytes"
	"strings"
)

func main(){
	fmt.Println(strings.Contains("Computer", "ter"))
	ioPkg()	
	ioutilPkg()
	bytesPkg()
	osPkg()
}

func ioPkg(){
	if _, err := io.WriteString(os.Stdout, "Good Morning!\n"); err != nil{
		log.Fatal(err)
	}
}

func ioutilPkg(){
	msg :=[]byte("Gophers goofing around!")
	err := os.WriteFile("readme.txt", msg, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func bytesPkg(){
	fmt.Printf("%s", bytes.Title([]byte("The way, the truth, and the life.")))
}

func osPkg(){
	f, err := os.OpenFile("notes.txt", os.O_RDWR|os.O_CREATE,0755)
	if err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
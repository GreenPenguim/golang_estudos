/*
Packages is a group of functions, that can be used in other files.
Importing a package allows you to use its functions and variables in your code.
*/
package main

import (
	"bytes"
	"container/list"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

func main(){
	fmt.Println(strings.Contains("Computer", "ter"))
	ioPkg()	
	ioutilPkg()
	bytesPkg()
	osPkg()
	filePath()
	defer errorPkg()
	containerList()
	sortPkg()
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
	fmt.Printf("%s", bytes.Title([]byte("The way, the truth, and the life.\n")))
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

func filePath(){
	filepath.Walk("./module-3",(func (path string, info os.FileInfo, err error) error  {
		fmt.Println(path)
		return nil
	}))
}

type MyError struct{
	When time.Time
	What string
}

func (e MyError) Error() string {
	return fmt.Sprintf("%v: %v", e.When, e.What)
}

func oops() error {
	return MyError{
		time.Now(),
		"File System has gone away",
	}
}

func errorPkg(){
	if err := oops(); err != nil {
		log.Println(err)
	}
}

func containerList(){
	l := list.New()
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

type Data struct{
	name string
	age int
}
type ByName []Data

func (bN ByName) Len() int {
	return len(bN)
}

func (bN ByName) Less(i, j int) bool {
	return bN[i].name < bN[j].name
}

func (bN ByName) Swap(i,j int){
	bN[i], bN[j] = bN[j], bN[i]
}

func sortPkg(){
	people := []Data{
		{"Charlie", 35},
		{"Bob", 25},
	}
	sort.Sort(ByName(people))
	fmt.Println(people)
}
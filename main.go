package main

import "fmt"

func main(){
	tipos()
	arrays()
	params(10)

	

}

func tipos(){
// É possível declarar variáveis sem atribuir tipagem.
	var nome string = "Houdini"
	var idade int = 20
	var versao float32 = 1.0
	fmt.Println("Nome:", nome,",", "Idade:", idade, ",", "Versão:", versao,)
}

func arrays(){
// Tipo unico de dados e tamanho fixo.
	var arr [7]int
	arr[4] = 35
	fmt.Println("Array:", arr)
}

func fatia(){
// É uma parte do Array, tem tipo unico de dados, mas tamanho dinâmico.
}

func mapa(){
// consulta de valor associado a uma chave, a.K.a: hash tables, arrays associativos, dicionários.
}

func params(p1 int){
	fmt.Println("Parâmetro:", p1)
}

func estrutura(){
	// Golang não tem classes, mas é possível criar estruturas (semelhante a classes).
	type Pessoa struct {
		Nome string
		Sobrenome string
		Idade int
	}
	pessoa := Pessoa{"Houdini","Silta", 20} 

	fmt.Println("Pessoa:", pessoa.Nome, pessoa.Sobrenome, pessoa.Idade)
}

func interfaces(){
	type StringDoida interface {
		String() string
	}
}


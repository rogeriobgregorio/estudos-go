package main

import "fmt"

// The init function is called before the main function
// It is typically used for initialization tasks
func init() {
	fmt.Println("Init function called")
}

// Exemplo de função que retorna um erro
func fazerAlgumaCoisa() error {
	// Simulando uma operação que pode falhar
	return fmt.Errorf("algo deu errado")
}

func main() {
	var (
		nome3  string = "Carlos"
		idade2 int    = 28
		ativo1 bool   = true
	)
	var p *int = new(int) // aloca memória para um int e retorna um ponteiro
	*p = 42               // atribui valor ao int apontado por p
	var idade int = 30
	nome1 := "João" // tipo inferido como string
	var x, y int = 10, 20
	var ativo bool // valor zero: false
	const pi = 3.14
	nome2, filhos := "Maria", 30
	//a, b := retornaDoisValores()  // por exemplo: return 10, "ok"
	_ = idade  // evitando erro de variável não utilizada
	_ = nome1  // evitando erro de variável não utilizada
	_ = x      // evitando erro de variável não utilizada
	_ = y      // evitando erro de variável não utilizada
	_ = ativo  // evitando erro de variável não utilizada
	_ = pi     // evitando erro de constante não utilizada
	_ = nome2  // evitando erro de variável não utilizada
	_ = filhos // evitando erro de variável não utilizada
	_ = nome3  // evitando erro de variável não utilizada
	_ = idade2 // evitando erro de variável não utilizada
	_ = ativo1 // evitando erro de variável não utilizada
	_ = p      // evitando erro de variável não utilizada
	//_ = a      // evitando erro de variável não utilizada
	//_ = b      // evitando erro de variável não utilizada

	var nome string = "Rogério"
	fmt.Println("Hello, World!")
	fmt.Println("Meu nome é", nome)

	if x > 10 {
		fmt.Println("Maior que 10")
	}

	if idade >= 18 {
		fmt.Println("Maior de idade")
	} else {
		fmt.Println("Menor de idade")
	}

	if x > 10 {
		fmt.Println("Maior que 10")
	} else if x == 10 {
		fmt.Println("Igual a 10")
	} else {
		fmt.Println("Menor que 10")
	}

	// Declarando e inicializando a variável idade dentro do if
	// A variável idade só existe dentro deste bloco if
	if idade := 20; idade >= 18 {
		fmt.Println("Maior de idade")
	}

	// Definindo uma função anônima que retorna um erro
	fazerAlgo := func() error {
		return fmt.Errorf("algo deu errado")
	}

	// Chamando a função e lidando com o erro
	if err := fazerAlgo(); err != nil {
		fmt.Println("Erro:", err)
	}

	if err := fazerAlgumaCoisa(); err != nil {
		fmt.Println("Erro:", err)
	}
}

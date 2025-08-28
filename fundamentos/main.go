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

	var i any = "uma string"

	// Usando type switch para determinar o tipo de i
	switch v := i.(type) {
	case int:
		fmt.Println("É um int:", v)
	case string:
		fmt.Println("É uma string:", v)
	default:
		fmt.Println("Tipo desconhecido")
	}

	// Loop tradicional
	for j := 0; j < 5; j++ {
		fmt.Println("Valor de j:", j)
	}

	// Loop estilo while
	soma := 0
	for soma < 10 {
		soma++
		fmt.Println("Soma:", soma)
	}

	// Loop infinito com break
	contador := 0
	for {
		if contador >= 3 {
			break
		}
		fmt.Println("Contador:", contador)
		contador++
	}

	// Definindo uma struct
	type Pessoa struct {
		Nome  string
		Idade int
	}

	// Criando uma instância da struct
	pessoa := Pessoa{
		Nome:  "Ana",
		Idade: 25,
	}
	fmt.Println("Pessoa:", pessoa.Nome, "Idade:", pessoa.Idade)

	// Usando ponteiros
	num := 100
	ptr := &num // obtendo o endereço de num
	fmt.Println("Valor de num:", num)
	fmt.Println("Endereço de num:", ptr)
	fmt.Println("Valor via ponteiro:", *ptr) // desreferenciando o ponteiro

	// Usando arrays
	var arr [3]int = [3]int{1, 2, 3}
	fmt.Println("Array:", arr)

	// Usando slices
	slice := []string{"Go", "Python", "Java"}
	slice = append(slice, "C++")
	fmt.Println("Slice:", slice)

	// Usando maps
	idades := map[string]int{
		"Alice": 30,
		"Bob":   25,
	}
	idades["Carlos"] = 28
	fmt.Println("Map:", idades)

	// Usando range para iterar sobre slices e maps
	for index, valor := range slice {
		fmt.Println("Index:", index, "Valor:", valor)
	}

	for chave, valor := range idades {
		fmt.Println("Chave:", chave, "Valor:", valor)
	}

	// Usando defer para garantir que uma função seja chamada no final
	defer
	// Esta mensagem será impressa no final do main
	fmt.Println("Execução finalizada")

	// A função init será chamada antes do main

	// A mensagem "Init function called" será impressa antes de "Hello, World!"
	// A mensagem "Execução finalizada" será impressa no final do main

	// tipos de tratamento de erro
	// 1. Retornar o erro para o chamador
	// Exemplo: if err := fazerAlgo(); err != nil { ... }

	// 2. Logar o erro e continuar
	// Exemplo: log.Println("Erro:", err)

	// 3. Logar o erro e terminar a execução
	// Exemplo: log.Fatal("Erro crítico:", err)

	// 4. Usar panic para erros críticos
	// Exemplo de uso de panic
	// panic("um erro crítico ocorreu") // panic interrompe a execução normal

	// 5. Usar recover para capturar um panic	
	// Exemplo de uso de panic e recover
	func() {
		defer func() { // função anônima para capturar panic
			if r := recover(); r != nil { // recover captura o panic
				fmt.Println("Recuperado de panic:", r) // r contém o valor passado para panic
			}
		}()
		panic("um erro crítico ocorreu") // panic interrompe a execução normal
	}() // função anônima auto-invocada

	// Exemplo de uso de defer para fechar um recurso
	func() {
		// Simulando abertura de um recurso
		fmt.Println("Recurso aberto")
		defer fmt.Println("Recurso fechado") // será executado no final da função
		fmt.Println("Usando o recurso")
		// Simulando erro durante o uso do recurso
		if true {
			fmt.Println("Erro ao usar o recurso")
			return // mesmo com return, o defer será executado
		}
		fmt.Println("Recurso usado com sucesso")
	}() // função anônima auto-invocada

	// Exemplo de uso de defer em loop
	for i := 0; i < 3; i++ {
		defer fmt.Println("Defer no loop:", i) // será executado na ordem inversa
	}
	// A mensagem "Defer no loop: 2", "Defer no loop: 1", "Defer no loop: 0" será impressa aqui

	// Exemplo de uso de ponteiros para modificar valores
	modificarValor := func(p *int) {
		*p = 20 // modifica o valor apontado por p
	}
	func() {
		num := 10
		fmt.Println("Antes:", num)
		modificarValor(&num) // passando o endereço de num
		fmt.Println("Depois:", num) // num será modificado
	}() // função anônima auto-invocada

	// Exemplo de uso de struct e ponteiros
	func() {
		type Retangulo struct {
			Largura, Altura float64
		}
		area := func(r *Retangulo) float64 {
			return r.Largura * r.Altura
		}
		r := Retangulo{Largura: 10, Altura: 5}
		fmt.Println("Área do retângulo:", area(&r)) // passando o endereço de r
	}() // função anônima auto-invocada

	// Mensagem final
	fmt.Println("Fim do programa")
}

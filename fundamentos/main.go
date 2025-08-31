package main

import (
	"fmt"
)

// The init function is called before the main function
// It is typically used for initialization tasks
func init() {
	fmt.Println("Init function called")
}

func main() {
	inicializaVariaveis()
	exibeMensagens()
	condicoes()
	loops()
	estruturaEDados()
	defer fmt.Println("Execução finalizada")
	tratamentoDeErros()
	exemploDefer()
	exemploDeferLoop()
	exemploPonteiros()
	exemploStructPonteiro()
	fmt.Println("Fim do programa")
	mainFuncaoAnonima()
	retornaDoisValores()
	funcComFuncaoComoRetorno()
	funcComVariaveisNomeadas()
	funcComPonteiros(new(int))
	funcComRetornoNomeado()
	funcComFuncaoComoParametro(func(x int) int { return x + 1 }, 10)
	funcComParametrosVariaveis("Alice", "Bob", "Charlie")
	retornaResultado(resultado)
	trataErro(nil)
	funcRecursiva(10)
	exemploContador()
	imprimir(10)
	imprimir("Hello, Generics!")
}

func inicializaVariaveis() {
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
}

func exibeMensagens() {
	var nome string = "Rogério"
	fmt.Println("Hello, World!")
	fmt.Println("Meu nome é", nome)
}

func condicoes() {
	var x int = 10
	var idade int = 30
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
}

func loops() {
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
}

func estruturaEDados() {
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
		fmt.Println("Index:", index, "Valor1:", valor)
	}

	for chave, valor := range idades {
		fmt.Println("Chave:", chave, "Valor:", valor)
	}

	var i any = "uma string"

	// Usando type switch para determinar o tipo de i,
	// o type switch permite verificar o tipo dinâmico de uma variável
	switch v := i.(type) {
	case int:
		fmt.Println("É um int:", v)
	case string:
		fmt.Println("É uma string:", v)
	default:
		fmt.Println("Tipo desconhecido")
	}
}

func tratamentoDeErros() {
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

	// Exemplo de uso de panic e recover
	func() {
		defer func() { // função anônima para capturar panic
			if r := recover(); r != nil { // recover captura o panic
				fmt.Println("Recuperado de panic:", r) // r contém o valor passado para panic
			}
		}()
		panic("um erro crítico ocorreu") // panic interrompe a execução normal
	}() // função anônima auto-invocada
}

func exemploDefer() {
	// Exemplo de uso de defer para fechar um recurso
	// o defer é usado para garantir que uma função seja chamada ao sair do escopo
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
}

func exemploDeferLoop() {
	// Exemplo de uso de defer em loop
	for i := 0; i < 3; i++ {
		defer fmt.Println("Defer no loop:", i) // será executado na ordem inversa
	}
	// A mensagem "Defer no loop: 2", "Defer no loop: 1", "Defer no loop: 0" será impressa aqui
}

func exemploPonteiros() {
	// Exemplo de uso de ponteiros para modificar valores
	modificarValor := func(p *int) {
		*p = 20 // modifica o valor apontado por p
	}
	func() {
		num := 10
		fmt.Println("Antes:", num)
		modificarValor(&num)        // passando o endereço de num
		fmt.Println("Depois:", num) // num será modificado
	}() // função anônima auto-invocada
}

func exemploStructPonteiro() {
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
}

// Exemplo de função que retorna um erro
func fazerAlgumaCoisa() error {
	// Simulando uma operação que pode falhar
	return fmt.Errorf("algo deu errado")
}

// Diferentes tipos de função

// Função que retorna múltiplos valores
func retornaDoisValores() (int, string) {
	return 42, "a resposta"
}

// Função com variáveis nomeadas
func funcComVariaveisNomeadas() (resultado int, mensagem string) {
	resultado = 42
	mensagem = "a resposta"
	return
}

// Função com retorno nomeado
func funcComRetornoNomeado() (resultado int) {
	resultado = 42
	return
}

// Função com ponteiros
// int é um tipo primitivo e um ponteiro é um tipo de referência
// p é um ponteiro para um int
func funcComPonteiros(p *int) {
	*p = 42
}

// Função com parâmetros variáveis
// nomes é um slice de strings
func funcComParametrosVariaveis(nomes ...string) {
	for i, nome := range nomes {
		fmt.Println("Nome", i, ":", nome)
	}
}

// Função com função como parâmetro
// f recebe uma função que recebe um int e retorna um int
// valor é o int que será passado para a função f
func funcComFuncaoComoParametro(f func(int) int, valor int) int {
	return f(valor)
}

// Função com função como retorno
// retorna uma função que recebe um int e retorna um int
func funcComFuncaoComoRetorno() func(int) int {
	return func(x int) int {
		return x * 2
	}
}

// Chamada da função com uma função anônima
var resultado int = funcComFuncaoComoParametro(func(x int) int {
	return x * 2
}, 5)

// Função anônima atribuída a uma variável
func mainFuncaoAnonima() {
	funcao := func(x int) int {
		return x * 2
	}
	fmt.Println("Função anônima chamada com 10:", funcao(10))
	_ = funcao // evitar erro de variável não utilizada
}

// Função que retorna um resultado
func retornaResultado(resultado int) int {
	return resultado
}

// Função que trata um erro
func trataErro(err error) {
	if err != nil {
		fmt.Println("Erro tratado:", err)
	}
}

// Função recursiva
func funcRecursiva(n int) int {
	if n == 0 {
		return 0
	}
	return n + funcRecursiva(n-1)
}

// Função genérica (disponível a partir do Go 1.18)
// uma função genérica é uma função que pode trabalhar com diferentes tipos de dados
func imprimir[T any](valor T) {
	fmt.Println("Valor:", valor)
}

// Função genérica para somar dois valores
func Somar[T int | float64](a, b T) T {
    return a + b
}

type Pessoa struct {
    Nome string
    Idade int
}

// Método vinculado ao tipo Pessoa
func (p Pessoa) Apresentar() string {
    return fmt.Sprintf("Olá, me chamo %s e tenho %d anos", p.Nome, p.Idade)
}

// Método com receptor de ponteiro
// Faz com que a função modifique o valor original
func (p *Pessoa) FazerAniversario() {
    p.Idade++
}

// Função que retorna uma função (closure)
// ela funciona capturando o estado da variável contador
// e permitindo que ele seja incrementado a cada chamada
// ou seja, ela captura o estado da variável contador mesmo após a função ter sido executada
func geraContador() func() int {
    contador := 0
    return func() int {
        contador++
        return contador
    }
}
// Exemplo de uso de closure contador
func exemploContador() {
	// Cria um contador
	var contador1 = geraContador()
	fmt.Println(contador1()) // Imprime: 1
	fmt.Println(contador1()) // Imprime: 2

	// Cria outro contador independente
	var contador2 = geraContador()
	fmt.Println(contador2()) // Imprime: 1 (começa do zero)
	fmt.Println(contador1()) // Imprime: 3 (continua de onde parou)
}

type StringWriter interface {
    WriteString(string) (int, error)
}

type MyWriter struct{}

// WriteString implementa o método WriteString da interface StringWriter
// dessa forma, MyWriter pode ser usado como um StringWriter
func (w MyWriter) WriteString(s string) (int, error) {
    fmt.Print(s)
    return len(s), nil
}
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const numMonitoramento = 2
const delay = 5

func main() {

	exibeIntroducao()

	for {
		exibirMenu()
		comando := leComando()

		switch comando {
		case 1:
			IniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs ...")
			imprimeLog()
		case 0:
			fmt.Println("Saindo do Programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço esse comando:;")
			os.Exit(-1)
		}
	}

}

func exibeIntroducao() {
	nome := "Odair"
	versao := 1.2
	fmt.Println("olá sr.", nome)
	fmt.Println("Você está na versão", versao)
}

func leComando() int {

	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O valor da variavel comando é:", comandoLido)
	//fmt.Println("O endereço da variavel comando é", &comandoLido)
	fmt.Println("")

	return comandoLido
}

func exibirMenu() {

	fmt.Println("1 - iniciar Monitoramento;")
	fmt.Println("2 - Exibir Logs;")
	fmt.Println("0 - Sair do programa;")

}

func IniciarMonitoramento() {
	fmt.Println("Iniciando o Monitoramento ...")

	//sites := []string{"https://random-status-code.herokuapp.com", "https://www.pinari.com.br", "https://moradiariopreto.com.br"}

	sites := leSitesDoArquivo()

	for i := 0; i < numMonitoramento; i++ {
		for i, site := range sites {
			fmt.Println("Testando posição", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")
}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("O site", site, "foi carregado com sucesso")
		registraLog(site, true)
	} else {
		fmt.Println("O site", site, "esta com problemas. StatusCode ", resp.StatusCode)
		registraLog(site, false)
	}

}

func leSitesDoArquivo() []string {

	var sites []string
	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)
		if err == io.EOF {
			break
		}

	}

	arquivo.Close()

	return sites
}

//IMPORTANTE::::
//
// func leSiteDoArquivo()
//
//usa *os (ordem de serviço) e para ler o arquivo "sites.txt", primeiro devemos cria-lo, e depois ...
//criar o "leitor"
//
//leitor que vai receber ferramenta bufio.newReader (novo leitor), vai ler linha por linha do ".TXT" e colocamos '/n' para ele
//ir até o final de cada linha.
//
//a ferramente TrimSpace é para excluir os espaços no final da linha e inclusive o /n que existe em todas as linhas de forma oculta.
//
//sites = append (sites, linha) significa que sites recebe cada linha que for lida pelo leitor e assim pode fazer o monitoramento
//de todo o arquivo ".TXT" // if err == io.EOF { para que após o termino da leitura da ultima linha ele saia da função.
//
// como é tudo linha a linha, fazer dentro de um for, para ele repetir o processo linha a linha do arquivo "".TXT"
//
//no final colocamos um return sites, para retornar para func testaSites() para o programa testar os sites ali inclusos.

func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("Log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/05/2006 15:04:05") + " - " + site + ": Online - " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

//func registraLog() - sites recebe uma stringo e status recebe boleana (verdadeiro ou false)
//**não esquecer de colocar para registrar as informações na func testaSite()
//A função os.Open só abri arquivos, nós iremos abrir e  escrever, então usar outra função os.OpenFile que permite fazer as
//alterações necessarias. olhar no google em "os openFile Golang" as Flags (funções). e pesquisar as autorizações (0666).
//Nesse programa usamos os.O_RDWR(para ler e escrever), os.O_CREATE(para criar o arquivo TXT caso ele não exista), os.O_APPEND(
//para escrever no arquivo o que foi pedido sem sobescrever uma linha sobre a outra.
//
//Para escrever em log.txt, chamamos a função WriteString e passamos os paramentros, mas como status é boloano, devemos chamar
//uma outro função strconv converte outros parametros para string
//
//Para deixar mais preciso e confiavel, colocamos a data e hora que foram feitos os testes, usando time.Now() e para formatar
// como iria aparecer usams time.Now.Format() a formatação pode ser pesquisado aqui, "https://go.dev/src/time/format.go"

func imprimeLog() {

	arquivo, err := ioutil.ReadFile("Log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))

}

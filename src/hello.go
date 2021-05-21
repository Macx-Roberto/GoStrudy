package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

const _repeticoes int = 1
const _tempoPausa int = 1
const _prefixSite string = "Site:"
const _msgSucesso string = " foi carregado com sucesso!"
const _msgErro string = " está com problemas. Status Code:"

func main() {
	for {
		exibeMenu()
		switch leComando() {
		case 1:
			iniciarMonitoramento()
		case 2:
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)
	fmt.Println("")
	return comandoLido
}

func iniciarMonitoramento() {
	sites := retornarSitesArquivo()
	// sites
	// if sites.1 <= 0 {
	// 	os.Exit(0)
	// }

	for i := 0; i <= _repeticoes; i++ {
		for j, site := range sites {
			fmt.Println("Testando site: ", j+1, " - ", site)
			testarSite(site)
		}
		// time.Sleep(100)
	}
}

func retornarSitesArquivo() []string {
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

func testarSite(site string) {

	var cMensagem string
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("erro")
	}

	if resp.StatusCode == 200 {
		cMensagem = _prefixSite + site + _msgSucesso
		fmt.Println(cMensagem)
		registraLog(cMensagem)
	} else {
		cMensagem = _prefixSite + site + _msgErro + err.Error()
		fmt.Println(cMensagem)
		registraLog(cMensagem)
	}
}

func registraLog(log string) {
	arquivo, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Ocorreu um erro ao salvar o log:", err)
	}
	log = time.Now().Format("02/01/2006 15:04:05") + " | " + log + "\n"
	arquivo.WriteString(log)
	arquivo.Close()
}

func imprimeLogs() {
	fmt.Println("Exibindo Logs...")
	arquivo, err := ioutil.ReadFile("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(arquivo))
}

package internal

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

// regex https:\/\/v1\.pinimg\.com\/videos\/mc\/720p\/.*

func GetVideoLink(link string) string {
	result, err := http.Get(link)
	if err != nil {
		fmt.Println("Erro na URL: ", err)
	}

	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		log.Fatalln("Erro no body: : ", err)
	}

	regexPattern := `https:\/\/v1\.pinimg\.com\/videos\/mc\/720p\/.*`
	re := regexp.MustCompile(regexPattern)

	// Encontrando o primeiro link no corpo da resposta que corresponde ao padrão
	links := re.FindString(string(body))

	// Verificando se foi encontrado um link correspondente ao padrão
	if links != "" {
		fmt.Println(strings.Split(links, `",`)[0])
		return strings.Split(links, `",`)[0]
	} else {
		fmt.Println("Nenhum link correspondente ao padrão foi encontrado.")
		return "Erro: nada encontrado."
	}
}

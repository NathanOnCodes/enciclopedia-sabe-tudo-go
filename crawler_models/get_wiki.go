package crawler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"web-crawler-wiki/utils"
)

func GetWiki(name string) (string, string, error) {
	url := fmt.Sprintf("https://pt.wikipedia.org/w/api.php?action=query&format=json&prop=extracts&titles=%s", name)

	resp, err := http.Get(url)
	utils.CheckError(err)
	
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	utils.CheckError(err)

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", "", err
	}

	pages, ok := data["query"].(map[string]interface{})["pages"].(map[string]interface{})
	if !ok {
		// A chave "pages" não está presente na resposta, retorna um erro
		return "", "", fmt.Errorf("Página não encontrada")
	}


var desc string
	for _, v := range pages {
		if v != nil && v.(map[string]interface{})["extract"] != nil {
			desc = v.(map[string]interface{})["extract"].(string)
			break
		}
	}

	if desc == "" {
		return "", "", errors.New("Descrição da página vazia")
	}

	paragraphs := strings.Split(desc, "\n")
	if len(paragraphs) > 2 {
		desc = strings.Join(paragraphs[:2], "\n")
	}

	return desc, url, nil
}
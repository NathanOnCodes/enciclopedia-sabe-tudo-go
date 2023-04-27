package controllers

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"web-crawler-wiki/crawler_models"
	"web-crawler-wiki/utils"
)


func ControllerGeWikiAndPerson(){
	search := flag.Bool("search", false, "pesquisa um arquivo")
	read := flag.Bool("read", false, "lê o histórico de busca")
	del := flag.Bool("del", false, "deleta o histórico de busca")
	flag.Parse()
 
	switch{
	case *search:
		name, err := getInput(os.Stdin, flag.Args()...)
		utils.CheckErrorOnResOfSwitch(err)

		desc, url, err := crawler.GetWiki(name)
		utils.CheckErrorOnResOfSwitch(err)

		var Persons crawler.Persons
		Persons.AddSearch(name, desc, url)
	
		for _, p := range Persons {
			err = p.WriteTxt()
			utils.CheckErrorOnResOfSwitch(err)
		}	

	case *read:
		err := crawler.ReadTxt()
		utils.CheckErrorOnResOfSwitch(err)
	
	case *del:
		err := crawler.DeleteLog(true)
		utils.CheckErrorOnResOfSwitch(err)
	default:
		fmt.Fprintln(os.Stdout, "comando inválido")
		os.Exit(0)
	}
		
}


func getInput (r io.Reader, args ...string) (string, error) {

	if len(args) >  0 {
		return strings.Join(args, " "), nil
	}

	scanner := bufio.NewScanner(r)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}

	text := scanner.Text()

	if len(text) == 0 {
		return "", errors.New("tarefas vazias não são permetidas")
	}
	return text, nil
}
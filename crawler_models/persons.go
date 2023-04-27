package crawler

import (
	"bufio"
	"fmt"
	"html"
	"os"
	"strings"
	"time"
	"web-crawler-wiki/utils"
)

type person struct {
	Name string 
	Desc string
	SearchedAt time.Time
	Url string
}

type Persons []person


func (p *Persons) AddSearch(name, desc, url string){
	personSearch := person{
		Name: name,
		Desc: desc,
		SearchedAt: time.Now(),
		Url: url,
	}
	*p = append(*p, personSearch)
}

func (p *person) WriteTxt() error {
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
	utils.CheckError(err)

	defer file.Close()
	formatDesc := html.EscapeString(p.Desc)

	line := fmt.Sprintf("%s,%s\n%s,%s\n\n", p.Name, formatDesc, p.SearchedAt.Format("02/01/2006 15:04:05"), p.Url)
	
	_ ,err = file.WriteString(line)
	utils.CheckError(err)

	return nil
} 



func ReadTxt() error {
	filename := "log.txt"
	file, err := os.Open(filename)
	utils.CheckError(err)
	defer file.Close()

	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return fmt.Errorf("o arquivo %s não existe", filename)
	}

	if info.Size() == 0 {
		return fmt.Errorf("o arquivo %s está vazio", filename)
	}

	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		line := scanner.Text()

		lineCount++

		if line == "" {
			fmt.Println()
			lineCount = 0
		} else {
			switch lineCount {
			case 1:
				fmt.Println("Nome:\t", line, "\n\nDescrição:\n")
			case 2:
				fmt.Println("\t", strings.ReplaceAll(line, "<p>", ""), "\n")
			case 3:
				parts := strings.Split(line, ",")
				fmt.Printf("\nData de pesquisa:\t%s\nUrl de referência:\t%s\n\n", parts[0], parts[1])
			default:
				return fmt.Errorf("linha inválida: %s", line)
			}

		}
	}

	return nil
}





func DeleteLog(del bool) error {
	if del {
		file, err := os.OpenFile("log.txt", os.O_TRUNC|os.O_WRONLY, 0644)
		utils.CheckError(err)

		defer file.Close()

		if _, err := file.Write(nil); err != nil {
			return err
		}
	}

	return nil
}
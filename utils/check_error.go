package utils

import (
	"fmt"
	"log"
	"os"
)

func CheckError(err error) {
	if err != nil {
		log.Fatal("houve um erro: ", err)
	}
}

func CheckErrorOnResOfSwitch(err error){
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

var (
	bsn       int = 561341758
	number01  int
	userInput int
)

func main() {
	//Having the user provide an input
	fmt.Println("Hello and welcome to this BSN(Sofie) number checker!\nBSN", bsn, "found, initiating check")

	str01 := strconv.Itoa(bsn)
	splitnum := strings.Split(str01, "")
	a := 9
	for _, x := range splitnum {
		dummyvar, err := strconv.Atoi(x)
		if err != nil {
			log.Fatal("Converteren mislukt, functie ValidateBSN ", err)
		}
		if a != 1 {
			number01 += (dummyvar * a)
			fmt.Println(number01)
			a--
		} else if a == 1 {
			number01 += (dummyvar * -1)
			fmt.Println(number01)
		}
	}

	if number01%11 == 0 {
		fmt.Println()
		fmt.Println(bsn, "is a valid BSN")
	}
	fmt.Scanln() //Stopping the program from closing itself.
}

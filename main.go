package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	keyLen           = 10
	defaultStartFlat = 1
	defaultFileName  = "keys.txt"
)

func main() {
	var firstFlat int
	var fileName string
	var input string
	var keysOnOneHuman int
	var i = 0
	flag.IntVar(&firstFlat, "ff", defaultStartFlat, "The apartment where the countdown starts")
	flag.IntVar(&keysOnOneHuman, "keys", defaultStartFlat, "Number of keys per apartment")
	flag.StringVar(&fileName, "f", defaultFileName, "Default name for the file where the keys are written to ")
	flag.Parse()

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()

	fmt.Printf("Начали чтение с квартиры %d: \n", firstFlat)

loop:
	for {
		_, err := fmt.Scanf("%s", &input)
		if err != nil {
			fmt.Printf("Eror on scan: %s", err.Error())
			continue
		}

		inputLen := len([]rune(input))

		switch inputLen {
		case keyLen:
			i++
			fmt.Printf("Key %s, in flat %d \n", input, firstFlat)
			file.WriteString(fmt.Sprintf("Key %s, %d\n", input, firstFlat))
			if i%keysOnOneHuman == 0 {
				firstFlat += 1
			}
		case 0:
			fmt.Println("End of read")
			break loop
		default:
			fmt.Printf("Not read key, need %d symb (your %d) \n", keyLen, inputLen)
			continue
		}
	}
}

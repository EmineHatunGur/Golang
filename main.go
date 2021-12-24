package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	//Kullanıcıdan string değer alma

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("enter stringValues : ")
	str, _ := reader.ReadString('\n')
	fmt.Println(str)

	//Kullanıcıdan int ya da Float değer alma

	values := bufio.NewReader(os.Stdin)
	fmt.Print("enter numbersValues : ")
	val, _ := values.ReadString('\n')
	aNumber, err := strconv.ParseFloat(strings.TrimSpace(val), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Number values: ", aNumber)
	}

}

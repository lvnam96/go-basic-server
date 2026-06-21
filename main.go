package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter number: ")
	str, _ := reader.ReadString('\n')
	f, err := strconv.ParseInt(strings.TrimSpace(str), 10, 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Value of %T: %v\n", int(f), f)
	}
}

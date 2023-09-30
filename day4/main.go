package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fileContents := string(file)
	fileContents = fileContents[:len(fileContents)-1]

	number := 0
	for {
		keyPlusNumber := []byte(fileContents + fmt.Sprint(number))
		md5Hash := md5.Sum(keyPlusNumber)
		// Five zeroes for part one six for part two
		// if strings.HasPrefix(hex.EncodeToString(md5Hash[:]), "00000") {
		if strings.HasPrefix(hex.EncodeToString(md5Hash[:]), "000000") {
			break
		}
		number += 1
	}

	fmt.Printf("%d\n", number)
}

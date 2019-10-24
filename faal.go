package main

import (
	"bufio"
	"crypto/rand"
	"log"
	"math/big"
	"os"
	"strings"
)

var hafez = []string{}

func init() {
	file, err := os.Open("data/hafez.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		hafez = append(hafez, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func GetFaal(str string) string {
	if !strings.Contains(str, KeywordFaal) {
		return ""
	}

	size := len(hafez)
	num, err := rand.Int(rand.Reader, big.NewInt(int64(size/2)))
	if err != nil {
		log.Fatal(err)
	}

	n := int(num.Int64())

	return hafez[2*n] + "\n" + hafez[2*n+1]
}

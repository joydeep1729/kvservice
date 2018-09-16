package main

import (
	"bufio"
	_ "fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//KeyValMap is a map which stores the data from the file
var KeyValMap map[string]int

func loadData(filename string) (map[string]int, error) {
	KvMap := make(map[string]int)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content := scanner.Text()
		temp := strings.Split(content, ",")
		inTemp, err := strconv.Atoi(temp[1])
		if err != nil {
			continue
		}
		KvMap[temp[0]] = inTemp
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return KvMap, nil
}

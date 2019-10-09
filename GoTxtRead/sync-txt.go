package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	dict := make(map[string]int)
	file, err := ioutil.ReadDir("../txt-repo/")
	if err != nil {
		log.Fatal(err)
	}
	fnc := func(name string) {
		data, _ := ioutil.ReadFile("../txt-repo/" + "\\" + name)
		convertString := string(data)
		splited := strings.Split(convertString, " ")

		for _, element := range splited {
			if _, ok := dict[element]; ok {
				dict[element] = dict[element] + 1
			} else {
				dict[element] = 1
			}
		}
	}
	for _, item := range file {
		fnc(item.Name())
	}
	elapsed := time.Since(start)

	fmt.Println("sync")
	fmt.Println()
	fmt.Printf("%v", elapsed)
	fmt.Println()
	fmt.Println(dict["is"])

}

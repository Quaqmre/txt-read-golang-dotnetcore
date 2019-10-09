package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"sync"
	"time"
)

func main() {
	start := time.Now()

	var mutex = &sync.Mutex{}
	var wg sync.WaitGroup

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
			mutex.Lock()
			if _, ok := dict[element]; ok {
				dict[element] = dict[element] + 1
			} else {
				dict[element] = 1
			}
			mutex.Unlock()
		}
		wg.Done()
	}
	for _, item := range file {
		wg.Add(1)
		go fnc(item.Name())
	}
	wg.Wait()
	elapsed := time.Since(start)

	fmt.Println("Async")
	fmt.Println()
	fmt.Printf("%v", elapsed)
	fmt.Println()
	fmt.Println(dict["is"])

}

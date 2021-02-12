package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		files, err := ioutil.ReadDir("./")
		if err != nil {
			log.Fatal(err)
		}

		for i, f := range files {
			if f.IsDir() == true {
				fmt.Printf("%d ---------- %s: %s\n", i+1, f.Name(), "DIRECTORY")
			} else {
				fmt.Printf("%d ---------- %s: %s\n", i+1, f.Name(), "FILE")
			}
		}
	} else {
		files, err := ioutil.ReadDir(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}

		for i, f := range files {
			if f.IsDir() == true {
				fmt.Printf("%d ---------- %s: %s\n", i+1, f.Name(), "DIRECTORY")
			} else {
				fmt.Printf("%d --------- %s: %s\n", i+1, f.Name(), "FILE")
			}
		}
	}
}

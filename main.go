package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	jsonFile, err := os.Open("files_to_copy.json")

	if err != nil {
		log.Fatal(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var files Files

	json.Unmarshal(byteValue, &files)

	for i := 0; i < len(files.Files); i++ {
		fmt.Println(files.Files[i].From)
		copy(files.Files[i].From, files.Files[i].To)
	}

}

func copy(fromPath, toPath string) (int64, error) {
	from, err := os.Open(fromPath)

	if err != nil {
		log.Fatal(err)
	}

	defer from.Close()

	destination, err := os.Create(toPath)

	if err != nil {
		log.Fatal(err)
	}

	defer destination.Close()

	return io.Copy(destination, from)
}

type Files struct {
	Files []File `json:"files"`
}

type File struct {
	From string `json:"from"`
	To   string `json:"to"`
}

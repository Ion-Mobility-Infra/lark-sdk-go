package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/Ion-Mobility-Infra/lark-sdk-go/wrapsanctify"
)

func main() {
	// cat $FILE | sanctify -p model -t $filename > $output
	fmt.Println("generate")

	// grab all files
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	path = filepath.Join(path, "sample")
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		wrapsanctify.HandleCommandLine(f, path)
	}

}

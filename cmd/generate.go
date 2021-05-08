package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/Ion-Mobility-Infra/lark-sdk-go/wrapsanctify"
)

func main() {
	// grab all files
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	srcDirName := "sample"
	srcDir := filepath.Join(path, srcDirName)

	destDirName := "model"
	destDir := filepath.Join(path, destDirName)

	files, err := ioutil.ReadDir(srcDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		// cat $FILE | sanctify -p model -t $filename > $output
		wrapsanctify.HandleFromCommandLine(f, srcDir, destDir)
	}

}

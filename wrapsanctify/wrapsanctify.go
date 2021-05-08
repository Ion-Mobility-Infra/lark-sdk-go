package wrapsanctify

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func HandleFromCommandLine(f os.FileInfo, srcDir string, destDir string) {
	basename := f.Name()
	name := strings.TrimSuffix(basename, filepath.Ext(basename))
	destFilename := name + ".go"

	// cat json file
	catCmd := exec.Command("cat", basename)
	catCmd.Dir = srcDir
	// sanctify json to go struct
	sanctifyCmd := exec.Command("sanctify", "-p", "model", "-t", underscoreToCamelCase(name))
	sanctifyCmd.Dir = srcDir

	// make a pipe
	reader, writer := io.Pipe()
	var buf bytes.Buffer

	// set the output of our cat command
	catCmd.Stdout = writer
	// set the input of the sanctify command
	sanctifyCmd.Stdin = reader

	// cache the outpit of sanctify to memory
	sanctifyCmd.Stdout = &buf

	// execute cat command
	catCmd.Start()
	// execute sanctify command
	sanctifyCmd.Start()

	// waiting for cat command to complete and close the writer
	catCmd.Wait()
	writer.Close()

	// waiting for sanctify command to complete and close the reader
	sanctifyCmd.Wait()
	reader.Close()

	dest := filepath.Join(destDir, destFilename)
	writeBuffer(dest, buf)
}

func writeBuffer(dest string, buf bytes.Buffer) error {
	fmt.Println("Writing to: ", dest)
	return ioutil.WriteFile(dest, buf.Bytes(), 0644)
}

func underscoreToCamelCase(underscore_name string) string {
	strList := strings.Split(underscore_name, "_")
	var CamelCaseNameList []string
	for _, str := range strList {
		CamelCaseNameList = append(CamelCaseNameList, strings.Title(str))
	}
	return strings.Join(CamelCaseNameList, "")
}

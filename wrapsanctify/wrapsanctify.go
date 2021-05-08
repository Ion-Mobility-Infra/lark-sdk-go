package wrapsanctify

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func HandleCommandLine(f os.FileInfo, path string) {
	basename := f.Name()
	name := strings.TrimSuffix(basename, filepath.Ext(basename))
	// output := name + ".go"

	// cmd := exec.Command("cat", basename, "|", "sanctify", "-p", "model", "-t", basename, ">", output)
	catCmd := exec.Command("cat", basename)
	catCmd.Dir = path
	// sanctifyCmd := exec.Command("sanctify", "-p", "model", "-t", basename, ">", output)
	sanctifyCmd := exec.Command("sanctify", "-p", "model", "-t", name)
	// sanctifyCmd := exec.Command("wc")
	sanctifyCmd.Dir = path

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

	io.Copy(os.Stdout, &buf)

	fmt.Println(&buf)
}

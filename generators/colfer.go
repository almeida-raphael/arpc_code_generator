package generators

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/token"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"time"
)

func colferFileToBytes(fileSet *token.FileSet, file *ast.File)([]byte, error){
	buf := new(bytes.Buffer)
	err := format.Node(buf, fileSet, file)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// ColferFile generates a colfer file using an *.arpc.go file pre processed by arpc code parsing functions
func ColferFile(fileSet *token.FileSet, file *ast.File, packagesRootPath, maxListSize, maxSerializedSize string)error{
	fmt.Printf("Generating colf file\n")
	dir, err := ioutil.TempDir("./", "temp_colf_files")
	if err != nil {
		return err
	}
	defer os.RemoveAll(dir)

	colfFilePath := path.Join(dir, fmt.Sprintf("%s,.colf", time.Now().Format("20060102150405")))
	colfFileBytes, err := colferFileToBytes(fileSet, file)
	err = ioutil.WriteFile(colfFilePath, colfFileBytes, 0664)
	if err != nil {
		return err
	}
	defer os.Remove(colfFilePath)

	cmd := exec.Command(
		"colf",
		"-s", maxSerializedSize,
		"-l", maxListSize,
		"-b", packagesRootPath,
		"Go", colfFilePath,
	)
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		return err
	}

	if stderr.Len() > 0 {
		return fmt.Errorf(string(stderr.Bytes()))
	}

	return nil
}

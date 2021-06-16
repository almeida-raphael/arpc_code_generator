package packages

import (
	"fmt"
	"github.com/almeida-raphael/arpc_code_generator/models"
	"github.com/almeida-raphael/arpc_code_generator/parsers"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"path"
)

func getAllPackages(fileSet *token.FileSet, dir string)([]map[string]*ast.Package, error){
	rootPackages, err := parser.ParseDir(fileSet, dir, parsers.FilterARPCFiles,0)
	if err != nil {
		return nil, err
	}
	dirPackages := []map[string]*ast.Package{
		rootPackages,
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			currentDirPackages, err := getAllPackages(fileSet, path.Join(dir, file.Name()))
			if err != nil {
				return nil, err
			}
			dirPackages = append(dirPackages, currentDirPackages...)
		}
	}

	return dirPackages, nil
}

// ParseDir reads a given directory looking for *.arpc.go files and packages and parses them
func ParseDir(dir string)(*token.FileSet, map[string]models.ARPCParsedFile, error){
	fileSet := token.NewFileSet()

	allPackages, err := getAllPackages(fileSet, dir)
	if err != nil {
		return nil, nil, err
	}

	result := make(map[string]models.ARPCParsedFile)
	for _, packages := range allPackages{
		for _, pkg := range packages {
			fmt.Printf("Parsing package: %s\n", pkg.Name)
			for fileName, file := range pkg.Files { // TODO: support more than one file
				fmt.Printf("Parsing file: %s\n", fileName)
				aRPCParsedFile, err := parsers.ParseARPCFile(file)
				if err != nil {
					fmt.Printf("WARNING: error parsing file: %s", fileName)
					continue
				}
				result[fileName] = *aRPCParsedFile
			}
		}
	}

	return fileSet, result, nil
}


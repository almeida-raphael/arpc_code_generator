package main

import (
	"fmt"
	"github.com/almeida-raphael/arpc_code_generator/configs"
	"github.com/almeida-raphael/arpc_code_generator/generators"
	"github.com/almeida-raphael/arpc_code_generator/packages"
	"github.com/fogodev/openvvar"
	"path"
)

func main(){
	config := configs.Config{}
	if err := openvvar.Load(&config); err != nil{
		if err.Error() == "flag: help requested" {
			return
		}
		panic(err)
	}

	fmt.Printf("Parsing:\n")
	fileSet, aRPCParsedFiles, err := packages.ParseDir(config.InputPath)
	if err != nil {
		panic(err)
	}

	fmt.Printf("\nGenerating Files:\n")
	for fileName, aRPCData := range aRPCParsedFiles{
		fmt.Printf("Generating files for: %s\n", fileName)
		err = generators.ColferFile(
			fileSet, aRPCData.ColferFile, config.PackagesRootPath, config.Max.ArraySize,
			config.Max.SerializedSize,
		)
		if err != nil {
			fmt.Printf("WARNING: Error while generating .colf files for %s:\n", fileName)
			continue
		}

		aRPCFileBasePath := path.Join(config.PackagesRootPath, aRPCData.PackageName)

		aRPCClientFilePath := path.Join(aRPCFileBasePath, "arpc_client.go")
		err = generators.ARPCClient(aRPCClientFilePath, aRPCData.PackageName, aRPCData.ARPCProcedures)
		if err != nil {
			fmt.Printf("WARNING: Error while generating aRPC client files for %s:\n", fileName)
			continue
		}

		aRPCServerFilePath := path.Join(aRPCFileBasePath, "arpc_server.go")
		err = generators.ARPCServer(aRPCServerFilePath, aRPCData.PackageName, aRPCData.ARPCProcedures)
		if err != nil {
			fmt.Printf("WARNING: Error while generating aRPC server files for %s:\n", fileName)
			continue
		}
	}
}


package generators

import (
	"bytes"
	"fmt"
	"github.com/almeida-raphael/arpc/helpers"
	"github.com/almeida-raphael/arpc_code_generator/models"
	"github.com/almeida-raphael/arpc_code_generator/templates"
	"github.com/pascaldekloe/name"
	"io/ioutil"
	"strings"
	"text/template"
)

func genARPCClientProcedures(packageName string, procedures []models.Procedure)(string, error){
	var procedureCodes []string
	for idx, procedure := range procedures{
		if procedure.Arg != nil {
			if procedure.Result != nil {
				templateData := templates.ARPCClientProcedureStruct{
					PackageName:   packageName,
					ServiceName:   name.CamelCase(packageName, true),
					ServiceHash:   fmt.Sprint(helpers.Hash(packageName)),
					ProcedureIdx:  idx,
					ProcedureName: procedure.Name,
					ArgName:       procedure.Arg.Name,
					ArgType:       procedure.Arg.TypeName,
					ResponseType:  *procedure.Result,
				}
				tmpl, err := template.New(
					fmt.Sprintf("%s_procedure", procedure.Name),
				).Parse(templates.ARPCClientProcedure)
				if err == nil {
					var procedureCodeBuf bytes.Buffer
					err = tmpl.Execute(&procedureCodeBuf, templateData)
					if err == nil {
						procedureCodes = append(procedureCodes, procedureCodeBuf.String())
						continue
					}
				}
			}
		}
		fmt.Printf(
			"WARNING: template parsing error on procedure %s from service %s", procedure.Name,
			packageName,
		)
	}

	return strings.Join(procedureCodes, "\n\n"), nil
}

// ARPCClient generates aRPC client file for a given service using a parsed procedure list and package name
func ARPCClient(aRPCFilePath, packageName string, procedures []models.Procedure) error {
	fmt.Printf("Generating aRPC client file\n")
	clientProcedures, err := genARPCClientProcedures(packageName, procedures)
	if err != nil {
		return err
	}

	templateData := templates.ARPCClientStruct{
		PackageName:          packageName,
		ServiceName:          name.CamelCase(packageName, true),
		ClientARPCProcedures: clientProcedures,
	}
	tmpl, err := template.New(fmt.Sprintf("%s_aRPC_client", packageName)).Parse(templates.ARPCClient)
	if err != nil {
		return err
	}
	var procedureCodeBuf bytes.Buffer
	err = tmpl.Execute(&procedureCodeBuf, templateData)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(aRPCFilePath, procedureCodeBuf.Bytes(), 0664)
	if err != nil {
		return err
	}

	return nil
}

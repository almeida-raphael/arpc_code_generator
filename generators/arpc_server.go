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

func genARPCServerProcedureDeclarations(packageName string, procedures []models.Procedure)(string, error) {
	var procedureDeclarationsCodes []string
	for _, procedure := range procedures{
		if procedure.Arg != nil {
			if procedure.Result != nil {
				templateData := templates.ARPCServerProcedureDeclarationStruct{
					ProcedureName: procedure.Name,
					ArgName:       name.CamelCase(procedure.Arg.Name, false),
					ArgType:       procedure.Arg.TypeName,
					ResponseType:  *procedure.Result,
				}
				tmpl, err := template.New(
					fmt.Sprintf("%s_procedure_declaration", procedure.Name),
				).Parse(templates.ARPCServerProcedureDeclaration)
				if err == nil {
					var procedureDeclarationCodeBuf bytes.Buffer
					err = tmpl.Execute(&procedureDeclarationCodeBuf, templateData)
					if err == nil {
						procedureDeclarationsCodes = append(
							procedureDeclarationsCodes, procedureDeclarationCodeBuf.String(),
						)
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

	return strings.Join(procedureDeclarationsCodes, "\n"), nil
}

func genARPCServerProcedures(packageName string, procedures []models.Procedure)(string, error) {
	var procedureDeclarationsCodes []string
	for _, procedure := range procedures{
		if procedure.Arg != nil {
			if procedure.Result != nil {
				templateData := templates.ARPCServerProcedureStruct{
					ServiceName:   name.CamelCase(packageName, true),
					ProcedureName: procedure.Name,
					ArgName:       name.CamelCase(procedure.Arg.Name, false),
					ArgType:       procedure.Arg.TypeName,
				}
				tmpl, err := template.New(
					fmt.Sprintf("%s_procedures", procedure.Name),
				).Parse(templates.ARPCServerProcedure)
				if err == nil {
					var procedureDeclarationCodeBuf bytes.Buffer
					err = tmpl.Execute(&procedureDeclarationCodeBuf, templateData)
					if err == nil {
						procedureDeclarationsCodes = append(
							procedureDeclarationsCodes, procedureDeclarationCodeBuf.String(),
						)
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

	return strings.Join(procedureDeclarationsCodes, "\n\n"), nil
}

func genARPCServerProcedureRegistrations(packageName string, procedures []models.Procedure)(string, error) {
	var procedureRegistrations []string
	for idx, procedure := range procedures{
		if procedure.Arg != nil {
			if procedure.Result != nil {
				templateData := templates.ServerARPCProcedureRegistrationStruct{
					ProcedureIdx:  idx,
					ProcedureName: procedure.Name,
				}
				tmpl, err := template.New(
					fmt.Sprintf("%s_procedures", procedure.Name),
				).Parse(templates.ServerARPCProcedureRegistration)
				if err == nil {
					var procedureDeclarationCodeBuf bytes.Buffer
					err = tmpl.Execute(&procedureDeclarationCodeBuf, templateData)
					if err == nil {
						procedureRegistrations = append(
							procedureRegistrations, procedureDeclarationCodeBuf.String(),
						)
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

	return strings.Join(procedureRegistrations, ""), nil
}

// ARPCServer generates aRPC server file for a given service using a parsed procedure list and package name
func ARPCServer(aRPCFilePath, packageName string, procedures []models.Procedure) error {
	fmt.Printf("Generating aRPC server file\n")
	serverProcedureDeclarations, err := genARPCServerProcedureDeclarations(packageName, procedures)
	if err != nil {
		return err
	}

	serverProcedures, err := genARPCServerProcedures(packageName, procedures)
	if err != nil {
		return err
	}

	serverProcedureRegistrations, err := genARPCServerProcedureRegistrations(packageName, procedures)
	if err != nil {
		return err
	}

	templateData := templates.ARPCServerStruct{
		PackageName:                     packageName,
		ServiceHash:                     fmt.Sprint(helpers.Hash(packageName)),
		ServiceName:                     name.CamelCase(packageName, true),
		ServerARPCProcedureDeclarations: serverProcedureDeclarations,
		ServerARPCProcedures:            serverProcedures,
		ServerARPCProcedureRegister:     serverProcedureRegistrations,
	}
	tmpl, err := template.New(fmt.Sprintf("%s_aRPC_server", packageName)).Parse(templates.ARPCServer)
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

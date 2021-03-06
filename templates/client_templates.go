package templates

// ARPCClientProcedureStruct struct for generating a client procedure
type ARPCClientProcedureStruct struct {
	PackageName   string
	ServiceName   string
	ProcedureIdx  int
	ProcedureName string
	ServiceHash   string
	ArgName  	  string
	ArgType  	  string
	ResponseType  string
}

// ARPCClientProcedure client procedure template
var ARPCClientProcedure = `
func ({{.PackageName}} *{{.ServiceName}}){{.ProcedureName}}({{.ArgName}} *{{.ArgType}}, ctx ...context.Context)(*{{.ResponseType}}, error){
	if ctx == nil || len(ctx) == 0{
		ctx = []context.Context{context.Background()}
	}
	response := {{.ResponseType}}{}
	err := {{.PackageName}}.controller.SendRPCCall(ctx[0], {{.ServiceHash}}, {{.ProcedureIdx}}, {{.ArgName}}, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
`

// ARPCClientStruct struct for generating client aRPC code for one service
type ARPCClientStruct struct {
	PackageName          string
	ServiceName          string
	ClientARPCProcedures string
}

// ARPCClient  client aRPC code for one service template
var ARPCClient = `
package {{.PackageName}}

// Code generated by aRPC; DO NOT EDIT.

import (
	"context"
	"github.com/almeida-raphael/arpc/controller"
)

type {{.ServiceName}} struct {
	controller *controller.RPC
}

func New{{.ServiceName}}(controller *controller.RPC) {{.ServiceName}} {
	return {{.ServiceName}}{
		controller: controller,
	}
}

{{.ClientARPCProcedures}}
`

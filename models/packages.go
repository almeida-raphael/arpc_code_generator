package models

import "go/ast"

// ARPCParsedFile structure that stores data to allow colfer file generation and arpc file generation for a service
type ARPCParsedFile struct {
	PackageName   string
	ColferFile    *ast.File
	ARPCProcedures []Procedure
}
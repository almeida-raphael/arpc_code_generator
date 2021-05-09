package parsers

import (
	"github.com/almeida-raphael/arpc_code_generator/models"
	"go/ast"
	"io/fs"
	"strings"
)

// ParseARPCFile receives an *.arpc.go file and retrieves the procedures, args, return types, and struct definitions
func ParseARPCFile(file *ast.File)(*models.ARPCParsedFile, error){
	result := models.ARPCParsedFile{
		PackageName: file.Name.Name,
		ColferFile: file,
	}

	var colferDecls []ast.Decl
	for _, decl := range file.Decls{
		if genDecl, ok := decl.(*ast.GenDecl); ok && len(genDecl.Specs) == 1{
			if typeSpec, ok := genDecl.Specs[0].(*ast.TypeSpec); ok{
				if funcType, ok := typeSpec.Type.(*ast.FuncType); ok {
					currentProcedure :=  models.Procedure{
						Name: typeSpec.Name.Name,
					}
					if len(funcType.Params.List) == 1 {
						if typeStarExpr, ok := funcType.Params.List[0].Type.(*ast.StarExpr); ok {
							if ident, ok := typeStarExpr.X.(*ast.Ident); ok {
								currentProcedure.Arg = &models.Arg{
									Name:     ident.Name,
									TypeName: ident.Obj.Name,
								}
							}
						}
					}
					if len(funcType.Results.List) == 1 || len(funcType.Results.List) == 2 {
						if typeStarExpr, ok := funcType.Results.List[0].Type.(*ast.StarExpr); ok {
							if ident, ok := typeStarExpr.X.(*ast.Ident); ok {
								resultType := ident.Obj.Name
								currentProcedure.Result = &resultType
							}
						}
					}
					result.ARPCProcedures = append(result.ARPCProcedures, currentProcedure)
					continue
				}
			}
		}
		colferDecls = append(colferDecls, decl)
	}

	file.Decls = colferDecls

	return &result, nil
}

// FilterARPCFiles filter function to allow only *.arpc.go files
func FilterARPCFiles(fileInfo fs.FileInfo) bool {
	return strings.HasSuffix(strings.ToLower(fileInfo.Name()),".arpc.go")
}
package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"strings"
)

var builtin = map[string]bool{
	"bool": true, "byte": true, "complex64": true, "complex128": true,
	"error": true, "float32": true, "float64": true, "int": true,
	"int8": true, "int16": true, "int32": true, "int64": true,
	"rune": true, "string": true, "uint": true, "uint8": true,
	"uint16": true, "uint32": true, "uint64": true, "uintptr": true,
	"any": true, "interface{}": true,
}

func exprString(fset *token.FileSet, expr ast.Expr) string {
	var b bytes.Buffer
	_ = printer.Fprint(&b, fset, expr)
	return b.String()
}

func qualify(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		if builtin[t.Name] {
			return t.Name
		}
		return "generated." + t.Name
	case *ast.SelectorExpr:
		return exprString(token.NewFileSet(), expr)
	case *ast.StarExpr:
		return "*" + qualify(t.X)
	case *ast.ArrayType:
		if t.Len == nil {
			return "[]" + qualify(t.Elt)
		}
		return "[" + exprString(token.NewFileSet(), t.Len) + "]" + qualify(t.Elt)
	case *ast.MapType:
		return "map[" + qualify(t.Key) + "]" + qualify(t.Value)
	case *ast.ChanType:
		dir := "chan "
		if t.Dir == ast.SEND {
			dir = "chan<- "
		} else if t.Dir == ast.RECV {
			dir = "<-chan "
		}
		return dir + qualify(t.Value)
	case *ast.InterfaceType:
		return "interface{}"
	case *ast.FuncType:
		return exprString(token.NewFileSet(), expr)
	case *ast.Ellipsis:
		return "..." + qualify(t.Elt)
	default:
		return exprString(token.NewFileSet(), expr)
	}
}

func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "internal/generated/genqlient.go", nil, 0)
	if err != nil {
		panic(err)
	}

	var out bytes.Buffer
	out.WriteString("package mosir_sdk_go\n\n")
	out.WriteString("import (\n\t\"context\"\n\n\tgenerated \"github.com/mosir-social/mosir_sdk_go/internal/generated\"\n)\n\n")
	for _, decl := range f.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok || fn.Recv != nil || !ast.IsExported(fn.Name.Name) || strings.HasSuffix(fn.Name.Name, "ForwardData") {
			continue
		}
		if fn.Type.Params == nil || len(fn.Type.Params.List) < 2 {
			continue
		}
		if exprString(fset, fn.Type.Params.List[0].Type) != "context.Context" || exprString(fset, fn.Type.Params.List[1].Type) != "graphql.Client" {
			continue
		}
		if fn.Type.Results == nil || len(fn.Type.Results.List) != 2 {
			continue
		}

		params := []string{"ctx context.Context"}
		callArgs := []string{"ctx", "c.graphqlClient()"}
		for i, p := range fn.Type.Params.List[2:] {
			typ := qualify(p.Type)
			if len(p.Names) == 0 {
				name := fmt.Sprintf("arg%d", i)
				params = append(params, fmt.Sprintf("%s %s", name, typ))
				callArgs = append(callArgs, name)
				continue
			}
			for _, name := range p.Names {
				params = append(params, fmt.Sprintf("%s %s", name.Name, typ))
				callArgs = append(callArgs, name.Name)
			}
		}

		resultSpecs := []struct{ name, typ string }{}
		for _, r := range fn.Type.Results.List {
			typ := qualify(r.Type)
			if len(r.Names) == 0 {
				resultSpecs = append(resultSpecs, struct{ name, typ string }{"", typ})
			} else {
				for _, name := range r.Names {
					resultSpecs = append(resultSpecs, struct{ name, typ string }{name.Name, typ})
				}
			}
		}
		parts := make([]string, len(resultSpecs))
		for i, r := range resultSpecs {
			if r.name != "" {
				parts[i] = fmt.Sprintf("%s %s", r.name, r.typ)
			} else {
				parts[i] = r.typ
			}
		}
		out.WriteString(fmt.Sprintf("func (c *Client) %s(%s) (%s) {\n", fn.Name.Name, strings.Join(params, ", "), strings.Join(parts, ", ")))
		out.WriteString("\tif c == nil {\n")
		zeroNames := make([]string, len(resultSpecs))
		for i, r := range resultSpecs {
			zeroName := fmt.Sprintf("zero%d", i)
			out.WriteString(fmt.Sprintf("\t\tvar %s %s\n", zeroName, r.typ))
			zeroNames[i] = zeroName
		}
		out.WriteString(fmt.Sprintf("\t\treturn %s\n", strings.Join(zeroNames, ", ")))
		out.WriteString("\t}\n")
		out.WriteString(fmt.Sprintf("\treturn generated.%s(%s)\n", fn.Name.Name, strings.Join(callArgs, ", ")))
		out.WriteString("}\n\n")
	}

	if err := os.WriteFile("client_generated.go", out.Bytes(), 0644); err != nil {
		panic(err)
	}

	aliases := []string{
		"type NotificationFilterInput = generated.NotificationFilterInput",
		"type PostDraftFilterInput = generated.PostDraftFilterInput",
		"type ReactionTypeInput = generated.ReactionTypeInput",
		"type PostType = generated.PostType",
	}
	var aliasOut bytes.Buffer
	aliasOut.WriteString("package mosir_sdk_go\n\nimport generated \"github.com/mosir-social/mosir_sdk_go/internal/generated\"\n\n")
	aliasOut.WriteString("// Re-exported input types used by generated client methods.\n")
	aliasOut.WriteString("type (\n")
	for _, line := range aliases {
		aliasOut.WriteString("\t" + strings.TrimPrefix(line, "type ") + "\n")
	}
	aliasOut.WriteString(")\n")
	if err := os.WriteFile("aliases.go", aliasOut.Bytes(), 0644); err != nil {
		panic(err)
	}
}

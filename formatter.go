package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/parser"
	"github.com/hashicorp/hcl2/hcl"
	"github.com/hashicorp/hcl2/hcl/hclsyntax"
	"github.com/hashicorp/hcl2/hclwrite"
)

// SEPERATOR seperates the original and formatted codes
var (
	changeSeperaor = []byte("**********@@@@@@@@@@&&&&&&&&&&")
	caseSeperator  = []byte("!!!!!!!!!!^^^^^^^^^^----------")
)

// HCLRender renders HCL
func HCLRender(doc ast.Node) []byte {
	var buf bytes.Buffer

	ast.WalkFunc(doc, func(node ast.Node, entering bool) ast.WalkStatus {
		return HCLRenderNode(&buf, node, entering)
	})

	return buf.Bytes()
}

// HCLRenderNode renders a markdown node to HTML
func HCLRenderNode(w io.Writer, node ast.Node, entering bool) ast.WalkStatus {
	switch node := node.(type) {
	case *ast.CodeBlock:
		FormatHCLCodeBlock(w, node)
	}
	return ast.GoToNext
}

// FormatHCLCodeBlock use hcl2 to format HCL configurations
func FormatHCLCodeBlock(w io.Writer, codeBlock *ast.CodeBlock) {
	if !codeBlock.IsFenced {
		return
	}
	// codes are stored  in codeBlock.Literal
	_, syntaxDiags := hclsyntax.ParseConfig(codeBlock.Literal, "<stdin>", hcl.Pos{Line: 1, Column: 1})
	if syntaxDiags.HasErrors() {
		log.Printf("Pass through the CodeBlock not in HCL format")
		return
	}
	// Write the original
	w.Write(codeBlock.Literal)

	// Write seperator
	w.Write([]byte(changeSeperaor))

	// Write the formatted
	w.Write(hclwrite.Format(codeBlock.Literal))

	w.Write(caseSeperator)
}

// Doer :
func Doer(file string) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Printf("Failed to read file:  %s\n", err)
		return
	}

	parser := parser.NewWithExtensions(parser.CommonExtensions)

	rootNode := markdown.Parse(content, parser)

	result := HCLRender(rootNode)

	caseParts := bytes.Split(result, caseSeperator)
	log.Printf("Format %d cases\n", len(caseParts))
	resultContent := content
	for _, v := range caseParts {
		if len(v) == 0 {
			continue
		}
		changeParts := bytes.Split(v, changeSeperaor)
		resultContent = bytes.Replace(resultContent, changeParts[0], changeParts[1], 1)
		ioutil.WriteFile(file, resultContent, 0644)
	}
}

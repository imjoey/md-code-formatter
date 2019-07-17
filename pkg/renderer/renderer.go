package renderer

import (
	"bytes"
	"io"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/parser"
	"github.com/imjoey/md-code-formatter/pkg/formatter"
)

// Statistic gets the details about renderer result
type Statistic struct {
	ChangedCount int
}

// MdCodeRenderer implements github.com/gomarkdown/markdown.Renderer
type MdCodeRenderer struct {
	Content   []byte
	Options   *Options
	CodeDiffs []*formatter.CodeDiff
	Stat      *Statistic
}

// Options :
type Options struct {
	Formatters map[formatter.Lang]formatter.Formatter
}

// Option :
type Option func(*Options)

// WithHCLEnabled :
func WithHCLEnabled() Option {
	return func(o *Options) {
		o.Formatters[formatter.HCL] = formatter.NewHCLFormatter()
	}
}

// NewMdCodeRenderer :
func NewMdCodeRenderer(content []byte, options ...Option) *MdCodeRenderer {
	renderer := &MdCodeRenderer{
		Content: content,
		Options: &Options{
			Formatters: make(map[formatter.Lang]formatter.Formatter),
		},
		Stat: &Statistic{},
	}

	for _, o := range options {
		o(renderer.Options)
	}

	return renderer
}

// RenderNode renders node
func (cr *MdCodeRenderer) RenderNode(w io.Writer, node ast.Node, entering bool) ast.WalkStatus {
	switch node := node.(type) {
	case *ast.CodeBlock:
		codeDiff := formatter.NewCodeDiff(node.Literal)
		if f, ok := cr.Options.Formatters[codeDiff.Lang]; ok {
			codeDiff.After, codeDiff.Errors = f.Format(codeDiff.Before)
		} else {
			codeDiff.Errors = append(codeDiff.Errors, formatter.ErrUnSupportedLang)
		}
		if codeDiff.IsChanged() {
			cr.CodeDiffs = append(cr.CodeDiffs, codeDiff)
		}
	}
	return ast.GoToNext
}

// RenderHeader :
func (cr *MdCodeRenderer) RenderHeader(w io.Writer, ast ast.Node) {
	// No need to implement
}

// RenderFooter :
func (cr *MdCodeRenderer) RenderFooter(w io.Writer, ast ast.Node) {
	// No need to implement
}

// Render :
func (cr *MdCodeRenderer) Render(w io.Writer) {
	parser := parser.NewWithExtensions(parser.CommonExtensions)

	rootNode := markdown.Parse(cr.Content, parser)

	ast.WalkFunc(rootNode, func(node ast.Node, entering bool) ast.WalkStatus {
		return cr.RenderNode(w, node, entering)
	})

	res := make([]byte, len(cr.Content))
	copy(res, cr.Content)
	for _, v := range cr.CodeDiffs {
		if v.IsChanged() {
			res = bytes.Replace(res, v.Before, v.After, 1)
			cr.Stat.ChangedCount++
		}
	}
	w.Write(res)
}

package formatter

import (
	"github.com/hashicorp/hcl2/hcl"
	"github.com/hashicorp/hcl2/hcl/hclsyntax"
	"github.com/hashicorp/hcl2/hclwrite"
)

// HCLFormatter :
type HCLFormatter struct {
}

// Lang :
func (h *HCLFormatter) Lang() Lang {
	return HCL
}

// Format :
func (h *HCLFormatter) Format(before Code) (Code, []error) {
	if before == nil {
		return nil, nil
	}
	if len(before) == 0 {
		return []byte{}, nil
	}

	_, syntaxDiags := hclsyntax.ParseConfig(before, "<stdin>", hcl.Pos{Line: 1, Column: 1})
	if syntaxDiags.HasErrors() {
		return nil, syntaxDiags.Errs()
	}
	after := hclwrite.Format(before)

	return after, nil
}

// NewHCLFormatter :
func NewHCLFormatter() *HCLFormatter {
	return &HCLFormatter{}
}

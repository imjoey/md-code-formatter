package formatter

import (
	"bytes"
	"errors"
)

// Lang :
type Lang string

const (
	// GO :
	GO = Lang("go")
	// HCL :
	HCL = Lang("HCL")
)

// Errors :
var (
	ErrUnSupportedLang = errors.New("UnSupportedLang")
)

// Code :
type Code []byte

// CodeDiff represents a code block before and after the formatting process
// If Errors is not nil, After will also be a nil to save the memory spaces
type CodeDiff struct {
	Lang   Lang
	Before Code
	After  Code
	Errors []error
}

// NewCodeDiff :
func NewCodeDiff(before Code) *CodeDiff {
	return &CodeDiff{
		Before: before,
		Lang:   HCL, // TODO: Add lang detector
	}
}

// IsChanged :
func (c *CodeDiff) IsChanged() bool {
	if len(c.Errors) != 0 {
		return false
	}
	return !bytes.Equal(c.Before, c.After)
}

// Formatter :
type Formatter interface {
	Lang() Lang
	Format(Code) (Code, []error)
}

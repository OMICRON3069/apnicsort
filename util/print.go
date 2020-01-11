package util

import (
	"apnicsort/apnic"
	"io"
)

// PrintSingleLine prints data to io
// without line break
// & insert "separator" between two object
func PrintSingleLine(al []apnic.ResultData, separator string) io.Reader {

}

// PrintLineByLine prints data to io
// with line break
// & insert "separator" between each object and line break
func PrintLineByLine(al []apnic.ResultData, separator string) io.Reader {

}

// ToFile direct writes data to file
func ToFile() {

}

// ToFileWithTemplate insert data to predefine template
func ToFileWithTemplate() {

}

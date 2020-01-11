package util

import (
	"apnicsort/apnic"
)

// Wants sets wanted data's attributes
// Not planing to deal date here
type Wants struct {
	TheType, CC []string
}

// LoadThenProcess returns slice depending on "Wants"
func LoadThenProcess(wt Wants) []apnic.ResultData {
	tql := apnic.LoadURL()
	tql = apnic.GetByCountryCode(tql, wt.CC...)
	tql = apnic.GetByType(tql, wt.TheType...)
	return tql
}

package apnic

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func Test_load(t *testing.T) {
	result := load()
	for _, v := range result {
		fmt.Println(v)
	}
}

func Test_processIP4(t *testing.T) {
	type args struct {
		start      string
		totalValue string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"first", args{"42.62.176.0", "1024"}, "42.62.176.0/22"},
		{"second", args{"42.52.0.0", "262144"}, "42.52.0.0/14"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := processIP4(tt.args.start, tt.args.totalValue); got != tt.want {
				t.Errorf("processIP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_processDate(t *testing.T) {
	type args struct {
		date string
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		// TODO: Add test cases.
		{"", args{"20160108"}, time.Date(2016, time.January, 8, 0, 0, 0, 0, time.UTC)},
		{"", args{"19980709"}, time.Date(1998, time.July, 9, 0, 0, 0, 0, time.UTC)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := processDate(tt.args.date); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("processDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

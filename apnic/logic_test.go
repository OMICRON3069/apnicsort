package apnic

import (
	"container/list"
	"fmt"
	"reflect"
	"testing"
)

func TestGetByCountryCode(t *testing.T) {
	type args struct {
		al *list.List
		cc []string
	}
	tests := []struct {
		name string
		args args
		want *list.List
	}{
		// TODO: Add test cases.
		{
			"",
			args{
				&list.List{},
				[]string{"CN", "JP"},
			},
			&list.List{},
		},
	}
	tests[0].args.al.PushBack(ResultData{
		CC:      "CN",
		TheType: "ipv4",
		IP:      "1.83.0.0/8",
	})
	tests[0].args.al.PushBack(ResultData{
		CC:      "NZ",
		TheType: "ipv4",
		IP:      "2.83.0.0/8",
	})
	tests[0].args.al.PushBack(ResultData{
		CC:      "JP",
		TheType: "ipv4",
		IP:      "3.123.0.0/8",
	})
	tests[0].args.al.PushBack(ResultData{
		CC:      "CN",
		TheType: "ipv4",
		IP:      "4.83.123.0/8",
	})

	got := GetByCountryCode(tests[0].args.al, tests[0].args.cc...)

	for item := got.Front(); item != nil; item = item.Next() {
		s := reflect.ValueOf(item.Value)
		vvvv := s.FieldByName("IP").Interface().(string)
		fmt.Println(vvvv)
	}

}

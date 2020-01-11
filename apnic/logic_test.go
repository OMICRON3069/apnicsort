package apnic

import (
	"container/list"
	"fmt"
	"testing"
)

func TestGetByCountryCode(t *testing.T) {
	type args struct {
		al []ResultData
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
				[]ResultData{
					ResultData{
						CC:      "CN",
						TheType: "ipv4",
						IP:      "1.83.0.0/8",
					},
					ResultData{
						CC:      "NZ",
						TheType: "ipv4",
						IP:      "2.83.0.0/8",
					},
					ResultData{
						CC:      "JP",
						TheType: "ipv4",
						IP:      "3.123.0.0/8",
					},
					ResultData{
						CC:      "CN",
						TheType: "ipv4",
						IP:      "4.83.123.0/8",
					},
				},
				[]string{"CN", "JP"},
			},
			&list.List{},
		},
	}

	got := GetByCountryCode(tests[0].args.al, tests[0].args.cc...)

	for _, v := range got {
		//s := reflect.ValueOf(item.Value).FieldByName("IP").Interface().(string)
		//vvvv := s.FieldByName("IP").Interface().(string)
		//still don't know why it can't works here
		//just move on
		fmt.Println(v)
	}

}

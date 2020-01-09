package apnic

import "testing"

func Test_load(t *testing.T) {
	//Test and view result in debug console
	//just to keep load function can works as expect
	if err := load(); err != nil {
		t.Errorf("hmmmm, interesting")
	}
}

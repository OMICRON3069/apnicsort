package apnic

import (
	"container/list"
	"reflect"
)

// GetByCountryCode returns a list of data by one or more country code
func GetByCountryCode(al *list.List, cc ...string) (resultList *list.List) {
	resultList = list.New()

	for item := al.Front(); item != nil; item = item.Next() {
		//after I switch to use language server, the auto complete
		// just works as I expect
		v := reflect.ValueOf(item.Value).FieldByName("CC").Interface().(string)
		//v := s.FieldByName("CC").Interface().(string)
		for _, ccn := range cc {
			if reflect.DeepEqual(v, ccn) {
				resultList.PushBack(item)
				break
			}
		}
	}
	return
}

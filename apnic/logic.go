package apnic

// GetByCountryCode returns a slice of data by one or more country code
func GetByCountryCode(al []ResultData, cc ...string) (resultList []ResultData) {
	resultList = make([]ResultData, 0)

	for _, v := range al {
		for _, ccn := range cc {
			if v.CC == ccn {
				resultList = append(resultList, v)
				break
			}
		}
	}
	return
}

// GetByType returns a slice by type
func GetByType(al []ResultData, tt ...string) (resultList []ResultData) {
	resultList = make([]ResultData, 0)

	for _, v := range al {
		for _, ttt := range tt {
			if v.TheType == ttt {
				resultList = append(resultList, v)
				break
			}
		}
	}
	return
}

// TODO:
// add more function here

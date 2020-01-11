package apnic

// GetByCountryCode returns a list of data by one or more country code
func (al *[]ResultData) GetByCountryCode(cc ...string) (resultList []ResultData) {
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

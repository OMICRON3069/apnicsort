package apnic

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// ResultData stores processed data
// The schema of apnic address file is
// registry|cc|type|start|value|date|status[|extensions...]
type ResultData struct {
	CC, TheType, IP, ASN string
	AllocationDate       time.Time
}

// example: start:42.62.176.0 value:1024
// 32-log(2\1024) is the cidr value
func processIP4(start, totalValue string) string {
	// convert string number to folat64
	floatValue, _ := strconv.ParseFloat(totalValue, 64)

	// simple calculate to get cidr
	cidr := 32 - math.Log2(floatValue)

	return start + "/" + fmt.Sprintf("%.0f", cidr)
}

// example 20110414
// order   01234567
func processDate(date string) time.Time {

	// hmmmmm
	// month in golang is a renamed int type
	year, _ := strconv.Atoi(date[0:4])
	month, _ := strconv.Atoi(date[4:6])
	day, _ := strconv.Atoi(date[6:])
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

// LoadURL parses result into slice
func LoadURL() (resultList []ResultData) {

	// TODO: make this url configurable
	remoteURL := "http://ftp.apnic.net/apnic/stats/apnic/delegated-apnic-latest"

	// I'm trying to replace list with slice
	// as list has many problem I can't handle
	resultList = make([]ResultData, 0)

	// requesting url
	log.Println("Download start, getting list from apnic......")
	resp, err := http.Get(remoteURL)
	if err != nil {
		log.Println("Error occurs while downloading...")
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()

	// use scanner to load it line by line
	scanner := bufio.NewScanner(resp.Body)
	var line string
	for scanner.Scan() {
		line = scanner.Text()
		// to print line data
		// uncomment the code below
		// log.Println(scanner.Text())

		// if this line start with "apnic"
		// then is a valied line
		if strings.HasPrefix(line, "apnic") {
			result := ResultData{}

			// split line with "|" charactor
			s := strings.Split(line, "|")
			result.CC, result.TheType = s[1], s[2]
			// TODO: deal with asn condition
			if result.TheType == "ipv4" {
				result.IP = processIP4(s[3], s[4])
			} else if result.TheType == "asn" {
				result.ASN = s[3]
			} else if result.TheType == "ipv6" {
				result.IP = s[3] + "/" + s[4]
			}

			// convert date to golang's schema
			result.AllocationDate = processDate(s[5])

			// finaly, creating a new slice
			// with the data just processed
			resultList = append(resultList, result)
		}
	}

	// check whether does it encounter an error
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return
}

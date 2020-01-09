package apnic

import (
	"bufio"
	"container/list"
	"log"
	"net/http"
	"strings"
	"time"
)

// ResultData stores processed data
// The schema of apnic address file is
// registry|cc|type|start|value|date|status[|extensions...]
type ResultData struct {
	cc, thetype, IP string
	allocateDate    time.Time
}

// start:42.62.176.0 value:1024
// 32-log(2\1024) is the cidr value
func processIP(start, totalValue string) string {

}

func processDate(date string) time.Time {

}

func load() (resultList *list.List, errors error) {
	remoteURL := "http://ftp.apnic.net/apnic/stats/apnic/delegated-apnic-latest"

	resultList = list.New()

	log.Println("Download start, get list from apnic......")
	resp, err := http.Get(remoteURL)
	if err != nil {
		errors = err
		log.Print("Error occurs while downloading...")
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	var line string
	for scanner.Scan() {
		line = scanner.Text()
		//log.Println(scanner.Text())

		// if this line start with apnic
		// then is a valied line
		if strings.HasPrefix(line, "apnic") {
			result := ResultData{}
			s := strings.Split(line, "|")
			result.cc, result.thetype = s[1], s[2]
			//TODO: use process functions here
			resultList.PushBack(result)
		}
	}

	if err := scanner.Err(); err != nil {
		errors = err
		log.Fatal(err)
	}
	return
}

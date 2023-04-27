package main

import ( "fmt"
	"net"
    "net/http"
	"time"
	_"log"
	"strings"
	"encoding/json"
	"io/ioutil"
	"github.com/fatih/color"


)


func main() {

	type CertObject struct {
		IssuerCaId int `json:"issuer_ca_id"`
		IssuerName string `json:"issuer_name"`
		CommonName string `json:"common_name"`
		NameValue string `json:"name_value"`
		Id int `json:"id"`
		EntryTimestamp string `json:"entry_timstamp"`
		NotBefore string `json:"not_before"`
		NotAfter string `json:"not_after"`
		SerialNumber string `json:"serial_number`
	}

	c := http.Client{
		Timeout: 60 * time.Second,
	}

    resp, err := c.Get("http://crt.sh/?q=google.com&output=json")
    if err != nil {
        panic(err) 
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var readJson = func(body []byte) (jsonData []CertObject) {
		errJson := json.Unmarshal(body, &jsonData)
		if errJson != nil {
			panic(errJson)
		}
		return 
	}
	
	certData := readJson(body)
	fmt.Println(len(certData))


	// Checkers (isWildCard, isAlive, isCached)
	var isWildCard = func(url string) (isValid bool) {
		if !strings.Contains(url, "*") {
			return true
		} else {
			return false
		}
	}

	var isCached = func(cache map[string]string , url string) (isValid bool) { 
		_, ok := cache[url]
		if ok {
			return false
		} else {
			return true
		}
	}

	
	
	var cache = make(map[string]string)

	for i := range(certData) {
	
		var isValid bool

		isValid = isCached(cache, certData[i].CommonName)
		isValid = isWildCard(certData[i].CommonName)

		if isValid {
			cache[certData[i].CommonName] = certData[i].CommonName
			//fmt.Printf("%s%s\n", isValid, certData[i].CommonName)
		

			

		}
	}

	// Make first round - Show unique sites)
	
	for i := range(cache) {
		// ping url ip, headers, and response code 
		timeout := 1 * time.Second
		_, err := net.DialTimeout("tcp", cache[i]+":80", timeout)
		if err != nil {
			color.Red("[-]" + cache[i])
		} else {
			color.Green("[+]" + cache[i])
		}
	}


}

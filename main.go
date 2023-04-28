package main

import ( 
    _"os"
    "fmt"
	"net"
    "net/http"
	"time"
	_"log"
	"strings"
	"encoding/json"
	"io/ioutil"
	"github.com/fatih/color"
)


// ascii art


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


    
    // Get Command-line arguments

	c := http.Client{
		Timeout: 60 * time.Second,
	}


    fmt.Println("\n Fetching SSL/TLS certificates's...")
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
		}
	}

	// Make first round - Show unique sites)
	
	for i := range(cache) {
		// ping url ip, headers, and response code 
		timeout := 1 * time.Second
		_, err := net.DialTimeout("tcp", cache[i]+":80", timeout)
		if err != nil {
			color.Red("[-] " + cache[i])
            delete(cache, cache[i])

		} else {
			color.Green("[+] " + cache[i])
		}
	}

       
    
    // get tty size. Next enumeration.

    //fmt.Println("-----------")
    fmt.Println("\n\n========================================================================================\n\n")   
    //fmt.Println("-----------")


    // Get Response Headers

    /* 
        Important ones by default.
        Will build a flag to include all response headers.
    */
      

    for i := range(cache) {
        time.Sleep(1 * time.Second)
        color.Yellow("\n[+] %s", cache[i] + "\n")
        resp, err := c.Get("http://" + cache[i])
        if err != nil {
            color.Red("%s", err)
        } else {
            // iterate through response Headers
            for key, val := range resp.Header {
                if key == "Server" || key == "P3p" {
                    color.Green("%s:%s", key, val)
                } else {
                    color.Green("%s", key)
                }
            }
        }

        
    }

}

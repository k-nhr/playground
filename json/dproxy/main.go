package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/koron/go-dproxy"
)

func main() {
	b, err := ioutil.ReadFile("./config")
	if err != nil {
		log.Fatal(err)
	}
	config := strings.Split(string(b), "\n")

	keys := map[string]string{}
	var sk []string
	for _, k := range config {
		key := strings.TrimSpace(k)
		if strings.HasPrefix(key, "#") {
			continue
		}
		split := strings.Split(key, ":")
		keys[split[0]] = split[1]
		sk = append(sk, split[0])
	}

	f, err := os.Create("agentsndlog_summary.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	jsonByteArray, err := ioutil.ReadFile("./agentsnd.log")
	if err != nil {
		log.Fatal(err)
	}

	var sndLodList []interface{}
	err = json.Unmarshal(jsonByteArray, &sndLodList)
	if err != nil {
		log.Fatal(err)
	}
	proxySet := dproxy.NewSet(sndLodList)
	objNum := proxySet.Len()
	fmt.Fprintf(f, "Number of JSON objects : %d\n", objNum)

	title := "#"
	for _, k := range sk {
		title = fmt.Sprintf("%s, %s", title, k)
	}
	fmt.Fprintf(f, "%s\n", title)

	for i := 0; i < objNum; i++ {
		substr := ""
		for _, k := range sk {
			t := keys[k]
			switch t {
			case "string":
				val, err := proxySet.A(i).M(k).String()
				if err != nil {
					log.Fatal(err)
				}
				substr = fmt.Sprintf("%s, %s", substr, val)
			case "float64":
				val, err := proxySet.A(i).M(k).Float64()
				if err != nil {
					log.Fatal(err)
				}
				substr = fmt.Sprintf("%s, %.0f", substr, val)
			default:
				log.Fatal(fmt.Errorf("invalid data type"))
			}
		}
		fmt.Fprintf(f, "%d, %s\n", i, substr[2:])
	}
}

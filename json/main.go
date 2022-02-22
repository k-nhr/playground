package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type sndLog struct {
	FileId           string `json:"fileId"`
	SndFileName      string `json:"sndFileName"`
	HostName         string `json:"hostName"`
	SndStartTimeUtc  int64  `json:"sndStartTimeUtc"`
	SndStartTimeStr  string `json:"sndStartTimeStr"`
	SndEndTimeUtc    int64  `json:"sndEndTimeUtc"`
	SndEndTimeStr    string `json:"sndEndTimeStr"`
	ExitCode         int    `json:"exitCode"`
	DetailCode       int    `json:"detailCode"`
	CompressedRate   int    `json:"compressedRate"`
	CipherType       string `json:"cipherType"`
	RecordCount      int    `json:"recordCount"`
	DataSize         int    `json:"dataSize"`
	TransferRate     int    `json:"transferRate"`
	TransferType     string `json:"transferType"`
	DTelegramLen     int    `json:"dTelegramLen"`
	TriggerEvent     int    `json:"triggerEvent"`
	TriggerStatus    int64  `json:"triggerStatus"`
	StartTransferId  string `json:"startTransferId"`
	LatestTransferId string `json:"latestTransferId"`
	StartProcessId   string `json:"startProcessId"`
	LatestProcessId  string `json:"latestProcessId"`
	PreJobName       string `json:"preJobName"`
	PreJobExitCode   int    `json:"preJobExitCode"`
}

func main() {
	jsonByteArray, err := ioutil.ReadFile("./agentsnd.log")
	if err != nil {
		log.Fatal(err)
	}

	var sndLodList []sndLog
	err = json.Unmarshal(jsonByteArray, &sndLodList)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("agentsndlog_summary.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fmt.Fprintf(f, "Number of JSON objects : %d\n", len(sndLodList))
	for i, v := range sndLodList {
		fmt.Fprintf(f, "%d: %s, exitcode=%d, detailCode=%d\n", i, v.SndEndTimeStr, v.ExitCode, v.DetailCode)
	}

}

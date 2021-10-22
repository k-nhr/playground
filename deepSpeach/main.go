package main

import "C"
import (
	"flag"
	"log"

	"github.com/k-nhr/playground/deepSpeach2/client"
)

func main() {
	model := flag.String("model", "", "Path to the model (protocol buffer binary file)")
	scorer := flag.String("scorer", "", "Path to the external scorer file")
	audio := flag.String("audio", "", "Path to the audio file to run (WAV format)")
	flag.Parse()

	if err := client.CreateModel(*model); err != nil {
		log.Fatal(err)
	}
	if err := client.EnableExternalScorer(*scorer); err != nil {
		log.Fatal(err)
	}
	client.ProcessFile(*audio)
	client.FreeModel()
	return
}

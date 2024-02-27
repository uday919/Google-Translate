package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"sync" //helps to work with wait groups

	"github.com/uday919/google-translate/cli"
)

var wg sync.WaitGroup //wg is of type

var sourceLang string
var targetLang string
var sourceText string

func init() {
	flag.StringVar(&sourceLang, "s", "en", "Source language[en]")
	flag.StringVar(&targetLang, "t", "fr", "Target language[fr]")
	flag.StringVar(&sourceText, "st", "", "Text to Translate")

}
func main() {
	flag.Parse()
	if flag.NFlag() == 0 { //NFlag to check the user has provided any flags or not
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(1)
	}
	strChan := make(chan string) //creating a channel
	wg.Add(1)                    //adding two tasks in the Wait Group
	reqBody := &cli.RequestBody{
		SourceLang: sourceLang,
		TargetLang: targetLang,
		SourceText: sourceText,
	}
	go cli.RequestTranslate(reqBody, strChan, &wg)
	processedStr := strings.ReplaceAll(<-strChan, "+", "  ") //printing the translated text
	fmt.Printf("%s\n", processedStr)
	wg.Wait() //waiting for all goroutines to finish their execution
}

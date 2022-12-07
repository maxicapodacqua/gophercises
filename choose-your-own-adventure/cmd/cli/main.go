package main

import (
	"encoding/json"
	"fmt"
	"github.com/maxicapodacqua/choose-your-own-adventure/internal"
	"io/ioutil"
	"os"
	"strings"
)

func printPage(p internal.Page) string {
	fmt.Printf("# %v #\n  %v", p.Title, strings.Join(p.Story, "\n"))
	if len(p.Options) > 0 {
		fmt.Printf("Options:\n ")
		for i, option := range p.Options {
			fmt.Printf("%v: %v\n ", i+1, option.Text)
		}
		var selection int
		fmt.Printf("Choose option:")
		fmt.Scan(&selection)
		fmt.Println()
		return p.Options[selection-1].Arc
	} else {
		fmt.Printf("\nThe End.\n")
		return ""
	}
}

func main() {

	f, _ := os.Open("./gopher.json")

	var allStories map[string]internal.Page

	byteVal, _ := ioutil.ReadAll(f)
	json.Unmarshal(byteVal, &allStories)

	for next := printPage(allStories["intro"]); next != ""; next = printPage(allStories[next]) {
	}
}

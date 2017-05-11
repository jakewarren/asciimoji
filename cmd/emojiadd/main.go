package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
	"io/ioutil"
)

type Emoji struct {
	Keywords []string `json:"keywords"`
	Emoji    string   `json:"emoji"`
}


var emojis []Emoji



func init() {

	file, e := ioutil.ReadFile("./emojis.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}


	if err := json.Unmarshal(file, &emojis); err != nil {
		panic(err)
	}


}

var (
	app = kingpin.New("emojiadd", "easily add emoji to json file")

	emoji = app.Arg("emoji", "the emoji").Required().String()
	keywords = app.Arg("keywords", "keywords of the emoji").Required().Strings()


)

func main() {

	app.Version("0.1").VersionFlag.Short('V')
	app.HelpFlag.Short('h')
	app.UsageTemplate(kingpin.SeparateOptionalFlagsUsageTemplate)
	kingpin.MustParse(app.Parse(os.Args[1:]))

	newEmoji := Emoji{*keywords,*emoji}

	emojis = append(emojis,newEmoji)

	//write out the file
	jsonstr,_ := json.MarshalIndent(emojis,"", "  ")

	e := ioutil.WriteFile("./emojis.json",jsonstr,0644)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
}


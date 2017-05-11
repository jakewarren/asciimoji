package main

import (
	"encoding/json"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/olekukonko/tablewriter"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
	"strconv"
	"strings"
)

type Emoji struct {
	Keywords []string `json:"keywords"`
	Emoji    string   `json:"emoji"`
}

var keywordLookUp map[string][]Emoji
var emojis []Emoji

func buildKeywordLookupMap() {
	keywordLookUp = make(map[string][]Emoji, 0)
	for _, emoji := range emojis {
		for _, keyword := range emoji.Keywords {
			keywordLookUp[keyword] = append(keywordLookUp[keyword], emoji)
		}
	}
}

func init() {
	/*
		file, e := ioutil.ReadFile("./emojis.json")
		if e != nil {
			fmt.Printf("File error: %v\n", e)
			os.Exit(1)
		}
	*/

	if err := json.Unmarshal(MustAsset("emojis.json"), &emojis); err != nil {
		panic(err)
	}

	buildKeywordLookupMap()

}

var (
	app = kingpin.New("asciimoji", "Look up ASCII emojis")

	list_all = app.Flag("list-all", "list all available emojis").Short('l').Bool()
	search   = app.Flag("search", "search all available emojis").Short('s').Strings()

	query = app.Arg("lookup", "the keyword to lookup.").String()
)

func main() {

	app.Version("0.1").VersionFlag.Short('V')
	app.HelpFlag.Short('h')
	app.UsageTemplate(kingpin.SeparateOptionalFlagsUsageTemplate)
	kingpin.MustParse(app.Parse(os.Args[1:]))

	if *list_all {
		listEmojis()
		os.Exit(0)
	}

	if len(*search) > 0 {
		for _, s := range *search {
			searchEmoji(s)
		}
		os.Exit(0)
	}

	if len(*query) > 0 {
		getEmoji(*query)
	} else {
		fmt.Println("no search term was entered.")
	}

}

func searchEmoji(query string) {

	table := tablewriter.NewWriter(os.Stdout)
	for _, e := range emojis {
		for _, k := range e.Keywords {

			if strings.Contains(k, query) {
				table.Append([]string{e.Emoji, strings.Join(e.Keywords, ", ")})
			}
		}
	}
	table.Render()
}

func getEmoji(query string) {

	emojis, ok := keywordLookUp[query]
	if !ok {
		fmt.Println("emoji not found :(")
		os.Exit(0)
	}

	var emoji Emoji
	if len(emojis) > 1 {
		for {
			for i, emj := range emojis {
				fmt.Printf("%d) %s\n", i+1, emj.Emoji)
			}

			var raw string
			fmt.Printf("choice [1-%d]: ", len(emojis))
			if _, err := fmt.Scanf("%s", &raw); err == nil {
				if choice, err := strconv.Atoi(raw); err == nil {
					if choice >= 1 && choice <= len(emojis) {
						emoji = emojis[choice-1]
						break
					}
				}
			}

			fmt.Println("invalid choice, please try again")
		}

	} else {
		emoji = emojis[0]
	}

	if err := clipboard.WriteAll(emoji.Emoji); err != nil {
		panic(err)
	}
	fmt.Printf("copied %s\n", emoji.Emoji)
}

func listEmojis() {
	table := tablewriter.NewWriter(os.Stdout)
	for _, emoji := range emojis {
		//fmt.Println(emoji.Emoji,"\t",emoji.Keywords)

		table.Append([]string{emoji.Emoji, strings.Join(emoji.Keywords, ", ")})
	}
	table.Render()
}

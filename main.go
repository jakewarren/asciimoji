package main

import (
	"encoding/json"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/olekukonko/tablewriter"
	"gopkg.in/AlecAivazis/survey.v1"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
	"strings"
)

//Emoji stores a struct to match keywords to an emoji
type Emoji struct {
	Keywords []string `json:"keywords"`
	Emoji    string   `json:"emoji"`
}

//map the keywords to emojis
var keywordLookUp map[string][]Emoji

//store a list of the emojis
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

	//read in the emoji "database"
	if err := json.Unmarshal(MustAsset("emojis.json"), &emojis); err != nil {
		panic(err)
	}

	buildKeywordLookupMap()

}

var (
	app = kingpin.New("asciimoji", "Look up ASCII emojis")

	listAll = app.Flag("list-all", "list all available emojis").Short('l').Bool()
	search  = app.Flag("search", "search all available emojis").Short('s').Strings()

	query = app.Arg("lookup", "the keyword to lookup.").String()
)

func main() {

	app.Version("0.1").VersionFlag.Short('V')
	app.HelpFlag.Short('h')
	app.UsageTemplate(kingpin.SeparateOptionalFlagsUsageTemplate)
	kingpin.MustParse(app.Parse(os.Args[1:]))

	if *listAll {
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

//list emojis with keywords that match a specified search term
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

//get the emoji(s) for a keyword
func getEmoji(query string) {

	emojiResults, ok := keywordLookUp[query]
	if !ok {
		fmt.Println("emoji not found :(")
		os.Exit(0)
	}

	var emoji string

	// if there is more than one emoji that matches the keyword, ask the user to pick which one they want
	if len(emojiResults) > 1 {
		var emojiList = make([]string, 0)
		for _, e := range emojiResults {
			emojiList = append(emojiList, e.Emoji)
		}

		var qs = []*survey.Question{
			{
				Name: "emoji_answer",
				Prompt: &survey.Select{
					Message: "Choose a emoji:",
					Options: emojiList,
				},
			},
		}

		answer := struct {
			EmojiAnswer string `survey:"emoji_answer"`
		}{}

		err := survey.Ask(qs, &answer)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		emoji = answer.EmojiAnswer

	} else {
		emoji = emojiResults[0].Emoji
	}

	// copy the selected emoji to the clipboard
	if err := clipboard.WriteAll(emoji); err != nil {
		panic(err)
	}
	fmt.Printf("copied %s\n", emoji)
}

//list all available emojis
func listEmojis() {
	table := tablewriter.NewWriter(os.Stdout)
	for _, emoji := range emojis {
		table.Append([]string{emoji.Emoji, strings.Join(emoji.Keywords, ", ")})
	}
	table.Render()
}

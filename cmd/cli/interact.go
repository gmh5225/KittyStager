package cli

import (
	"KittyStager/cmd/httpUtil"
	"KittyStager/cmd/util"
	"fmt"
	"github.com/c-bata/go-prompt"
	"os"
	"strings"
)

func Interact(kittenName string) error {
	in := fmt.Sprintf("KittyStager - %s❯ ", kittenName)
	for {
		t := prompt.Input(in, completerInteract,
			prompt.OptionPrefixTextColor(prompt.Blue),
			prompt.OptionPreviewSuggestionTextColor(prompt.Green),
			prompt.OptionSelectedSuggestionBGColor(prompt.LightGray),
			prompt.OptionSelectedSuggestionTextColor(prompt.Blue),
			prompt.OptionDescriptionBGColor(prompt.Blue),
			prompt.OptionSuggestionBGColor(prompt.DarkGray))
		input := strings.Split(t, " ")
		switch input[0] {
		case "exit":
			os.Exit(1337)
		case "back":
			return nil
		case "target":
			printTarget()
		case "interact":
			interact()
		case "payload":
			payload(kittenName)
		case "sleep":
			sleep(input, kittenName)
		case "recon":
			initChecks := httpUtil.Targets[kittenName].GetInitChecks()
			util.PrintRecon(initChecks)
		}
	}
	return nil
}

func completerInteract(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "exit", Description: "Exit the program"},
		{Text: "back", Description: "Go back to the main menu"},
		{Text: "target", Description: "Show targets"},
		{Text: "interact", Description: "Interact with a target"},
		{Text: "payload", Description: "Host a payload"},
		{Text: "sleep", Description: "Set sleep time"},
		{Text: "recon", Description: "Show recon information"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}
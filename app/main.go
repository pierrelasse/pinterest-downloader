package main

import (
	"encoding/json"
	"pinterest-downloader/app/api"
	"pinterest-downloader/app/utils"
)

func formatMap(data any) string {
	formatted, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return "{}"
	}
	return string(formatted)
}

func main() {
	pinID := "422281209475293"

	result, err := api.Suggestions(pinID, "")

	// result, err := api.GetPin(pinID)

	if err != nil {
		utils.Console_writeln(utils.Fmt("Error: %v", err))
		return
	}

	utils.Console_writeln(formatMap(result))
	// utils.Console_writeln(formatMap(result))

	utils.Console_writeln("============")
}

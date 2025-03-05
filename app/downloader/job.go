package downloader

import (
	"io"
	"os"
	"pinterest-downloader/app/fetch"
	"pinterest-downloader/app/utils"
)

type Job struct {
	URL      string
	FileName string
}

func (th *Job) Run() error {
	resp, err := fetch.Request(th.URL, "GET")
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	file, err := os.Create("out/" + th.FileName)
	if err != nil {
		return err
	}
	defer file.Close()

	file.Write(body)

	utils.Console_writeln(utils.FGreen + "✅ " + th.FileName + utils.Reset)

	return nil
}

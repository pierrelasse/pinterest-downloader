package main

import (
	"os"
	"path/filepath"
	"pinterest-downloader/app/api"
	"pinterest-downloader/app/downloader"
	"pinterest-downloader/app/utils"
	"sync"
)

type Ctx struct {
	Amount         uint16
	Images         bool
	Videos         bool
	Multiply       uint8
	MultiplyChance uint8
	discovered     *utils.ConcurrentStringSet
	worker         downloader.Worker
	discoverWG     sync.WaitGroup
}

func (th *Ctx) IsAmountReached() bool {
	var count int
	err := filepath.Walk("out", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			count++
		}
		return nil
	})
	if err != nil {
		return true
	}
	return count >= int(th.Amount)
}

func discover(bookmark string, pinId string, ctx *Ctx) {
	result, err := api.Suggestions(pinId, bookmark)
	if err != nil {
		utils.Console_writeln(utils.Fmt("Error: %v", err))
		return
	}

	for _, i := range result.Response {
		if ctx.IsAmountReached() && ctx.Multiply == 0 {
			break
		}

		if !ctx.discovered.Exists(i.ID) {
			ctx.discovered.Add(i.ID)

			if ctx.Videos && i.Video != nil {
				ctx.worker.Submit(downloader.Job{
					URL:      *i.Video,
					FileName: i.ID + ".mp4",
				})
			} else if ctx.Images && i.ImageURL != "" {
				ctx.worker.Submit(downloader.Job{
					URL:      i.ImageURL,
					FileName: i.ID + ".jpg",
				})
			} else {
				continue
			}
		}

		if ctx.Multiply > 0 && uint8(utils.Random_randInt(0, 100)) < ctx.MultiplyChance {
			ctx.Multiply--
			ctx.discoverWG.Add(1)
			go func(pinID string) {
				defer ctx.discoverWG.Done()
				// utils.Console_writeln(utils.Fmt("multiplied, remaining: %d", ctx.Multiply))
				discover("", pinID, ctx)
			}(i.ID)
		}
	}

	if !ctx.IsAmountReached() || ctx.Multiply > 0 {
		discover(result.Bookmark, pinId, ctx)
	}
}

func main() {
	pinId := "10344274145522856"

	ctx := &Ctx{
		Images:         true,
		Videos:         true,
		Amount:         uint16(500),
		Multiply:       8,
		MultiplyChance: 5,
		discovered:     utils.NewConcurrentStringSet(),
		worker:         downloader.Worker{},
	}
	go ctx.worker.StartWorkers(3)

	os.MkdirAll("out", 0)

	discover("", pinId, ctx)

	utils.Console_writeln("Waiting for workers")
	ctx.worker.Wait()

	utils.Console_writeln("Done")
}

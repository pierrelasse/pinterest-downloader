package downloader

import (
	"pinterest-downloader/app/utils"
	"sync"
	"time"
)

type Worker struct {
	jobs chan Job
	wg   *sync.WaitGroup
}

// Starts a specified amount of new workers.
func (th *Worker) StartWorkers(amount uint8) {
	if th.wg != nil {
		return
	}
	th.jobs = make(chan Job, 100)
	th.wg = &sync.WaitGroup{}

	for i := uint8(0); i < amount; i++ {
		th.wg.Add(1)
		go func() {
			defer th.wg.Done()

			for {
				select {
				case job, ok := <-th.jobs:
					if !ok {
						return
					}
					if err := job.Run(); err != nil {
						utils.Console_writeln(utils.FRed + utils.Fmt("Error running job: %v", err) + utils.Reset)
						time.Sleep(time.Second)
					}
				default:
					if th.wg == nil {
						return
					} else {
						time.Sleep(time.Second)
					}
				}
			}
		}()
	}
}

// Waits until all jobs are finished.
func (th *Worker) Wait() {
	if th.wg != nil {
		wg := th.wg
		th.wg = nil
		wg.Wait()
	}
}

// Submits a new job.
func (th *Worker) Submit(job Job) {
	if th.wg == nil {
		return
	}
	select {
	case th.jobs <- job:
	default:
	}
}

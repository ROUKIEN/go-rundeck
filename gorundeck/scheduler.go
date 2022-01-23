package gorundeck

import (
	"fmt"
	"time"
)

type AJob struct {
	Name             string
	Description      string
	IntervalDuration time.Duration
	Fn               func(input string, output string, startAt time.Time)
	OnGracefulStop   *func()
}

func NewScheduler() {
	j1 := AJob{
		Name:        "Do stuff",
		Description: "This job does some stuff",
		Fn: func(input, output string, startAt time.Time) {
			time.Sleep(2000 * time.Millisecond)
			fmt.Printf("\tI did some stuff !\n")
		},
		IntervalDuration: 10 * time.Second,
	}

	j2 := AJob{
		Name:        "Do another stuff",
		Description: "Computing some stuff",
		Fn: func(input, output string, startAt time.Time) {
			fmt.Printf("\tStarted at %s !\n", startAt.Format(time.RFC3339))
		},
		IntervalDuration: 5 * time.Second,
	}

	jobs := []*AJob{
		&j1,
		&j2,
	}

	done := make(chan bool)
	tickers := make([]*time.Ticker, 0)

	for _, job := range jobs {
		tickers = append(tickers, schedule(job, done))
	}

	<-done
}

func schedule(j *AJob, done <-chan bool) *time.Ticker {
	ticker := time.NewTicker(j.IntervalDuration)

	go func(j *AJob) {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Printf("[%s][%s] Starting...\n", t.Format(time.RFC3339), j.Name)
				j.Fn("coucou", "sortie", t)
				fmt.Printf("[%s][%s] Done.\n", time.Now().Format(time.RFC3339), j.Name)
			}
		}
	}(j)

	return ticker
}

package jobs

import (
	"fmt"
	"sync"

	"gop/sensibull/consts"
)

var jobQueue chan SubscribeJob

func Init() chan SubscribeJob {
	jobQueue = make(chan SubscribeJob, consts.DefaultNumOfJobs)
	return jobQueue
}

type SubscribeJob struct {
	ID     int
	Symbol string
	Token  int
	Action func(token int)
}
type Result struct {
	JobID int
	Data  interface{} // Replace this with the actual result data type.
}

func (job *SubscribeJob) DoJob() Result {
	// Simulate some work processing time (replace this with your actual job processing logic).
	job.Action(job.Token)
	resultData := fmt.Sprintf("Result data for job %d", job.ID)
	return Result{JobID: job.ID, Data: resultData}
}

func StartSubscribingChannel(quotesMap map[string]int, action func(token int)) {
	workerSize := consts.DefaultNumOfWorkers
	results := make(chan Result, len(quotesMap))
	var wg sync.WaitGroup
	for i := 1; i <= workerSize; i++ {
		wg.Add(1)
		go worker(i, jobQueue, results, &wg)
	}

	for key, value := range quotesMap {
		job := SubscribeJob{Symbol: key, Token: value, Action: action}
		jobQueue <- job
	}
	wg.Wait()
	close(results)
	for result := range results {
		fmt.Printf("Result received: %d\n", result)
	}
}

func worker(id int, jobs <-chan SubscribeJob, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		result := job.DoJob()
		fmt.Printf("Worker %d processing job %d\n", id, job.ID)
		results <- result // Sending the result to the results channel.
	}
}

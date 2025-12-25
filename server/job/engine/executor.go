package engine

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/qwganker/boring/comm/table"
	"github.com/qwganker/boring/conf"
)

type JobEngineExecutor struct {
}

func NewJobEngineExecutor() *JobEngineExecutor {
	return &JobEngineExecutor{}
}

func (e *JobEngineExecutor) Exec(jobtask *table.TJobTask) error {

	log.Printf("Exec jobTask ID:%d, SchedID:%s\n", jobtask.ID, jobtask.SchedID)

	body, err := json.Marshal(jobtask)

	httpReq, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/job/worker/process", conf.GetConfig().JobWorkerAddress), bytes.NewReader(body))
	if err != nil {
		log.Println(err.Error())
		return err
	}

	client := &http.Client{Timeout: 5 * time.Second}

	resp, err := client.Do(httpReq)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode/100 != 2 {
		s := fmt.Sprintf("Call JobWorker returned HTTP status %s", resp.Status)
		log.Println(s)
		return errors.New(s)
	}

	return nil
}

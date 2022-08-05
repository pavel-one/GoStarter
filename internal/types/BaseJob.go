package types

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

type BaseJob struct {
	ID   string
	Data interface{}
}

func (j *BaseJob) Init(data interface{}) {
	j.ID = uuid.NewString()
	j.Data = data
}

func (j *BaseJob) Run() {
	time.Sleep(time.Second * 5)
	fmt.Println(fmt.Sprintf("Handle test job id: %s", j.ID))

	defer j.Close()
}

func (j *BaseJob) Close() {
	fmt.Println(fmt.Sprintf("Close job %s", j.ID))
}

func (j *BaseJob) Error(err error) {
	fmt.Println(fmt.Sprintf("Error job %s", j.ID))
}

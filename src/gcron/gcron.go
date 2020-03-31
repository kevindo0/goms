package gcron

import (
	"fmt"
	"goms/pkg/logging"
	"github.com/robfig/cron"
)

func CleanAllTag() {
    logging.Info("Cleaned") 
}


type TestJob struct {
}

func (this TestJob)Run() {
    fmt.Println("testJob1...")
}

type Test2Job struct {
}

func (this Test2Job)Run() {
    fmt.Println("testJob2...")
}

func Cron() {
    fmt.Println("Starting...")
    i := 0
    c := cron.New()
    
    //AddFunc
    spec := "*/5 * * * * ?"
    c.AddFunc(spec, func() {
        i++
        fmt.Println("cron running:", i)
    })
    
    //AddJob方法
    c.AddJob(spec, TestJob{})
    c.AddJob(spec, Test2Job{})
    
    //启动计划任务
    c.Start()
}
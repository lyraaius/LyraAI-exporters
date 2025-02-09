package cron_job

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"runtime"
	"time"
)

type Job interface {
	Do(ctx context.Context)
	Interval() time.Duration
	DisableSerializable() bool
}

type CronJob struct {
	job Job
}

func NewCronJob(job Job) *CronJob {
	return &CronJob{
		job: job,
	}
}

func (c *CronJob) Run(ctx context.Context) {
	recoverFun := func() {
		if err := recover(); err != nil {
			hlog.CtxErrorf(ctx, "recover fail, error: %+v", err)
			buf := make([]byte, 2048)
			n := runtime.Stack(buf, false)
			hlog.CtxErrorf(ctx, "panic stack: %s", string(buf[:n]))
			fmt.Printf("error stack trace: %s \n", string(buf[:n]))
		}
	}
	defer recoverFun()
	ctx = context.Background()
	c.job.Do(ctx)

	ticker := time.Tick(c.job.Interval())

	for {
		select {
		case <-ticker:
			ctx = context.Background()

			if !c.job.DisableSerializable() {
				c.job.Do(ctx)
			} else {
				go func() {
					defer recoverFun()
					c.job.Do(ctx)
				}()
			}

		case <-ctx.Done():
			return
		}
	}
}

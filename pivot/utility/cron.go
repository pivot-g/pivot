package utility

import (
	"fmt"
	"time"

	"github.com/pivot-g/pivot/pivot/log"
	"github.com/robfig/cron"
)

type CronExpression struct {
	Spec string
}

func (c *CronExpression) ExpandCronExpression() (int64, error) {
	c.ValidateCron()
	p, err := cron.Parse(c.Spec)
	now := time.Now().In(time.Local)
	if err != nil {
		log.Debug("Unable to expand cron expression")
	}
	fmt.Println(err)
	fmt.Println(p)
	return p.Next(now).Unix(), err
}

func (c *CronExpression) ValidateCron() bool {
	o := true
	if len(c.Spec) < 5 {
		o = false
	}
	if len(c.Spec) == 5 {
		c.Spec = fmt.Sprintf("0 %s", c.Spec)
	}
	log.Debug("spec ", c.Spec)
	_, err := cron.Parse(c.Spec)
	if err != nil {
		log.Debug("Invalid cron expression")
		o = false
	}
	return o
}

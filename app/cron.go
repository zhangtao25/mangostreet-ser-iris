package app

import (
	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"

	//"mangostreet-ser-iris/services/task"
	"mangostreet-ser-iris/services/task"
)

func startSchedule() {
	c := cron.New()
	addCronFunc(c, "@every 30m", task.RssTask)
	addCronFunc(c, "@every 1h", task.SitemapTask)
	c.Start()
}

func addCronFunc(c *cron.Cron, sepc string, cmd func()) {
	err := c.AddFunc(sepc, cmd)
	if err != nil {
		logrus.Error(err)
	}
}

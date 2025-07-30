package cronpkg

import (
	"context"

	"github.com/robfig/cron/v3"
)

/*
Cron Spec Examples:

// ┌───────────── minute (0 - 59)
// │ ┌───────────── hour (0 - 23)
// │ │ ┌───────────── day of month (1 - 31)
// │ │ │ ┌───────────── month (1 - 12 or names)
// │ │ │ │ ┌───────────── day of week (0 - 6) (Sunday=0 or 7, or names)
// │ │ │ │ │
// │ │ │ │ │
// * * * * *  command to execute

Examples:

"@every 1m"           // Every 1 minute
"0 0 * * *"           // Every day at midnight
"30 14 * * 1-5"       // At 14:30 on Monday through Friday
"0 0 1 * *"           // At midnight on the first day of every month
"0 0 * * 0"           // Every Sunday at midnight
"15 10 * * *"         // Every day at 10:15 AM
"0 9-17 * * 1-5"      // Every hour from 9 AM to 5 PM Monday through Friday (at minute 0)
"@hourly"             // Every hour (same as "0 * * * *")
"@daily"              // Every day at midnight (same as "0 0 * * *")
"@weekly"             // Every week at Sunday midnight
"@monthly"            // Every month at midnight on day 1
"@yearly" or "@annually" // Every year at midnight on Jan 1

Note:
- You can use ranges (e.g. 1-5), lists (e.g. 1,2,5), and steps
- Names of months and days can be used (e.g. "JAN", "MON")
*/

type Manager struct {
	c *cron.Cron
}

func New() *Manager {
	return &Manager{
		c: cron.New(),
	}
}

func (m *Manager) AddJob(spec string, job func()) error {
	_, err := m.c.AddFunc(spec, job)
	return err
}

func (m *Manager) Start() {
	m.c.Start()
}

func (m *Manager) Stop() context.Context {
	return m.c.Stop()
}

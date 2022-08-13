package cron

func init() {
	_ = Initialize
}

// Initialize create a new instance of cron singleton.
func Initialize() {
	mutex.Lock()
	defer mutex.Unlock()

	Close()
	c = NewCron()
}

// Close remove cron singleton instance.
func Close() {
	if c != nil {
		c.Close()
	}
}

// Singleton return a single instance of cron.
func Singleton() *Cron {
	if c == nil {
		Initialize()
	}

	return c
}

package scheduler

import "time"

type Task struct {
	ID       string
	Name     string
	Interval time.Duration
	Function func()
	IsActive bool
	LastRun  time.Time
	NextRun  time.Time
}

func NewTask(id, name string, interval time.Duration, fn func()) *Task {
	return &Task{
		ID:       id,
		Name:     name,
		Interval: interval,
		Function: fn,
		IsActive: true,
	}
}
func (t *Task) UpdateInterval(newInterval time.Duration) {
	t.Interval = newInterval
}

func (t *Task) SetNextRun(nextRun time.Time) {
	t.NextRun = nextRun
}
func (t *Task) Run() {
	if t.IsActive {
		t.Function()
		t.LastRun = time.Now()
	}
}
func (t *Task) Stop() {
	t.IsActive = false
}

func (t *Task) Restart() {
	t.IsActive = true
}

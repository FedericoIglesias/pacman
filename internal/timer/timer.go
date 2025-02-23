package timer

type Timer struct {
	Sticks int
	Count  int
}

func NewTimer(milisecond int) *Timer {
	Sticks := int((float64(milisecond) * (0.03)))

	return &Timer{
		Sticks: Sticks,
		Count:  0,
	}
}

func (t *Timer) Update() {
	t.Count += 1
}

func (t *Timer) IsTimerDone() bool {
	return t.Sticks <= t.Count
}

func (t *Timer) Reset() {
	t.Count = 0
}

package jobqueue

var concurrencyLimit = Defaults.concurrencyLimit
var inProgress = 0
var Queue = make([](func()), 0, 700)

func SetConcurrecny(num int) {
	concurrencyLimit = num
}

func QueuePush(f func()) {
	Queue = append(Queue, f)
	if inProgress < concurrencyLimit && len(Queue) > 0 {
		inProgress++
		fn := Queue[0]
		Queue = Queue[1:]
		go fn()
	}
}

func QueueNext() {
	inProgress--
	if inProgress < concurrencyLimit && len(Queue) > 0 {
		inProgress++
		fn := Queue[0]
		Queue = Queue[1:]
		go fn()
	}
}

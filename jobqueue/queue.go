package jobqueue

var concurrencyLimit = Defaults.concurrencyLimit
var inProgress = 0

var queue = make([](func()), 0, 700)

//SetConcurrency is used to set limit on concuurent execution of task being pushed in the queue.
func SetConcurrency(num int) {
	concurrencyLimit = num
}

//QueuePush function is used to push a task to queue
//
//  task := func(num int){
//		println(num)
//  }
//
//  jobqueue.QueuePush(func(){
//		task(7)
//		jobqueue.QueueNext()
//  })
func QueuePush(f func()) {
	queue = append(queue, f)
	if inProgress < concurrencyLimit && len(queue) > 0 {
		inProgress++
		fn := queue[0]
		queue = queue[1:]
		go fn()
	}
}

//QueueNext function is used to send queue an acknowledgement when a task is completed.
func QueueNext() {
	inProgress--
	if inProgress < concurrencyLimit && len(queue) > 0 {
		inProgress++
		fn := queue[0]
		queue = queue[1:]
		go fn()
	}
}

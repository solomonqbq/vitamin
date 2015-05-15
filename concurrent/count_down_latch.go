package concurrent

type CountDownLatch struct {
	count       int
	signal_chan chan int
}

func NewCountDownLatch(count int) *CountDownLatch {
	cdl := new(CountDownLatch)
	cdl.count = count
	cdl.signal_chan = make(chan int, count)
	return cdl
}

func (cdl *CountDownLatch) CountDown() {
	cdl.signal_chan <- 1
}

func (cdl *CountDownLatch) Await() {
	for i := 0; i < cdl.count; i++ {
		<-cdl.signal_chan
	}
}

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	nodes       = 5
	transactions = 20000
	batchSize    = 20
)

type Txn struct {
	id int
}

type Vote struct {
	ok bool
}

func participant(id int, prepare <-chan []Txn, decision <-chan bool, vote chan<- Vote, ack chan<- bool, wg *sync.WaitGroup) {
	for batch := range prepare {
		_ = batch
		ok := rand.Intn(100) > 5
		vote <- Vote{ok: ok}
		<-decision
		ack <- true
	}
	wg.Done()
}

func coordinator() {
	prepare := make(chan []Txn)
	vote := make(chan Vote, nodes)
	decision := make(chan bool)
	ack := make(chan bool, nodes)

	var wg sync.WaitGroup

	for i := 0; i < nodes; i++ {
		wg.Add(1)
		go participant(i, prepare, decision, vote, ack, &wg)
	}

	queue := make([]Txn, 0, batchSize)
	start := time.Now()

	for i := 0; i < transactions; i++ {
		queue = append(queue, Txn{id: i})

		if len(queue) == batchSize {
			for j := 0; j < nodes; j++ {
				prepare <- queue
			}

			commit := true
			for j := 0; j < nodes; j++ {
				v := <-vote
				if !v.ok {
					commit = false
				}
			}

			for j := 0; j < nodes; j++ {
				decision <- commit
			}

			for j := 0; j < nodes; j++ {
				<-ack
			}

			queue = queue[:0]
		}
	}

	close(prepare)
	wg.Wait()

	fmt.Println("Elapsed:", time.Since(start))
}

func main() {
	coordinator()
}

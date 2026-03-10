import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	participants = 5
	transactions = 20000
)

type Vote struct {
	id   int
	yes  bool
}

type Ack struct {
	id int
}

func participant(id int, prepare <-chan int, decision <-chan bool, vote chan<- Vote, ack chan<- Ack, wg *sync.WaitGroup) {
	for range prepare {
		ok := rand.Intn(100) > 5
		vote <- Vote{id: id, yes: ok}
		d := <-decision
		if d {
			_ = id
		}
		ack <- Ack{id: id}
	}
	wg.Done()
}

func coordinator() {
	prepare := make(chan int)
	vote := make(chan Vote, participants)
	decision := make(chan bool)
	ack := make(chan Ack, participants)

	var wg sync.WaitGroup

	for i := 0; i < participants; i++ {
		wg.Add(1)
		go participant(i, prepare, decision, vote, ack, &wg)
	}

	start := time.Now()

	for t := 0; t < transactions; t++ {
		for i := 0; i < participants; i++ {
			prepare <- t
		}

		commit := true
		for i := 0; i < participants; i++ {
			v := <-vote
			if !v.yes {
				commit = false
			}
		}

		for i := 0; i < participants; i++ {
			decision <- commit
		}

		for i := 0; i < participants; i++ {
			<-ack
		}
	}

	close(prepare)
	wg.Wait()

	fmt.Println("Elapsed:", time.Since(start))
}

func main() {
	coordinator()
}

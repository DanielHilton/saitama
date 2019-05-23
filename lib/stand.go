package star_platinum

import (
	"fmt"
	//"net/http"
	"time"
)

type Stand struct {
	Endpoint  string
	Frequency int
	attacking bool
}

func NewStand(endpoint string, frequency int) *Stand {
	return &Stand{
		endpoint,
		frequency,
		false,
	}
}

func (s Stand) StartAttacking() {
	s.attacking = true
	for s.attacking {
		go requestEndpoint(s.Endpoint)
		time.Sleep(time.Duration(1/s.Frequency) * time.Second)
	}
}

func (s Stand) StopAttacking() {
	s.attacking = false
}

func requestEndpoint(e string) {
	fmt.Println("ORA ORA ORA ORA")

	//resp, _ := http.Get(e)
	//fmt.Println(resp)
}

package star_platinum

import "fmt"

type Stand struct {
	Endpoint  string
	Frequency int
}

func NewStand(endpoint string, frequency int) *Stand {
	return &Stand{
		endpoint,
		frequency,
	}
}

func Attack(s Stand) {
	go requestEndpoint(s.Endpoint)
}

func requestEndpoint(e string) {
	fmt.Println("ORA ORA ORA ORA")
}

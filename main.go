package main

import (
	"encoding/csv"
	"fmt"
	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
	"os"
)

func FileParserToTickets(f *os.File, arrTickets *[]tickets.Ticket) {
	csvReader := csv.NewReader(f)
	for {
		rec, err := csvReader.Read()
		if err != nil && err.Error() == "EOF" {
			break
		}
		t := tickets.GenerateTicket(rec)
		*arrTickets = append(*arrTickets, t)
	}
}

func main() {
	arrTickets := []tickets.Ticket{}
	f, err := os.Open("/Users/nstrangis/GoBoot/desafio-go-bases/tickets.csv")
	if err != nil {
		println(err.Error())
		return
	}
	defer f.Close()

	FileParserToTickets(f, &arrTickets)

	//REQ1
	total := tickets.GetTotalTickets(&arrTickets, "Brazil")
	fmt.Println(total)

	//REQ2
	totalHora, err := tickets.GetTicketsByPeriod(&arrTickets, "05")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(totalHora)
	}

	//REQ3
	average, err := tickets.AverageDestination(&arrTickets, "Argentina", len(arrTickets))
	fmt.Println(average)
}

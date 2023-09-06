package tickets

import (
	"errors"
	"strconv"
)

type Ticket struct {
	Id          string
	Name        string
	Email       string
	Destination string
	Time        string
	Tickets     []Ticket
}

const (
	_pos_id          = 0
	_pos_name        = 1
	_pos_email       = 2
	_pos_destination = 3
	_pos_time        = 4
)

const (
	_early = iota
	_morning
	_late
	_night
)

func GenerateTicket(datos []string) Ticket {
	t := Ticket{}
	t.Id = datos[_pos_id]
	t.Name = datos[_pos_name]
	t.Email = datos[_pos_email]
	t.Destination = datos[_pos_destination]
	t.Time = datos[_pos_time]
	return t
}

// ejemplo 1
func GetTotalTickets(arrTickets *[]Ticket, destination string) int {
	count := 0
	for _, t := range *arrTickets {
		if t.Destination == destination {
			count++
		}
	}
	return count
}

// ejemplo 2
func GetPeriod(hour int) int {
	switch {
	case hour >= 0 && hour < 6:
		return _early
	case hour >= 6 && hour < 12:
		return _morning
	case hour >= 12 && hour < 18:
		return _late
	default:
		return _night
	}
}

func GetTicketsByPeriod(arrTickets *[]Ticket, time string) (int, error) {
	var period, count int
	hour, err := strconv.Atoi(time[:2])
	count = 0
	period = GetPeriod(int(hour))
	for _, t := range *arrTickets {
		var ticketHour int
		ticketHour, err = strconv.Atoi(t.Time[:2])
		if period == GetPeriod(ticketHour) {
			count++
		}
	}
	return count, err
}

// ejemplo 3
func AverageDestination(arrTickets *[]Ticket, destination string, total int) (float64, error) {
	count := 0
	if total == 0 {
		return 0, errors.New("No hay tickets")
	}
	for _, t := range *arrTickets {
		if destination == t.Destination {
			count++
		}
	}
	return float64(count) / float64(total) * 100, nil
}

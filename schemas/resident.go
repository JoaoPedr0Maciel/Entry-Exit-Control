package schemas

import (
	"time"
)

type Resident struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Sexo        string    `json:"sexo"`
	Cpf         string    `json:"cpf"`
	HouseNumber int64     `json:"house_number"`
	EntryTime   time.Time `json:"entry_time"`
	ExitTime    time.Time `json:"exit_time"`
}

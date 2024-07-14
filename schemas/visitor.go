package schemas

import (
	"time"
)

type Visitor struct {
	ID           uint       `json:"id"`
	Name         string     `json:"name"`
	Sexo         string     `json:"sexo"`
	Cpf          string     `json:"cpf"`
	HouseToVisit int64      `json:"house_to_visit_number"`
	WhoAllowed   string      `json:"who_allowed_cpf"`
	EntryTime    time.Time  `json:"entry_time"`
	ExitTime     *time.Time `json:"exit_time"`
}

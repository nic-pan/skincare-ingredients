package models

type Contradiction struct {
	ID       int    `json:id`
	Effects  Effect `json:effects`
	Reason   string `json:reason`
	Severity string `json:severity`
}

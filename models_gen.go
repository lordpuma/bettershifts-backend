// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package bettershifts

import (
	"github.com/lordpuma/bettershifts/models"
)

type EditBenefit struct {
	Date   *string `json:"date"`
	Reason *string `json:"reason"`
	User   *string `json:"user"`
	Amount *int    `json:"amount"`
}

type EditShift struct {
	Start     *string `json:"start"`
	End       *string `json:"end"`
	Workplace *string `json:"workplace"`
	User      *string `json:"user"`
}

type EditTodo struct {
	Name      *string `json:"name"`
	Workplace *string `json:"workplace"`
	Date      *string `json:"date"`
	Benefit   *int    `json:"benefit"`
}

type EditUser struct {
	FirstName  *string   `json:"firstName"`
	LastName   *string   `json:"lastName"`
	Username   *string   `json:"username"`
	IsAdmin    *bool     `json:"isAdmin"`
	Wage       *int      `json:"wage"`
	Workplaces []*string `json:"workplaces"`
}

type EditWorkplace struct {
	Name *string `json:"name"`
}

type LoginPayload struct {
	Token       string `json:"token"`
	HasPassword bool   `json:"hasPassword"`
}

type NewBenefit struct {
	Date   string `json:"date"`
	Reason string `json:"reason"`
	User   string `json:"user"`
	Amount int    `json:"amount"`
}

type NewShift struct {
	Start     string  `json:"start"`
	End       *string `json:"end"`
	Workplace string  `json:"workplace"`
	User      string  `json:"user"`
}

type NewTodo struct {
	Name      string `json:"name"`
	Workplace string `json:"workplace"`
	Date      string `json:"date"`
	Benefit   int    `json:"benefit"`
}

type NewUser struct {
	FirstName  string   `json:"firstName"`
	LastName   string   `json:"lastName"`
	Username   string   `json:"username"`
	IsAdmin    bool     `json:"isAdmin"`
	Wage       int      `json:"wage"`
	Workplaces []string `json:"workplaces"`
}

type NewWorkplace struct {
	Name string `json:"name"`
}

type Wage struct {
	User   models.User `json:"user"`
	Amount int         `json:"amount"`
}

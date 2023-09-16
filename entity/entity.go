package entity

import (
	"time"
)

type ExpenseID int
type UserID int
type AccountTypeID int

type Expenses []*Expense
type Expense struct {
	ID            ExpenseID     `json:"id"`
	Title         string        `json:"title"`
	Expense       int           `json:"expense"`
	UserID        UserID        `json:"user_id"`
	AccountTypeID AccountTypeID `json:"account_type_id"`
	ImageURL      string        `json:"image_url"`
	Created       time.Time     `json:"created"`
	Modified      time.Time     `json:"modified"`
}

type ExpenseDetail struct {
	ID      ExpenseID `json:"id"`
	Content string    `json:"content"`
}

type User struct {
	ID        UserID `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type AccountType struct {
	ID    AccountTypeID `json:"id"`
	Title int           `json:"title"`
}

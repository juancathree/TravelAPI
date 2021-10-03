package domain

import "strings"

type Expenses map[string]UserExpense

type UserExpense struct {
	Amount  float64   `json:"amount" bson:"amount"`
	Expense []Expense `json:"expense" bson:"expense"`
}

type Expense struct {
	Amount      float64 `validate:"required" json:"amount" bson:"amount"`
	IsPersonal  bool    `validate:"required" json:"isPersonal" bson:"isPersonal"`
	Category    string  `json:"category" bson:"category"`
	Description string  `validate:"required,max=30" json:"description" bson:"description"`
}

func (e *Expenses) Initialize(userID *string) {
	*e = make(Expenses)
	userExpense := UserExpense{
		Amount:  0.0,
		Expense: []Expense{},
	}
	(*e)[strings.Split(*userID, "@")[0]] = userExpense
}

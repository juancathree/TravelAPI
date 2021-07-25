package domain

type Expenses map[string]UserExpense

type UserExpense struct {
	Amount  float64   `json:"amount" bson:"amount"`
	Expense []Expense `json:"expense" bson:"expense"`
}

type Expense struct {
	Amount      float64 `validate:"required" json:"amount" bson:"amount"`
	IsPersonal  bool    `validate:"required" json:"isPersonal" bson:"isPersonal"`
	Category    string  `validate:"required" json:"category" bson:"category"`
	Description string  `validate:"required,max=30" json:"description" bson:"description"`
}

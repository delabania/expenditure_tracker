package db

type Database interface {
	SavePerson(person Person) error
	SaveExpenditure(expenditure Expenditure) error
	SaveSettlement(settlement Settlement) error

	GetPerson(id int64) (Person, error)
	GetExpenditure(id int64) (Expenditure, error)
	GetSettlement(id int64) (Settlement, error)

	GetAllPersons() ([]Person, error)
	GetAllExpenditures() ([]Expenditure, error)
	GetAllSettlements() ([]Settlement, error)
}

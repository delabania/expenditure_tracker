package db

type Database interface {
	SavePerson(person Person) (int64, error)
	SaveExpenditure(expenditure Expenditure) (int64, error)
	SaveSettlement(settlement Settlement) (int64, error)

	GetPerson(id int64) (Person, error)
	GetExpenditure(id int64) (Expenditure, error)
	GetSettlement(id int64) (Settlement, error)

	GetAllPersons() ([]Person, error)
	GetAllExpenditures() ([]Expenditure, error)
	GetAllSettlements() ([]Settlement, error)
}

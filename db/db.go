package db

type Database interface {
	savePerson(person Person) error
	saveExpenditure(expenditure Expenditure) error
	saveSettlement(settlement Settlement) error

	getPerson(id int64) (Person, error)
	getExpenditure(id int64) (Expenditure, error)
	getSettlement(id int64) (Settlement, error)

	getAllPersons() ([]Person, error)
	getAllExpenditures() ([]Expenditure, error)
	getAllSettlements() ([]Settlement, error)
}
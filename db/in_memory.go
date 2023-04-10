package db

import "errors"

type InMemoryDatabase struct {
	Persons      []Person
	Expenditures []Expenditure
	Settlements  []Settlement
}

func (db *InMemoryDatabase) SavePerson(person Person) error {
	db.Persons = append(db.Persons, person)
	return nil
}

func (db *InMemoryDatabase) SaveExpenditure(expenditure Expenditure) error {
	db.Expenditures = append(db.Expenditures, expenditure)
	return nil
}

func (db *InMemoryDatabase) SaveSettlement(settlement Settlement) error {
	db.Settlements = append(db.Settlements, settlement)
	return nil
}

func (db *InMemoryDatabase) GetPerson(id int64) (Person, error) {
	for _, person := range db.Persons {
		if person.ID == id {
			return person, nil
		}
	}
	return Person{}, errors.New("person not found")
}

func (db *InMemoryDatabase) GetExpenditure(id int64) (Expenditure, error) {
	for _, expenditure := range db.Expenditures {
		if expenditure.ID == id {
			return expenditure, nil
		}
	}
	return Expenditure{}, errors.New("expenditure not found")
}

func (db *InMemoryDatabase) GetSettlement(id int64) (Settlement, error) {
	for _, settlement := range db.Settlements {
		if settlement.ID == id {
			return settlement, nil
		}
	}
	return Settlement{}, errors.New("settlement not found")
}

func (db *InMemoryDatabase) GetAllPersons() ([]Person, error) {
	return db.Persons, nil
}

func (db *InMemoryDatabase) GetAllExpenditures() ([]Expenditure, error) {
	return db.Expenditures, nil
}

func (db *InMemoryDatabase) GetAllSettlements() ([]Settlement, error) {
	return db.Settlements, nil
}

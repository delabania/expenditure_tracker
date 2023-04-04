package db

import "errors"

type InMemoryDatabase struct {
	Persons      []Person
	Expenditures []Expenditure
	Settlements  []Settlement
}

func (db *InMemoryDatabase) savePerson(person Person) error {
	db.Persons = append(db.Persons, person)
	return nil
}

func (db *InMemoryDatabase) saveExpenditure(expenditure Expenditure) error {
	db.Expenditures = append(db.Expenditures, expenditure)
	return nil
}

func (db *InMemoryDatabase) saveSettlement(settlement Settlement) error {
	db.Settlements = append(db.Settlements, settlement)
	return nil
}

func (db *InMemoryDatabase) getPerson(id int64) (Person, error) {
	for _, person := range db.Persons {
		if person.ID == id {
			return person, nil
		}
	}
	return Person{}, errors.New("person not found")
}

func (db *InMemoryDatabase) getExpenditure(id int64) (Expenditure, error) {
	for _, expenditure := range db.Expenditures {
		if expenditure.ID == id {
			return expenditure, nil
		}
	}
	return Expenditure{}, errors.New("expenditure not found")
}

func (db *InMemoryDatabase) getSettlement(id int64) (Settlement, error) {
	for _, settlement := range db.Settlements {
		if settlement.ID == id {
			return settlement, nil
		}
	}
	return Settlement{}, errors.New("settlement not found")
}

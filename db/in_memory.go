package db

import "errors"

type InMemoryDatabase struct {
	Persons      []Person
	Expenditures []Expenditure
	Settlements  []Settlement
}

func (db *InMemoryDatabase) SavePerson(person Person) (int64, error) {
	personId, err := db.getNextPersonID()
	if err != nil {
		return 0, err
	}
	person.ID = personId
	db.Persons = append(db.Persons, person)
	return person.ID, nil
}

func (db *InMemoryDatabase) SaveExpenditure(expenditure Expenditure) (int64, error) {
	expenditureId, err := db.getNextExpenditureID()
	if err != nil {
		return 0, err
	}
	expenditure.ID = expenditureId
	db.Expenditures = append(db.Expenditures, expenditure)
	return expenditure.ID, nil
}

func (db *InMemoryDatabase) SaveSettlement(settlement Settlement) (int64, error) {
	settlementId, err := db.getNextSettlementID()
	if err != nil {
		return 0, err
	}
	settlement.ID = settlementId
	db.Settlements = append(db.Settlements, settlement)
	return settlement.ID, nil
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

func (db *InMemoryDatabase) getNextPersonID() (int64, error) {
	var max int64 = 0
	for _, person := range db.Persons {
		if person.ID > max {
			max = person.ID
		}
	}

	return max, nil
}

func (db *InMemoryDatabase) getNextExpenditureID() (int64, error) {
	var max int64 = 0
	for _, expenditure := range db.Expenditures {
		if expenditure.ID > max {
			max = expenditure.ID
		}
	}

	return max, nil
}

func (db *InMemoryDatabase) getNextSettlementID() (int64, error) {
	var max int64 = 0
	for _, settlement := range db.Settlements {
		if settlement.ID > max {
			max = settlement.ID
		}
	}

	return max, nil
}

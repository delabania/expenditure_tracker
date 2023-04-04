package db

import (
	"reflect"
	"testing"
)

func TestInMemoryDatabase_getAllExpenditures(t *testing.T) {
	type fields struct {
		Persons      []Person
		Expenditures []Expenditure
		Settlements  []Settlement
	}
	tests := []struct {
		name    string
		fields  fields
		want    []Expenditure
		wantErr bool
	}{
		{
			"test all expenditures are returned",
			fields{[]Person{}, []Expenditure{{ID: 1}, {ID: 2}}, []Settlement{}},
			[]Expenditure{{ID: 1}, {ID: 2}},
			false,
		}, {
			"test empty expenditures are returned",
			fields{[]Person{}, []Expenditure{}, []Settlement{}},
			[]Expenditure{},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &InMemoryDatabase{
				Persons:      tt.fields.Persons,
				Expenditures: tt.fields.Expenditures,
				Settlements:  tt.fields.Settlements,
			}
			got, err := db.getAllExpenditures()
			if (err != nil) != tt.wantErr {
				t.Errorf("getAllExpenditures() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAllExpenditures() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInMemoryDatabase_getExpenditure(t *testing.T) {
	type fields struct {
		Persons      []Person
		Expenditures []Expenditure
		Settlements  []Settlement
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Expenditure
		wantErr bool
	}{
		{
			"test expenditure is returned",
			fields{[]Person{}, []Expenditure{{ID: 1}, {ID: 2}}, []Settlement{}},
			args{1},
			Expenditure{ID: 1},
			false,
		}, {
			"test error is returned",
			fields{[]Person{}, []Expenditure{{ID: 1}, {ID: 2}}, []Settlement{}},
			args{3},
			Expenditure{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &InMemoryDatabase{
				Persons:      tt.fields.Persons,
				Expenditures: tt.fields.Expenditures,
				Settlements:  tt.fields.Settlements,
			}
			got, err := db.getExpenditure(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("getExpenditure() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getExpenditure() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInMemoryDatabase_saveExpenditure(t *testing.T) {
	type fields struct {
		Persons      []Person
		Expenditures []Expenditure
		Settlements  []Settlement
	}
	type args struct {
		expenditure Expenditure
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		result  []Expenditure
	}{
		{
			"test expenditure is saved",
			fields{[]Person{}, []Expenditure{}, []Settlement{}},
			args{Expenditure{ID: 1}},
			false,
			[]Expenditure{{ID: 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &InMemoryDatabase{
				Persons:      tt.fields.Persons,
				Expenditures: tt.fields.Expenditures,
				Settlements:  tt.fields.Settlements,
			}
			if err := db.saveExpenditure(tt.args.expenditure); (err != nil) != tt.wantErr {
				t.Errorf("saveExpenditure() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(db.Expenditures, tt.result) {
				t.Errorf("saveExpenditure() result = %v, want %v", db.Expenditures, tt.result)
			}
		})
	}
}

package relational

import "gorm.io/gorm"

type Model any

type TableI[M Model] interface {
	Create(row M) error
	Select(query string, args any) (result []M)
	Update(query string, newRow M, args ...any) error
	Delete(query string, args... any) error
}

type Table[M Model] struct {
	db *gorm.DB
}

func NewTable[M Model] (db *gorm.DB) *Table[M] { return &Table[M]{ db } }

func (t Table[M]) Create(row M) error {
	result := t.db.Create(row)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t Table[M]) Select(query string, args ...any) (result []M) {
	t.db.Where(query, args...).Find(&result)
	return result
}

func (t Table[M]) Update(oldRow M, newRow M) error{
	result := t.db.Model(oldRow).Updates(&newRow)
	return result.Error
}

func (t Table[M]) Delete(query string, args... any) error {
	result := t.db.Delete(query, args...)
	if result.Error != nil{
		return result.Error
	}
	return nil
}
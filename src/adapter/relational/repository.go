package relational

import (
	"fmt"

	"gorm.io/gorm"
)

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

func NewTable[M Model] (db *gorm.DB) *Table[M] { 
	return &Table[M]{ db } 
}

func (t Table[M]) Create(row M) error {
	result := t.db.Create(row)
	fmt.Println(result.RowsAffected)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t Table[M]) Select(orderTable, query string, args ...any) (result []M) {
	t.db.Where(query, args...).Order(orderTable).Limit(10).Find(&result)
	return result
}

func (t Table[M]) Update(oldRow string, newRow M) error{
	result := t.db.Where("id = ?", oldRow).Updates(&newRow)
	return result.Error
}

func (t Table[M]) Delete(id string, m M) error {
	result := t.db.Where("id = ?", id).Delete(m)
	if result.Error != nil{
		return result.Error
	}
	return nil
}
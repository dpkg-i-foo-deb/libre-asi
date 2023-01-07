// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import (
	"time"
)

const TableNameDrugTranslation = "drug_translations"

// DrugTranslation mapped from table <drug_translations>
type DrugTranslation struct {
	Drug      string    `gorm:"column:drug" json:"drug"`
	Language  string    `gorm:"column:language;primaryKey" json:"language"`
	IsDefault bool      `gorm:"column:is_default;not null" json:"is_default"`
	DrugName  string    `gorm:"column:drug_name;primaryKey" json:"drug_name"`
	CreatedAt time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
}

// TableName DrugTranslation's table name
func (*DrugTranslation) TableName() string {
	return TableNameDrugTranslation
}

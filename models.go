package main

type DynamoDB struct {
	HashKey   string `gorm:"not null"`
	SortKey   string `gorm:"not null"`
	Value     string `gorm:"not null"`
	Tombstone bool   `gorm:"not null;default:false"`
}

func (DynamoDB) TableName() string {
	return "dynamodb"
}

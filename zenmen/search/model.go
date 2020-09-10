package search

type Search struct {
	ItemID        uint   `gorm:"column:item_id" json:"item_id"`
	Item          string `gorm:"column:item" json:"item"`
	Title         string `gorm:"column:title" json:"title"`
	Brief         string `gorm:"column:brief;size:500" json:"brief"` // 简介
	Description   string `gorm:"column:description;type:text" json:"description"`
	Tags          string `gorm:"column:tags;type:text" json:"tags"`
	SpecialColumn string `gorm:"column:special_column;size:50" json:"special_column"`
}

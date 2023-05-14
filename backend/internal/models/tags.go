package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	ID   uuid.UUID
	Name string
}

type TagModel struct {
}

func NewTagModel() *TagModel {
	return &TagModel{}
}

func (m *TagModel) CreateTag(db *gorm.DB, tag *Tag) error {
	if result := db.Create(tag); result.Error != nil {
		return result.Error
	}
	return nil
}

func (m *TagModel) GetTagByID(db *gorm.DB, tagID uuid.UUID) (*Tag, error) {
	target := &Tag{
		ID: tagID,
	}
	if result := db.First(&target); result.Error != nil {
		return nil, result.Error
	}
	return target, nil
}

func (m *TagModel) UpdateTag(db *gorm.DB, tagID uuid.UUID, tagName string) (*Tag, error) {
	target := &Tag{ID: tagID}
	if result := db.Model(target).Updates(Tag{Name: tagName}).First(target); result.Error != nil {
		return nil, result.Error
	}
	return target, nil
}

func (m *TagModel) DeleteTag(db *gorm.DB, tagID uuid.UUID) (*Tag, error) {
	target := &Tag{ID: tagID}
	if result := db.Delete(&target); result.Error != nil {
		return nil, result.Error
	}
	db.Unscoped().Find(&target)
	return target, nil
}

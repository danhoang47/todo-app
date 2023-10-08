package common

import "time"

type SQLModel struct {
	ID        uint       `json:"-" gorm:"column:id;"`
	FakeID    *UID       `json:"id" gorm:"-"`
	Status    int8       `json:"status" gorm:"column:status;default:1"`
	CreatedAt *time.Time `json:"createdAt,omitempty" gorm:"column:created_at;type:timestamp"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty" gorm:"column:updated_at;type:timestamp"`
}

func (m *SQLModel) GenUID(dbType int) {
	uid := NewUID(uint32(m.ID), dbType, 1)
	m.FakeID = &uid
}

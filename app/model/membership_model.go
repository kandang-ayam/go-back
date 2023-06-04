package model

import (
	"gorm.io/gorm"
	"strings"
	"time"
)

type Membership struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	MemberCode string    `json:"member_code"`
	Phone      int       `json:"phone"`
	BirthDay   string    `json:"birth_day"`
	Level      string    `json:"level"`
	Point      int       `json:"point"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (m *Membership) AfterFind(tx *gorm.DB) (err error) {
	m.BirthDay = strings.Split(m.BirthDay, "T")[0]

	return
}

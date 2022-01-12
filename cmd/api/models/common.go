//models.common

package models

import "time"

// ModelBase of fields for models.
type ModelBase struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement;<-:false"`
	CreatedAt time.Time `json:"create_at" gorm:"<-:false"`
	UpdatedAt time.Time `json:"update_at" gorm:"<-:false"`
	DeleteAt  time.Time `json:"delete_at" gorm:"index; <-:false"`
}

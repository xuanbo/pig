package entity

import (
	"github.com/xuanbo/pig/util"

	"gorm.io/gorm"
)

// Entity 实体
type Entity struct {
	ID        string     `json:"id" gorm:"primaryKey;type:string;size:30"`
	CreatedAt *util.Time `json:"createdAt" gorm:"<-:create"`
	UpdatedAt *util.Time `json:"updatedAt" gorm:"<-:create;<-:update"`
	CreatedBy string     `json:"createdBy" gorm:"<-:create;type:string;size:30"`
	UpdatedBy string     `json:"updatedBy" gorm:"<-:create;<-:update;type:string;size:30"`
}

// BeforeCreate 创建前
func (entity *Entity) BeforeCreate(tx *gorm.DB) error {
	ctx := tx.Statement.Context
	user := ctx.Value("userId")
	if user != nil {
		if userID, ok := user.(string); ok {
			entity.CreatedBy = userID
		}
	}
	// 创建时，未指定ID则生成一个
	if entity.ID == "" {
		entity.ID = GenerateID()
	}
	return nil
}

// BeforeUpdate 更新前
func (entity *Entity) BeforeUpdate(tx *gorm.DB) error {
	ctx := tx.Statement.Context
	user := ctx.Value("userId")
	if user != nil {
		if userID, ok := user.(string); ok {
			entity.UpdatedBy = userID
		}
	}
	return nil
}

package model

import "gorm.io/gorm"

type UpgradeResource struct {
	gorm.Model
	SourceName string
}

package model

import "gorm.io/gorm"

type Geoip struct {
	gorm.Model
}

func (m *Geoip) TableName() string {
    return "geoip"
}

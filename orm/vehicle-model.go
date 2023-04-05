package orm

import orm "github.com/HinekoTech/middleware/orm"

type VehicleModel struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	Name           string `db:"name"  json:"name" gorm:"not null;type:varchar(50)"`
	VehicleBrandID string `db:"vehicle_brand_id"  json:"vehicle_brand_id" gorm:"type:varchar(36)"`
}

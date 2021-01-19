package model

import "time"

//AppModel
type AppModel struct {
	Id       int `gorm:"primary_key"`
	Created  time.Time
	Modified time.Time
}

func (m *AppModel) Validate() {

}

func Create() {

}

func Get() {

}

func Update() {

}

func Delete() {

}

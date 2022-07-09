package model

import "gorm.io/gorm"

type Call struct {
	gorm.Model
	Time       int64
	StatusCode int
	Successful bool
	Url        Url
	UrlID      uint
}

type CallUrlResult struct {
	Id         uint
	StatusCode int
	Threshhold int
	ResetTime  int64
	Time       int64
}

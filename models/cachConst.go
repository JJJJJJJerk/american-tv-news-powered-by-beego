package models

import (
	"time"
)

//http://jinzhu.me/gorm/ gorm 文档

const CK_QUOTE = "ck.base.quotes.5random"
const CK_IMG_5_RANDOM = "ck.base.images.5random"
const CK_TAG_ALL = "ck.base.tags.all"

const C_EXPIRE_TIME_FOREVER = -1
const C_EXPIRE_TIME_YEAR = time.Hour * 24 * 365
const C_EXPIRE_TIME_WEEK = time.Hour * 24 * 7
const C_EXPIRE_TIME_DAY = time.Hour * 24
const C_EXPIRE_TIME_HOUR_12 = time.Hour * 12
const C_EXPIRE_TIME_HOUR_06 = time.Hour * 6
const C_EXPIRE_TIME_HOUR_03 = time.Hour * 3
const C_EXPIRE_TIME_HOUR_01 = time.Hour * 1
const C_EXPIRE_TIME_MIN_30 = time.Minute * 30
const C_EXPIRE_TIME_MIN_15 = time.Minute * 15
const C_EXPIRE_TIME_MIN_01 = time.Minute * 1

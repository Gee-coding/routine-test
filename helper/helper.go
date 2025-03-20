package helper

import "time"

var (
	loc, _  = time.LoadLocation("Asia/Bangkok")
	timeLog = time.Now().In(loc)
	TimeInc = time.Hour * 7
	made    = "test" // prod, test
)

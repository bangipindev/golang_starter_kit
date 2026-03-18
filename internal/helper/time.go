package helper

import "time"

var jakartaLoc *time.Location

func init() {
	var err error
	jakartaLoc, err = time.LoadLocation("Asia/Jakarta")
	if err != nil {
		panic("failed to load Asia/Jakarta timezone: " + err.Error())
	}
}

func ToWIB(t time.Time) time.Time {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	return t.In(loc)
}

func ToWIBString(t time.Time) string {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	return t.In(loc).Format("2006-01-02 15:04:05")
}

func NowWIB() time.Time {
	return time.Now().In(time.FixedZone("WIB", 7*60*60))
}

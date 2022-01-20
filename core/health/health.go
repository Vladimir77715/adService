package health

import (
	"time"
)

var ServiceStartTime time.Time

func Init() {
	ServiceStartTime = time.Now()
}

func GetServiceUptime() string {
	return time.Since(ServiceStartTime).Round(time.Second).String()
}

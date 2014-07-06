package main

import (
	"fmt"
	"time"
)

func UpdateApiCounters() {
	key := fmt.Sprintf("%d%d%d", time.Now().Month(), time.Now().Day(), time.Now().Year())
	stat := ApiStat{}
	DB.FirstOrCreate(&stat, ApiStat{DateKey: key})
	DB.Exec("UPDATE api_stats SET counter = counter + 1 WHERE date_key = ?", key)
}

package config

import "time"

type UserCollect struct {
	Collector string        `json:"collector"`
	Duration  time.Duration `json:"duration"`
}

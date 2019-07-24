package config

import (
	"time"
)


var (
	Redis_Max_Active int = 0
	Redis_Max_Idle int = 16
	Redis_Idle_Timeout = 3*time.Second
	Redis_Address = "localhost:6379"
)
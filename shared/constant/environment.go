package constant

import (
	"os"
	"strconv"
)

var (
	DatabaseHost = os.Getenv("APP_DB_HOST")
	DatabasePort = os.Getenv("APP_DB_PORT")
	DatabaseUser = os.Getenv("APP_DB_USERNAME")
	DatabasePass = os.Getenv("APP_DB_PASSWORD")
	DatabaseName = os.Getenv("APP_DB_NAME")

	AppSecretKey = os.Getenv("APP_SECRET_KEY")
	AppIssuer    = os.Getenv("APP_ISSUER")
	Timeout      = os.Getenv("APP_TIMEOUT_LIMIT")

	UserId         = "UserId"
	DefaultTimeout = 5
)

func GetTimeout() int {
	num, err := strconv.Atoi(Timeout)
	if err != nil {
		return DefaultTimeout
	}
	return num
}

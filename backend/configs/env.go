package configs

import "os"

var (
	DataSourceName = os.Getenv("DATA_SOURCE_NAME")
)

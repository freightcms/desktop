package logging

import (
	"os"

	ll "github.com/freightcms/logging"
)

var (
	Logger = ll.New(os.Stdout, ll.LogLevels())
)

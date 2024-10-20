package tasker

import (
	"runtime"

	"path/filepath"
)

var (
	// Retrieves the caller's information (file, line, etc.) at runtime.
	_, b, _, _ = runtime.Caller(0)

	// RootPath represents the application's root directory path.
	RootPath = filepath.Join(
		filepath.Dir(b), "/",
	)
)

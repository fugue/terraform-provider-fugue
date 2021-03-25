package fugue

import "time"

const (
	// EnvironmentRetryTimeout defines the maximum time to retry on
	// errors when changing an environment
	EnvironmentRetryTimeout = 30 * time.Second
)

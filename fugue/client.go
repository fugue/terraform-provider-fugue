package fugue

import (
	"fmt"
	"os"

	"github.com/fugue/fugue-client/client"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

const (
	// DefaultHost is the default hostname of the Fugue API
	DefaultHost = "api.riskmanager.fugue.co"

	// DefaultBase is the base path of the Fugue API
	DefaultBase = "v0"
)

func mustGetEnv(name string) string {
	value := os.Getenv(name)
	if value == "" {
		fmt.Fprintf(os.Stderr, "Missing environment variable: %s\n", name)
		os.Exit(1)
	}
	return value
}

func getEnvWithDefault(name, defaultValue string) string {
	value := os.Getenv(name)
	if value == "" {
		return defaultValue
	}
	return value
}

func getClient() (*client.Fugue, runtime.ClientAuthInfoWriter) {

	clientID := mustGetEnv("FUGUE_API_ID")
	clientSecret := mustGetEnv("FUGUE_API_SECRET")

	host := getEnvWithDefault("FUGUE_API_HOST", DefaultHost)
	base := getEnvWithDefault("FUGUE_API_BASE", DefaultBase)

	transport := httptransport.New(host, base, []string{"https"})
	apiClient := client.New(transport, strfmt.Default)

	auth := httptransport.BasicAuth(clientID, clientSecret)

	return apiClient, auth
}

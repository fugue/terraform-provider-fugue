package fugue

import (
	"errors"
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

// Client supports interactions with the Fugue API
type Client struct {
	*client.Fugue
	Auth runtime.ClientAuthInfoWriter
}

func getFugueClient(clientID, clientSecret string) (*Client, error) {

	if clientID == "" {
		return nil, errors.New("FUGUE_API_ID is not set")
	}

	if clientSecret == "" {
		return nil, errors.New("FUGUE_API_SECRET is not set")
	}

	host := getEnvWithDefault("FUGUE_API_HOST", DefaultHost)
	base := getEnvWithDefault("FUGUE_API_BASE", DefaultBase)

	transport := httptransport.New(host, base, []string{"https"})
	apiClient := client.New(transport, strfmt.Default)
	auth := httptransport.BasicAuth(clientID, clientSecret)

	return &Client{Fugue: apiClient, Auth: auth}, nil
}

func getEnvWithDefault(name, defaultValue string) string {
	value := os.Getenv(name)
	if value == "" {
		return defaultValue
	}
	return value
}

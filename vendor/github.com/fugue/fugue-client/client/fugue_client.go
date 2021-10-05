// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/fugue/fugue-client/client/audit_log"
	"github.com/fugue/fugue-client/client/custom_rules"
	"github.com/fugue/fugue-client/client/environments"
	"github.com/fugue/fugue-client/client/events"
	"github.com/fugue/fugue-client/client/families"
	"github.com/fugue/fugue-client/client/groups"
	"github.com/fugue/fugue-client/client/invites"
	"github.com/fugue/fugue-client/client/metadata"
	"github.com/fugue/fugue-client/client/notifications"
	"github.com/fugue/fugue-client/client/rule_waivers"
	"github.com/fugue/fugue-client/client/scans"
	"github.com/fugue/fugue-client/client/users"
)

// Default fugue HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "api.riskmanager.fugue.co"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/v0"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new fugue HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Fugue {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new fugue HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Fugue {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new fugue client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Fugue {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(Fugue)
	cli.Transport = transport
	cli.AuditLog = audit_log.New(transport, formats)
	cli.CustomRules = custom_rules.New(transport, formats)
	cli.Environments = environments.New(transport, formats)
	cli.Events = events.New(transport, formats)
	cli.Families = families.New(transport, formats)
	cli.Groups = groups.New(transport, formats)
	cli.Invites = invites.New(transport, formats)
	cli.Metadata = metadata.New(transport, formats)
	cli.Notifications = notifications.New(transport, formats)
	cli.RuleWaivers = rule_waivers.New(transport, formats)
	cli.Scans = scans.New(transport, formats)
	cli.Users = users.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// Fugue is a client for fugue
type Fugue struct {
	AuditLog audit_log.ClientService

	CustomRules custom_rules.ClientService

	Environments environments.ClientService

	Events events.ClientService

	Families families.ClientService

	Groups groups.ClientService

	Invites invites.ClientService

	Metadata metadata.ClientService

	Notifications notifications.ClientService

	RuleWaivers rule_waivers.ClientService

	Scans scans.ClientService

	Users users.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Fugue) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.AuditLog.SetTransport(transport)
	c.CustomRules.SetTransport(transport)
	c.Environments.SetTransport(transport)
	c.Events.SetTransport(transport)
	c.Families.SetTransport(transport)
	c.Groups.SetTransport(transport)
	c.Invites.SetTransport(transport)
	c.Metadata.SetTransport(transport)
	c.Notifications.SetTransport(transport)
	c.RuleWaivers.SetTransport(transport)
	c.Scans.SetTransport(transport)
	c.Users.SetTransport(transport)
}

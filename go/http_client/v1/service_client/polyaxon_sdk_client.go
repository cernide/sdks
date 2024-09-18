// Code generated by go-swagger; DO NOT EDIT.

package service_client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/cernide/sdks/go/http_client/v1/service_client/agents_v1"
	"github.com/cernide/sdks/go/http_client/v1/service_client/artifacts_stores_v1"
	"github.com/cernide/sdks/go/http_client/v1/service_client/auth_v1"
	"github.com/cernide/sdks/go/http_client/v1/service_client/connections_v1"
	"github.com/cernide/sdks/go/http_client/v1/service_client/dashboards_v1"
	"github.com/cernide/sdks/go/http_client/v1/service_client/organizations_v1"
	"github.com/cernide/sdks/go/http_client/v1/service_client/policies_v1"
	"github.com/cernide/sdks/go/http_client/v1/service_client/presets_v1"
	"github.com/cernide/sdks/go/http_client/v1/service_client/project_dashboards_v1"
	"github.com/cernide/sdks/go/http_client/v1/service_client/project_searches_v1"
	"github.com/cernide/sdks/go/http_client/v1/service_client/projects_v1"
	"github.com/cernide/sdks/go/http_client/v1/service_client/queues_v1"
	"github.com/cernide/sdks/go/http_client/v1/service_client/runs_v1"
	"github.com/cernide/sdks/go/http_client/v1/service_client/schemas_v1"
	"github.com/cernide/sdks/go/http_client/v1/service_client/searches_v1"
	"github.com/cernide/sdks/go/http_client/v1/service_client/service_accounts_v1"
	"github.com/cernide/sdks/go/http_client/v1/service_client/tags_v1"
	"github.com/cernide/sdks/go/http_client/v1/service_client/teams_v1"
	"github.com/cernide/sdks/go/http_client/v1/service_client/users_v1"
	"github.com/cernide/sdks/go/http_client/v1/service_client/versions_v1"
)

// Default polyaxon sdk HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http", "https"}

// NewHTTPClient creates a new polyaxon sdk HTTP client.
func NewHTTPClient(formats strfmt.Registry) *PolyaxonSdk {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new polyaxon sdk HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *PolyaxonSdk {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new polyaxon sdk client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *PolyaxonSdk {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(PolyaxonSdk)
	cli.Transport = transport
	cli.AgentsV1 = agents_v1.New(transport, formats)
	cli.ArtifactsStoresV1 = artifacts_stores_v1.New(transport, formats)
	cli.AuthV1 = auth_v1.New(transport, formats)
	cli.ConnectionsV1 = connections_v1.New(transport, formats)
	cli.DashboardsV1 = dashboards_v1.New(transport, formats)
	cli.OrganizationsV1 = organizations_v1.New(transport, formats)
	cli.PoliciesV1 = policies_v1.New(transport, formats)
	cli.PresetsV1 = presets_v1.New(transport, formats)
	cli.ProjectDashboardsV1 = project_dashboards_v1.New(transport, formats)
	cli.ProjectSearchesV1 = project_searches_v1.New(transport, formats)
	cli.ProjectsV1 = projects_v1.New(transport, formats)
	cli.QueuesV1 = queues_v1.New(transport, formats)
	cli.RunsV1 = runs_v1.New(transport, formats)
	cli.SchemasV1 = schemas_v1.New(transport, formats)
	cli.SearchesV1 = searches_v1.New(transport, formats)
	cli.ServiceAccountsV1 = service_accounts_v1.New(transport, formats)
	cli.TagsV1 = tags_v1.New(transport, formats)
	cli.TeamsV1 = teams_v1.New(transport, formats)
	cli.UsersV1 = users_v1.New(transport, formats)
	cli.VersionsV1 = versions_v1.New(transport, formats)
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

// PolyaxonSdk is a client for polyaxon sdk
type PolyaxonSdk struct {
	AgentsV1 agents_v1.ClientService

	ArtifactsStoresV1 artifacts_stores_v1.ClientService

	AuthV1 auth_v1.ClientService

	ConnectionsV1 connections_v1.ClientService

	DashboardsV1 dashboards_v1.ClientService

	OrganizationsV1 organizations_v1.ClientService

	PoliciesV1 policies_v1.ClientService

	PresetsV1 presets_v1.ClientService

	ProjectDashboardsV1 project_dashboards_v1.ClientService

	ProjectSearchesV1 project_searches_v1.ClientService

	ProjectsV1 projects_v1.ClientService

	QueuesV1 queues_v1.ClientService

	RunsV1 runs_v1.ClientService

	SchemasV1 schemas_v1.ClientService

	SearchesV1 searches_v1.ClientService

	ServiceAccountsV1 service_accounts_v1.ClientService

	TagsV1 tags_v1.ClientService

	TeamsV1 teams_v1.ClientService

	UsersV1 users_v1.ClientService

	VersionsV1 versions_v1.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *PolyaxonSdk) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.AgentsV1.SetTransport(transport)
	c.ArtifactsStoresV1.SetTransport(transport)
	c.AuthV1.SetTransport(transport)
	c.ConnectionsV1.SetTransport(transport)
	c.DashboardsV1.SetTransport(transport)
	c.OrganizationsV1.SetTransport(transport)
	c.PoliciesV1.SetTransport(transport)
	c.PresetsV1.SetTransport(transport)
	c.ProjectDashboardsV1.SetTransport(transport)
	c.ProjectSearchesV1.SetTransport(transport)
	c.ProjectsV1.SetTransport(transport)
	c.QueuesV1.SetTransport(transport)
	c.RunsV1.SetTransport(transport)
	c.SchemasV1.SetTransport(transport)
	c.SearchesV1.SetTransport(transport)
	c.ServiceAccountsV1.SetTransport(transport)
	c.TagsV1.SetTransport(transport)
	c.TeamsV1.SetTransport(transport)
	c.UsersV1.SetTransport(transport)
	c.VersionsV1.SetTransport(transport)
}

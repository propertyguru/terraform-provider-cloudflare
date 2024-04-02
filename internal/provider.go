// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package internal

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-terraform/internal/dns_records"
	"github.com/cloudflare/cloudflare-terraform/internal/email_routing_rules"
	"github.com/cloudflare/cloudflare-terraform/internal/filters"
	"github.com/cloudflare/cloudflare-terraform/internal/firewall_rules"
	"github.com/cloudflare/cloudflare-terraform/internal/hyperdrive_configs"
	"github.com/cloudflare/cloudflare-terraform/internal/logpush_jobs"
	"github.com/cloudflare/cloudflare-terraform/internal/magic_transit_gre_tunnels"
	"github.com/cloudflare/cloudflare-terraform/internal/magic_transit_ipsec_tunnels"
	"github.com/cloudflare/cloudflare-terraform/internal/pagerules"
	"github.com/cloudflare/cloudflare-terraform/internal/rate_limits"
	"github.com/cloudflare/cloudflare-terraform/internal/rules_lists"
	"github.com/cloudflare/cloudflare-terraform/internal/waiting_rooms"
	"github.com/cloudflare/cloudflare-terraform/internal/waiting_rooms_events"
	"github.com/cloudflare/cloudflare-terraform/internal/workers_routes"
	"github.com/cloudflare/cloudflare-terraform/internal/zero_trust_access_applications"
	"github.com/cloudflare/cloudflare-terraform/internal/zero_trust_access_certificates"
	"github.com/cloudflare/cloudflare-terraform/internal/zero_trust_access_custom_pages"
	"github.com/cloudflare/cloudflare-terraform/internal/zero_trust_access_tags"
	"github.com/cloudflare/cloudflare-terraform/internal/zero_trust_devices_dex_tests"
	"github.com/cloudflare/cloudflare-terraform/internal/zero_trust_identity_providers"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ provider.Provider = &CloudflareProvider{}

// CloudflareProvider defines the provider implementation.
type CloudflareProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// CloudflareProviderModel describes the provider data model.
type CloudflareProviderModel struct {
	APIKey         types.String `tfsdk:"api_key" json:"api_key"`
	APIEmail       types.String `tfsdk:"api_email" json:"api_email"`
	APIToken       types.String `tfsdk:"api_token" json:"api_token"`
	UserServiceKey types.String `tfsdk:"user_service_key" json:"user_service_key"`
}

func (p *CloudflareProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "cloudflare"
	resp.Version = p.version
}

func (p CloudflareProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"api_key": schema.StringAttribute{
				Optional: true,
			},
			"api_email": schema.StringAttribute{
				Optional: true,
			},
			"api_token": schema.StringAttribute{
				Optional: true,
			},
			"user_service_key": schema.StringAttribute{
				Optional: true,
			},
		},
	}
}

func (p *CloudflareProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {

	// TODO(terraform): apiKey := os.Getenv("API_KEY")

	var data CloudflareProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	opts := []option.RequestOption{}

	if !data.APIKey.IsNull() {
		opts = append(opts, option.WithAPIKey(data.APIKey.ValueString()))
	}
	if !data.APIEmail.IsNull() {
		opts = append(opts, option.WithAPIEmail(data.APIEmail.ValueString()))
	}
	if !data.APIToken.IsNull() {
		opts = append(opts, option.WithAPIToken(data.APIToken.ValueString()))
	}
	if !data.UserServiceKey.IsNull() {
		opts = append(opts, option.WithUserServiceKey(data.UserServiceKey.ValueString()))
	}

	client := cloudflare.NewClient(
		opts...,
	)

	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *CloudflareProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		dns_records.NewResource,
		email_routing_rules.NewResource,
		filters.NewResource,
		firewall_rules.NewResource,
		logpush_jobs.NewResource,
		pagerules.NewResource,
		rate_limits.NewResource,
		waiting_rooms.NewResource,
		waiting_rooms_events.NewResource,
		workers_routes.NewResource,
		magic_transit_gre_tunnels.NewResource,
		magic_transit_ipsec_tunnels.NewResource,
		rules_lists.NewResource,
		zero_trust_devices_dex_tests.NewResource,
		zero_trust_identity_providers.NewResource,
		zero_trust_access_applications.NewResource,
		zero_trust_access_certificates.NewResource,
		zero_trust_access_custom_pages.NewResource,
		zero_trust_access_tags.NewResource,
		hyperdrive_configs.NewResource,
	}
}

func (p *CloudflareProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func NewProvider(version string) func() provider.Provider {
	return func() provider.Provider {
		return &CloudflareProvider{
			version: version,
		}
	}
}
// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package internal

import (
	"context"
	"os"

	"github.com/Metronome-Industries/metronome-go"
	"github.com/Metronome-Industries/metronome-go/option"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ provider.ProviderWithConfigValidators = (*MetronomeProvider)(nil)

// MetronomeProvider defines the provider implementation.
type MetronomeProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// MetronomeProviderModel describes the provider data model.
type MetronomeProviderModel struct {
	BaseURL       types.String `tfsdk:"base_url" json:"base_url,optional"`
	BearerToken   types.String `tfsdk:"bearer_token" json:"bearer_token,optional"`
	WebhookSecret types.String `tfsdk:"webhook_secret" json:"webhook_secret,optional"`
}

func (p *MetronomeProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "metronome"
	resp.Version = p.version
}

func ProviderSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"base_url": schema.StringAttribute{
				Description: "Set the base url that the provider connects to.",
				Optional:    true,
			},
			"bearer_token": schema.StringAttribute{
				Optional: true,
			},
			"webhook_secret": schema.StringAttribute{
				Optional: true,
			},
		},
	}
}

func (p *MetronomeProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = ProviderSchema(ctx)
}

func (p *MetronomeProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {

	var data MetronomeProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	opts := []option.RequestOption{}

	if !data.BaseURL.IsNull() && !data.BaseURL.IsUnknown() {
		opts = append(opts, option.WithBaseURL(data.BaseURL.ValueString()))
	} else if o, ok := os.LookupEnv("METRONOME_BASE_URL"); ok {
		opts = append(opts, option.WithBaseURL(o))
	}

	if !data.BearerToken.IsNull() && !data.BearerToken.IsUnknown() {
		opts = append(opts, option.WithBearerToken(data.BearerToken.ValueString()))
	} else if o, ok := os.LookupEnv("METRONOME_BEARER_TOKEN"); ok {
		opts = append(opts, option.WithBearerToken(o))
	} else {
		resp.Diagnostics.AddAttributeError(
			path.Root("bearer_token"),
			"Missing bearer_token value",
			"The bearer_token field is required. Set it in provider configuration or via the \"METRONOME_BEARER_TOKEN\" environment variable.",
		)
		return
	}

	if !data.WebhookSecret.IsNull() && !data.WebhookSecret.IsUnknown() {
		opts = append(opts, option.WithWebhookSecret(data.WebhookSecret.ValueString()))
	} else if o, ok := os.LookupEnv("METRONOME_WEBHOOK_SECRET"); ok {
		opts = append(opts, option.WithWebhookSecret(o))
	}

	client := metronome.NewClient(
		opts...,
	)

	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *MetronomeProvider) ConfigValidators(_ context.Context) []provider.ConfigValidator {
	return []provider.ConfigValidator{}
}

func (p *MetronomeProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}

func (p *MetronomeProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func NewProvider(version string) func() provider.Provider {
	return func() provider.Provider {
		return &MetronomeProvider{
			version: version,
		}
	}
}

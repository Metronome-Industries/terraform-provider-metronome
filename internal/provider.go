// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package internal

import (
	"context"

	"github.com/Metronome-Industries/metronome-go"
	"github.com/Metronome-Industries/metronome-go/option"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/metronome-terraform/internal/resources/billable_metric"
)

var _ provider.Provider = &MetronomeProvider{}

// MetronomeProvider defines the provider implementation.
type MetronomeProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// MetronomeProviderModel describes the provider data model.
type MetronomeProviderModel struct {
	BaseURL       types.String `tfsdk:"base_url" json:"base_url"`
	BearerToken   types.String `tfsdk:"bearer_token" json:"bearer_token"`
	WebhookSecret types.String `tfsdk:"webhook_secret" json:"webhook_secret"`
}

func (p *MetronomeProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "metronome"
	resp.Version = p.version
}

func (p MetronomeProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"base_url": schema.StringAttribute{
				Description: "Set the base url that the provider connects to. This can be used for testing in other environments.",
				Optional:    true,
			},
			"bearer_token": schema.StringAttribute{
				Required: true,
			},
			"webhook_secret": schema.StringAttribute{
				Optional: true,
			},
		},
	}
}

func (p *MetronomeProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {

	// TODO(terraform): apiKey := os.Getenv("API_KEY")

	var data MetronomeProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	opts := []option.RequestOption{}

	if !data.BaseURL.IsNull() {
		opts = append(opts, option.WithBaseURL(data.BaseURL.ValueString()))
	}
	if !data.BearerToken.IsNull() {
		opts = append(opts, option.WithBearerToken(data.BearerToken.ValueString()))
	}
	if !data.WebhookSecret.IsNull() {
		opts = append(opts, option.WithWebhookSecret(data.WebhookSecret.ValueString()))
	}

	client := metronome.NewClient(
		opts...,
	)

	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *MetronomeProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		billable_metric.NewResource,
	}
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

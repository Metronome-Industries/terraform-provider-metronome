// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package internal

import (
	"context"
	"os"

	"github.com/Metronome-Industries/metronome-go"
	"github.com/Metronome-Industries/metronome-go/option"
	"github.com/Metronome-Industries/terraform-provider-metronome/internal/services/v1_alert"
	"github.com/Metronome-Industries/terraform-provider-metronome/internal/services/v1_audit_log"
	"github.com/Metronome-Industries/terraform-provider-metronome/internal/services/v1_billable_metric"
	"github.com/Metronome-Industries/terraform-provider-metronome/internal/services/v1_contract"
	"github.com/Metronome-Industries/terraform-provider-metronome/internal/services/v1_contract_named_schedule"
	"github.com/Metronome-Industries/terraform-provider-metronome/internal/services/v1_contract_product"
	"github.com/Metronome-Industries/terraform-provider-metronome/internal/services/v1_contract_rate_card"
	"github.com/Metronome-Industries/terraform-provider-metronome/internal/services/v1_contract_rate_card_named_schedule"
	"github.com/Metronome-Industries/terraform-provider-metronome/internal/services/v1_contract_rate_card_product_order"
	"github.com/Metronome-Industries/terraform-provider-metronome/internal/services/v1_contract_rate_card_rate"
	"github.com/Metronome-Industries/terraform-provider-metronome/internal/services/v1_credit_grant"
	"github.com/Metronome-Industries/terraform-provider-metronome/internal/services/v1_customer"
	"github.com/Metronome-Industries/terraform-provider-metronome/internal/services/v1_customer_alert"
	"github.com/Metronome-Industries/terraform-provider-metronome/internal/services/v1_customer_billing_config"
	"github.com/Metronome-Industries/terraform-provider-metronome/internal/services/v1_customer_commit"
	"github.com/Metronome-Industries/terraform-provider-metronome/internal/services/v1_customer_credit"
	"github.com/Metronome-Industries/terraform-provider-metronome/internal/services/v1_customer_invoice"
	"github.com/Metronome-Industries/terraform-provider-metronome/internal/services/v1_customer_named_schedule"
	"github.com/Metronome-Industries/terraform-provider-metronome/internal/services/v1_customer_plan"
	"github.com/Metronome-Industries/terraform-provider-metronome/internal/services/v1_plan"
	"github.com/Metronome-Industries/terraform-provider-metronome/internal/services/v1_pricing_unit"
	"github.com/Metronome-Industries/terraform-provider-metronome/internal/services/v2_contract"
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
				Description: "Set the base url that the provider connects to. This can be used for testing in other environments.",
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
	return []func() resource.Resource{
		v2_contract.NewResource,
		v1_alert.NewResource,
		v1_credit_grant.NewResource,
		v1_customer.NewResource,
		v1_customer_billing_config.NewResource,
		v1_customer_commit.NewResource,
		v1_customer_credit.NewResource,
		v1_customer_named_schedule.NewResource,
		v1_billable_metric.NewResource,
		v1_contract.NewResource,
		v1_contract_product.NewResource,
		v1_contract_rate_card.NewResource,
		v1_contract_rate_card_product_order.NewResource,
		v1_contract_rate_card_named_schedule.NewResource,
		v1_contract_named_schedule.NewResource,
	}
}

func (p *MetronomeProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		v2_contract.NewV2ContractDataSource,
		v1_plan.NewV1PlansDataSource,
		v1_credit_grant.NewV1CreditGrantsDataSource,
		v1_pricing_unit.NewV1PricingUnitsDataSource,
		v1_customer.NewV1CustomerDataSource,
		v1_customer.NewV1CustomersDataSource,
		v1_customer_alert.NewV1CustomerAlertDataSource,
		v1_customer_plan.NewV1CustomerPlansDataSource,
		v1_customer_invoice.NewV1CustomerInvoiceDataSource,
		v1_customer_invoice.NewV1CustomerInvoicesDataSource,
		v1_customer_billing_config.NewV1CustomerBillingConfigDataSource,
		v1_customer_named_schedule.NewV1CustomerNamedScheduleDataSource,
		v1_audit_log.NewV1AuditLogsDataSource,
		v1_billable_metric.NewV1BillableMetricDataSource,
		v1_billable_metric.NewV1BillableMetricsDataSource,
		v1_contract.NewV1ContractDataSource,
		v1_contract_product.NewV1ContractProductDataSource,
		v1_contract_product.NewV1ContractProductsDataSource,
		v1_contract_rate_card.NewV1ContractRateCardDataSource,
		v1_contract_rate_card.NewV1ContractRateCardsDataSource,
		v1_contract_rate_card_rate.NewV1ContractRateCardRatesDataSource,
		v1_contract_rate_card_named_schedule.NewV1ContractRateCardNamedScheduleDataSource,
		v1_contract_named_schedule.NewV1ContractNamedScheduleDataSource,
	}
}

func NewProvider(version string) func() provider.Provider {
	return func() provider.Provider {
		return &MetronomeProvider{
			version: version,
		}
	}
}

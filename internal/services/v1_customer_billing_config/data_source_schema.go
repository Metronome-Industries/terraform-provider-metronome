// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_customer_billing_config

import (
	"context"

	"github.com/Metronome-Industries/terraform-provider-metronome/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*V1CustomerBillingConfigDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"billing_provider_type": schema.StringAttribute{
				Description: `Available values: "aws_marketplace", "stripe", "netsuite", "custom", "azure_marketplace", "quickbooks_online", "workday", "gcp_marketplace".`,
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"aws_marketplace",
						"stripe",
						"netsuite",
						"custom",
						"azure_marketplace",
						"quickbooks_online",
						"workday",
						"gcp_marketplace",
					),
				},
			},
			"customer_id": schema.StringAttribute{
				Required: true,
			},
			"data": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[V1CustomerBillingConfigDataDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"aws_expiration_date": schema.StringAttribute{
						Description: "Contract expiration date for the customer. The expected format is RFC 3339 and can be retrieved from [AWS's GetEntitlements API](https://docs.aws.amazon.com/marketplaceentitlement/latest/APIReference/API_GetEntitlements.html).",
						Computed:    true,
						CustomType:  timetypes.RFC3339Type{},
					},
					"aws_product_code": schema.StringAttribute{
						Computed: true,
					},
					"aws_region": schema.StringAttribute{
						Description: `Available values: "af-south-1", "ap-east-1", "ap-northeast-1", "ap-northeast-2", "ap-northeast-3", "ap-south-1", "ap-southeast-1", "ap-southeast-2", "ca-central-1", "cn-north-1", "cn-northwest-1", "eu-central-1", "eu-north-1", "eu-south-1", "eu-west-1", "eu-west-2", "eu-west-3", "me-south-1", "sa-east-1", "us-east-1", "us-east-2", "us-gov-east-1", "us-gov-west-1", "us-west-1", "us-west-2".`,
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"af-south-1",
								"ap-east-1",
								"ap-northeast-1",
								"ap-northeast-2",
								"ap-northeast-3",
								"ap-south-1",
								"ap-southeast-1",
								"ap-southeast-2",
								"ca-central-1",
								"cn-north-1",
								"cn-northwest-1",
								"eu-central-1",
								"eu-north-1",
								"eu-south-1",
								"eu-west-1",
								"eu-west-2",
								"eu-west-3",
								"me-south-1",
								"sa-east-1",
								"us-east-1",
								"us-east-2",
								"us-gov-east-1",
								"us-gov-west-1",
								"us-west-1",
								"us-west-2",
							),
						},
					},
					"azure_expiration_date": schema.StringAttribute{
						Description: "Subscription term start/end date for the customer. The expected format is RFC 3339 and can be retrieved from [Azure's Get Subscription API](https://learn.microsoft.com/en-us/partner-center/marketplace/partner-center-portal/pc-saas-fulfillment-subscription-api#get-subscription).",
						Computed:    true,
						CustomType:  timetypes.RFC3339Type{},
					},
					"azure_plan_id": schema.StringAttribute{
						Computed: true,
					},
					"azure_start_date": schema.StringAttribute{
						Description: "Subscription term start/end date for the customer. The expected format is RFC 3339 and can be retrieved from [Azure's Get Subscription API](https://learn.microsoft.com/en-us/partner-center/marketplace/partner-center-portal/pc-saas-fulfillment-subscription-api#get-subscription).",
						Computed:    true,
						CustomType:  timetypes.RFC3339Type{},
					},
					"azure_subscription_status": schema.StringAttribute{
						Description: `Available values: "Subscribed", "Unsubscribed", "Suspended", "PendingFulfillmentStart".`,
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"Subscribed",
								"Unsubscribed",
								"Suspended",
								"PendingFulfillmentStart",
							),
						},
					},
					"billing_provider_customer_id": schema.StringAttribute{
						Computed: true,
					},
					"stripe_collection_method": schema.StringAttribute{
						Description: `Available values: "charge_automatically", "send_invoice".`,
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("charge_automatically", "send_invoice"),
						},
					},
				},
			},
		},
	}
}

func (d *V1CustomerBillingConfigDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *V1CustomerBillingConfigDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

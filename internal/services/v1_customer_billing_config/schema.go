// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_customer_billing_config

import (
	"context"

	"github.com/Metronome-Industries/terraform-provider-metronome/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*V1CustomerBillingConfigResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
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
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"customer_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"billing_provider_customer_id": schema.StringAttribute{
				Description:   "The customer ID in the billing provider's system. For Azure, this is the subscription ID.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"aws_product_code": schema.StringAttribute{
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"aws_region": schema.StringAttribute{
				Description: `Available values: "af-south-1", "ap-east-1", "ap-northeast-1", "ap-northeast-2", "ap-northeast-3", "ap-south-1", "ap-southeast-1", "ap-southeast-2", "ca-central-1", "cn-north-1", "cn-northwest-1", "eu-central-1", "eu-north-1", "eu-south-1", "eu-west-1", "eu-west-2", "eu-west-3", "me-south-1", "sa-east-1", "us-east-1", "us-east-2", "us-gov-east-1", "us-gov-west-1", "us-west-1", "us-west-2".`,
				Optional:    true,
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
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"stripe_collection_method": schema.StringAttribute{
				Description: `Available values: "charge_automatically", "send_invoice".`,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("charge_automatically", "send_invoice"),
				},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"data": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[V1CustomerBillingConfigDataModel](ctx),
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

func (r *V1CustomerBillingConfigResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *V1CustomerBillingConfigResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}

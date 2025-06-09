// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_customer

import (
	"context"

	"github.com/Metronome-Industries/terraform-provider-metronome/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*V1CustomerResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"customer_id": schema.StringAttribute{
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Description:   "This will be truncated to 160 characters if the provided name is longer.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"external_id": schema.StringAttribute{
				Description:   "(deprecated, use ingest_aliases instead) an alias that can be used to refer to this customer in usage events",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"custom_fields": schema.MapAttribute{
				Optional:      true,
				ElementType:   types.StringType,
				PlanModifiers: []planmodifier.Map{mapplanmodifier.RequiresReplace()},
			},
			"ingest_aliases": schema.ListAttribute{
				Description:   "Aliases that can be used to refer to this customer in usage events",
				Optional:      true,
				ElementType:   types.StringType,
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
			},
			"billing_config": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"billing_provider_customer_id": schema.StringAttribute{
						Required: true,
					},
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
					"aws_is_subscription_product": schema.BoolAttribute{
						Description: "True if the aws_product_code is a SAAS subscription product, false otherwise.",
						Optional:    true,
					},
					"aws_product_code": schema.StringAttribute{
						Optional: true,
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
					},
					"stripe_collection_method": schema.StringAttribute{
						Description: `Available values: "charge_automatically", "send_invoice".`,
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("charge_automatically", "send_invoice"),
						},
					},
				},
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
			},
			"customer_billing_provider_configurations": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"billing_provider": schema.StringAttribute{
							Description: "The billing provider set for this configuration.\nAvailable values: \"aws_marketplace\", \"azure_marketplace\", \"gcp_marketplace\", \"stripe\", \"netsuite\".",
							Required:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"aws_marketplace",
									"azure_marketplace",
									"gcp_marketplace",
									"stripe",
									"netsuite",
								),
							},
						},
						"configuration": schema.MapAttribute{
							Description: "Configuration for the billing provider. The structure of this object is specific to the billing provider and delivery provider combination. Defaults to an empty object, however, for most billing provider + delivery method combinations, it will not be a valid configuration.",
							Optional:    true,
							ElementType: jsontypes.NormalizedType{},
						},
						"delivery_method": schema.StringAttribute{
							Description: "The method to use for delivering invoices to this customer. If not provided, the `delivery_method_id` must be provided.\nAvailable values: \"direct_to_billing_provider\", \"aws_sqs\", \"tackle\", \"aws_sns\".",
							Optional:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"direct_to_billing_provider",
									"aws_sqs",
									"tackle",
									"aws_sns",
								),
							},
						},
						"delivery_method_id": schema.StringAttribute{
							Description: "ID of the delivery method to use for this customer. If not provided, the `delivery_method` must be provided.",
							Optional:    true,
						},
					},
				},
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
			},
			"data": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[V1CustomerDataModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description: "the Metronome ID of the customer",
						Computed:    true,
					},
					"created_at": schema.StringAttribute{
						Description: "RFC 3339 timestamp indicating when the customer was created.",
						Computed:    true,
						CustomType:  timetypes.RFC3339Type{},
					},
					"custom_fields": schema.MapAttribute{
						Computed:    true,
						CustomType:  customfield.NewMapType[types.String](ctx),
						ElementType: types.StringType,
					},
					"customer_config": schema.SingleNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectType[V1CustomerDataCustomerConfigModel](ctx),
						Attributes: map[string]schema.Attribute{
							"salesforce_account_id": schema.StringAttribute{
								Description: "The Salesforce account ID for the customer",
								Computed:    true,
							},
						},
					},
					"external_id": schema.StringAttribute{
						Description: "(deprecated, use ingest_aliases instead) the first ID (Metronome or ingest alias) that can be used in usage events",
						Computed:    true,
					},
					"ingest_aliases": schema.ListAttribute{
						Description: "aliases for this customer that can be used instead of the Metronome customer ID in usage events",
						Computed:    true,
						CustomType:  customfield.NewListType[types.String](ctx),
						ElementType: types.StringType,
					},
					"name": schema.StringAttribute{
						Computed: true,
					},
					"archived_at": schema.StringAttribute{
						Description: "RFC 3339 timestamp indicating when the customer was archived. Null if the customer is active.",
						Computed:    true,
						CustomType:  timetypes.RFC3339Type{},
					},
					"current_billable_status": schema.SingleNestedAttribute{
						Description: "This field's availability is dependent on your client's configuration.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[V1CustomerDataCurrentBillableStatusModel](ctx),
						Attributes: map[string]schema.Attribute{
							"value": schema.StringAttribute{
								Description: `Available values: "billable", "unbillable".`,
								Computed:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("billable", "unbillable"),
								},
							},
							"effective_at": schema.StringAttribute{
								Computed:   true,
								CustomType: timetypes.RFC3339Type{},
							},
						},
					},
				},
			},
		},
	}
}

func (r *V1CustomerResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *V1CustomerResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}

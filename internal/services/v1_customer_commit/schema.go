// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_customer_commit

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/metronome-terraform/internal/customfield"
)

var _ resource.ResourceWithConfigValidators = (*V1CustomerCommitResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"customer_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"priority": schema.Float64Attribute{
				Description:   "If multiple credits or commits are applicable, the one with the lower priority will apply first.",
				Required:      true,
				PlanModifiers: []planmodifier.Float64{float64planmodifier.RequiresReplace()},
			},
			"product_id": schema.StringAttribute{
				Description:   "ID of the fixed product associated with the commit. This is required because products are used to invoice the commit amount.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"type": schema.StringAttribute{
				Description: `Available values: "PREPAID", "POSTPAID".`,
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("PREPAID", "POSTPAID"),
				},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"access_schedule": schema.SingleNestedAttribute{
				Description: `Schedule for distributing the commit to the customer. For "POSTPAID" commits only one schedule item is allowed and amount must match invoice_schedule total.`,
				Required:    true,
				Attributes: map[string]schema.Attribute{
					"schedule_items": schema.ListNestedAttribute{
						Required: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"amount": schema.Float64Attribute{
									Required: true,
								},
								"ending_before": schema.StringAttribute{
									Description: "RFC 3339 timestamp (exclusive)",
									Required:    true,
									CustomType:  timetypes.RFC3339Type{},
								},
								"starting_at": schema.StringAttribute{
									Description: "RFC 3339 timestamp (inclusive)",
									Required:    true,
									CustomType:  timetypes.RFC3339Type{},
								},
							},
						},
					},
					"credit_type_id": schema.StringAttribute{
						Description: "Defaults to USD (cents) if not passed",
						Optional:    true,
					},
				},
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
			},
			"description": schema.StringAttribute{
				Description:   "Used only in UI/API. It is not exposed to end customers.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"invoice_contract_id": schema.StringAttribute{
				Description:   `The contract that this commit will be billed on. This is required for "POSTPAID" commits and for "PREPAID" commits unless there is no invoice schedule above (i.e., the commit is 'free').`,
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Description:   "displayed on invoices",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"netsuite_sales_order_id": schema.StringAttribute{
				Description:   "This field's availability is dependent on your client's configuration.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"rate_type": schema.StringAttribute{
				Description: `Available values: "COMMIT_RATE", "LIST_RATE".`,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("COMMIT_RATE", "LIST_RATE"),
				},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"salesforce_opportunity_id": schema.StringAttribute{
				Description:   "This field's availability is dependent on your client's configuration.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"uniqueness_key": schema.StringAttribute{
				Description:   "Prevents the creation of duplicates. If a request to create a commit or credit is made with a uniqueness key that was previously used to create a commit or credit, a new record will not be created and the request will fail with a 409 error.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"applicable_contract_ids": schema.ListAttribute{
				Description:   "Which contract the commit applies to. If not provided, the commit applies to all contracts.",
				Optional:      true,
				ElementType:   types.StringType,
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
			},
			"applicable_product_ids": schema.ListAttribute{
				Description:   "Which products the commit applies to. If both applicable_product_ids and applicable_product_tags are not provided, the commit applies to all products.",
				Optional:      true,
				ElementType:   types.StringType,
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
			},
			"applicable_product_tags": schema.ListAttribute{
				Description:   "Which tags the commit applies to. If both applicable_product_ids and applicable_product_tags are not provided, the commit applies to all products.",
				Optional:      true,
				ElementType:   types.StringType,
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
			},
			"custom_fields": schema.MapAttribute{
				Optional:      true,
				ElementType:   types.StringType,
				PlanModifiers: []planmodifier.Map{mapplanmodifier.RequiresReplace()},
			},
			"invoice_schedule": schema.SingleNestedAttribute{
				Description: `Required for "POSTPAID" commits: the true up invoice will be generated at this time and only one schedule item is allowed; the total must match accesss_schedule amount. Optional for "PREPAID" commits: if not provided, this will be a "complimentary" commit with no invoice.`,
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"credit_type_id": schema.StringAttribute{
						Description: "Defaults to USD (cents) if not passed.",
						Optional:    true,
					},
					"recurring_schedule": schema.SingleNestedAttribute{
						Description: "Enter the unit price and quantity for the charge or instead only send the amount. If amount is sent, the unit price is assumed to be the amount and quantity is inferred to be 1.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"amount_distribution": schema.StringAttribute{
								Description: `Available values: "DIVIDED", "DIVIDED_ROUNDED", "EACH".`,
								Required:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive(
										"DIVIDED",
										"DIVIDED_ROUNDED",
										"EACH",
									),
								},
							},
							"ending_before": schema.StringAttribute{
								Description: "RFC 3339 timestamp (exclusive).",
								Required:    true,
								CustomType:  timetypes.RFC3339Type{},
							},
							"frequency": schema.StringAttribute{
								Description: `Available values: "MONTHLY", "QUARTERLY", "SEMI_ANNUAL", "ANNUAL".`,
								Required:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive(
										"MONTHLY",
										"QUARTERLY",
										"SEMI_ANNUAL",
										"ANNUAL",
									),
								},
							},
							"starting_at": schema.StringAttribute{
								Description: "RFC 3339 timestamp (inclusive).",
								Required:    true,
								CustomType:  timetypes.RFC3339Type{},
							},
							"amount": schema.Float64Attribute{
								Description: "Amount for the charge. Can be provided instead of unit_price and quantity. If amount is sent, the unit_price is assumed to be the amount and quantity is inferred to be 1.",
								Optional:    true,
							},
							"quantity": schema.Float64Attribute{
								Description: "Quantity for the charge. Will be multiplied by unit_price to determine the amount and must be specified with unit_price. If specified amount cannot be provided.",
								Optional:    true,
							},
							"unit_price": schema.Float64Attribute{
								Description: "Unit price for the charge. Will be multiplied by quantity to determine the amount and must be specified with quantity. If specified amount cannot be provided.",
								Optional:    true,
							},
						},
					},
					"schedule_items": schema.ListNestedAttribute{
						Description: "Either provide amount or provide both unit_price and quantity.",
						Optional:    true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"timestamp": schema.StringAttribute{
									Description: "timestamp of the scheduled event",
									Required:    true,
									CustomType:  timetypes.RFC3339Type{},
								},
								"amount": schema.Float64Attribute{
									Description: "Amount for the charge. Can be provided instead of unit_price and quantity. If amount is sent, the unit_price is assumed to be the amount and quantity is inferred to be 1.",
									Optional:    true,
								},
								"quantity": schema.Float64Attribute{
									Description: "Quantity for the charge. Will be multiplied by unit_price to determine the amount and must be specified with unit_price. If specified amount cannot be provided.",
									Optional:    true,
								},
								"unit_price": schema.Float64Attribute{
									Description: "Unit price for the charge. Will be multiplied by quantity to determine the amount and must be specified with quantity. If specified amount cannot be provided.",
									Optional:    true,
								},
							},
						},
					},
				},
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
			},
			"data": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[V1CustomerCommitDataModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
				},
			},
		},
	}
}

func (r *V1CustomerCommitResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *V1CustomerCommitResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}

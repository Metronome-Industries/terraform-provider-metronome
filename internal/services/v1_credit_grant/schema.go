// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_credit_grant

import (
	"context"

	"github.com/Metronome-Industries/terraform-provider-metronome/internal/customfield"
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
)

var _ resource.ResourceWithConfigValidators = (*V1CreditGrantResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"customer_id": schema.StringAttribute{
				Description:   "the Metronome ID of the customer",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"priority": schema.Float64Attribute{
				Required:      true,
				PlanModifiers: []planmodifier.Float64{float64planmodifier.RequiresReplace()},
			},
			"grant_amount": schema.SingleNestedAttribute{
				Description: "the amount of credits granted",
				Required:    true,
				Attributes: map[string]schema.Attribute{
					"amount": schema.Float64Attribute{
						Required: true,
					},
					"credit_type_id": schema.StringAttribute{
						Description: "the ID of the pricing unit to be used. Defaults to USD (cents) if not passed.",
						Required:    true,
					},
				},
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
			},
			"paid_amount": schema.SingleNestedAttribute{
				Description: "the amount paid for this credit grant",
				Required:    true,
				Attributes: map[string]schema.Attribute{
					"amount": schema.Float64Attribute{
						Required: true,
					},
					"credit_type_id": schema.StringAttribute{
						Description: "the ID of the pricing unit to be used. Defaults to USD (cents) if not passed.",
						Required:    true,
					},
				},
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
			},
			"effective_at": schema.StringAttribute{
				Description:   "The credit grant will only apply to usage or charges dated on or after this timestamp",
				Optional:      true,
				CustomType:    timetypes.RFC3339Type{},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"invoice_date": schema.StringAttribute{
				Description:   "The date to issue an invoice for the paid_amount.",
				Optional:      true,
				CustomType:    timetypes.RFC3339Type{},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"reason": schema.StringAttribute{
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"uniqueness_key": schema.StringAttribute{
				Description:   "Prevents the creation of duplicates. If a request to create a record is made with a previously used uniqueness key, a new record will not be created and the request will fail with a 409 error.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"custom_fields": schema.MapAttribute{
				Description:   "Custom fields to attach to the credit grant.",
				Optional:      true,
				ElementType:   types.StringType,
				PlanModifiers: []planmodifier.Map{mapplanmodifier.RequiresReplace()},
			},
			"product_ids": schema.ListAttribute{
				Description:   "The product(s) which these credits will be applied to. (If unspecified, the credits will be applied to charges for all products.). The array ordering specified here will be used to determine the order in which credits will be applied to invoice line items",
				Optional:      true,
				ElementType:   types.StringType,
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
			},
			"rollover_settings": schema.SingleNestedAttribute{
				Description: "Configure a rollover for this credit grant so if it expires it rolls over a configured amount to a new credit grant. This feature is currently opt-in only. Contact Metronome to be added to the beta.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"expires_at": schema.StringAttribute{
						Description: "The date to expire the rollover credits.",
						Required:    true,
						CustomType:  timetypes.RFC3339Type{},
					},
					"priority": schema.Float64Attribute{
						Description: "The priority to give the rollover credit grant that gets created when a rollover happens.",
						Required:    true,
					},
					"rollover_amount": schema.SingleNestedAttribute{
						Description: "Specify how much to rollover to the rollover credit grant",
						Required:    true,
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Description: "Rollover up to a percentage of the original credit grant amount.\nAvailable values: \"MAX_PERCENTAGE\", \"MAX_AMOUNT\".",
								Required:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("MAX_PERCENTAGE", "MAX_AMOUNT"),
								},
							},
							"value": schema.Float64Attribute{
								Description: "The maximum percentage (0-1) of the original credit grant to rollover.",
								Required:    true,
							},
						},
					},
				},
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
			},
			"expires_at": schema.StringAttribute{
				Description: "The credit grant will only apply to usage or charges dated before this timestamp",
				Required:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"name": schema.StringAttribute{
				Description: "the name of the credit grant as it will appear on invoices",
				Required:    true,
			},
			"credit_grant_type": schema.StringAttribute{
				Optional: true,
			},
			"id": schema.StringAttribute{
				Description: "the ID of the credit grant",
				Optional:    true,
			},
			"data": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[V1CreditGrantDataModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
				},
			},
		},
	}
}

func (r *V1CreditGrantResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *V1CreditGrantResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}

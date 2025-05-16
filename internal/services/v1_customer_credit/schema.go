// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_customer_credit

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

var _ resource.ResourceWithConfigValidators = (*V1CustomerCreditResource)(nil)

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
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"access_schedule": schema.SingleNestedAttribute{
				Description: "Schedule for distributing the credit to the customer.",
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
				Description:   "Which contract the credit applies to. If not provided, the credit applies to all contracts.",
				Optional:      true,
				ElementType:   types.StringType,
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
			},
			"applicable_product_ids": schema.ListAttribute{
				Description:   "Which products the credit applies to. If both applicable_product_ids and applicable_product_tags are not provided, the credit applies to all products.",
				Optional:      true,
				ElementType:   types.StringType,
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
			},
			"applicable_product_tags": schema.ListAttribute{
				Description:   "Which tags the credit applies to. If both applicable_product_ids and applicable_product_tags are not provided, the credit applies to all products.",
				Optional:      true,
				ElementType:   types.StringType,
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
			},
			"custom_fields": schema.MapAttribute{
				Optional:      true,
				ElementType:   types.StringType,
				PlanModifiers: []planmodifier.Map{mapplanmodifier.RequiresReplace()},
			},
			"data": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[V1CustomerCreditDataModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
				},
			},
		},
	}
}

func (r *V1CustomerCreditResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *V1CustomerCreditResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}

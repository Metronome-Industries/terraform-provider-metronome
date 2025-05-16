// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_alert

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/metronome-terraform/internal/customfield"
)

var _ resource.ResourceWithConfigValidators = (*V1AlertResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"alert_type": schema.StringAttribute{
				Description: "Type of the alert\nAvailable values: \"low_credit_balance_reached\", \"spend_threshold_reached\", \"monthly_invoice_total_spend_threshold_reached\", \"low_remaining_days_in_plan_reached\", \"low_remaining_credit_percentage_reached\", \"usage_threshold_reached\", \"low_remaining_days_for_commit_segment_reached\", \"low_remaining_commit_balance_reached\", \"low_remaining_commit_percentage_reached\", \"low_remaining_days_for_contract_credit_segment_reached\", \"low_remaining_contract_credit_balance_reached\", \"low_remaining_contract_credit_percentage_reached\", \"low_remaining_contract_credit_and_commit_balance_reached\", \"invoice_total_reached\".",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"low_credit_balance_reached",
						"spend_threshold_reached",
						"monthly_invoice_total_spend_threshold_reached",
						"low_remaining_days_in_plan_reached",
						"low_remaining_credit_percentage_reached",
						"usage_threshold_reached",
						"low_remaining_days_for_commit_segment_reached",
						"low_remaining_commit_balance_reached",
						"low_remaining_commit_percentage_reached",
						"low_remaining_days_for_contract_credit_segment_reached",
						"low_remaining_contract_credit_balance_reached",
						"low_remaining_contract_credit_percentage_reached",
						"low_remaining_contract_credit_and_commit_balance_reached",
						"invoice_total_reached",
					),
				},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Description:   "Name of the alert",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"threshold": schema.Float64Attribute{
				Description:   "Threshold value of the alert policy.  Depending upon the alert type, this number may represent a financial amount, the days remaining, or a percentage reached.",
				Required:      true,
				PlanModifiers: []planmodifier.Float64{float64planmodifier.RequiresReplace()},
			},
			"billable_metric_id": schema.StringAttribute{
				Description:   "For alerts of type `usage_threshold_reached`, specifies which billable metric to track the usage for.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"credit_type_id": schema.StringAttribute{
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"customer_id": schema.StringAttribute{
				Description:   "If provided, will create this alert for this specific customer. To create an alert for all customers, do not specify a `customer_id`.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"evaluate_on_create": schema.BoolAttribute{
				Description:   "If true, the alert will evaluate immediately on customers that already meet the alert threshold. If false, it will only evaluate on future customers that trigger the alert threshold. Defaults to true.",
				Optional:      true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplace()},
			},
			"plan_id": schema.StringAttribute{
				Description:   "If provided, will create this alert for this specific plan. To create an alert for all customers, do not specify a `plan_id`.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"uniqueness_key": schema.StringAttribute{
				Description:   "Prevents the creation of duplicates. If a request to create a record is made with a previously used uniqueness key, a new record will not be created and the request will fail with a 409 error.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"credit_grant_type_filters": schema.ListAttribute{
				Description:   "An array of strings, representing a way to filter the credit grant this alert applies to, by looking at the credit_grant_type field on the credit grant. This field is only defined for CreditPercentage and CreditBalance alerts",
				Optional:      true,
				ElementType:   types.StringType,
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
			},
			"invoice_types_filter": schema.ListAttribute{
				Description:   "Only supported for invoice_total_reached alerts. A list of invoice types to evaluate.",
				Optional:      true,
				ElementType:   types.StringType,
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
			},
			"custom_field_filters": schema.ListNestedAttribute{
				Description: "A list of custom field filters for alert types that support advanced filtering. Only present for contract invoices.",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"entity": schema.StringAttribute{
							Description: `Available values: "Contract", "Commit", "ContractCredit".`,
							Required:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"Contract",
									"Commit",
									"ContractCredit",
								),
							},
						},
						"key": schema.StringAttribute{
							Required: true,
						},
						"value": schema.StringAttribute{
							Required: true,
						},
					},
				},
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
			},
			"group_key_filter": schema.SingleNestedAttribute{
				Description: "Scopes alert evaluation to a specific presentation group key on individual line items. Only present for spend alerts.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"key": schema.StringAttribute{
						Required: true,
					},
					"value": schema.StringAttribute{
						Required: true,
					},
				},
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
			},
			"data": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[V1AlertDataModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
				},
			},
		},
	}
}

func (r *V1AlertResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *V1AlertResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}

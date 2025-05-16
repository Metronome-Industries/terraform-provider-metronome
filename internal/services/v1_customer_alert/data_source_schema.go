// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_customer_alert

import (
	"context"

	"github.com/Metronome-Industries/terraform-provider-metronome/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*V1CustomerAlertDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"data": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[V1CustomerAlertDataDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"alert": schema.SingleNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectType[V1CustomerAlertDataAlertDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"id": schema.StringAttribute{
								Description: "the Metronome ID of the alert",
								Computed:    true,
							},
							"name": schema.StringAttribute{
								Description: "Name of the alert",
								Computed:    true,
							},
							"status": schema.StringAttribute{
								Description: "Status of the alert\nAvailable values: \"enabled\", \"archived\", \"disabled\".",
								Computed:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive(
										"enabled",
										"archived",
										"disabled",
									),
								},
							},
							"threshold": schema.Float64Attribute{
								Description: "Threshold value of the alert policy",
								Computed:    true,
							},
							"type": schema.StringAttribute{
								Description: "Type of the alert\nAvailable values: \"low_credit_balance_reached\", \"spend_threshold_reached\", \"monthly_invoice_total_spend_threshold_reached\", \"low_remaining_days_in_plan_reached\", \"low_remaining_credit_percentage_reached\", \"usage_threshold_reached\", \"low_remaining_days_for_commit_segment_reached\", \"low_remaining_commit_balance_reached\", \"low_remaining_commit_percentage_reached\", \"low_remaining_days_for_contract_credit_segment_reached\", \"low_remaining_contract_credit_balance_reached\", \"low_remaining_contract_credit_percentage_reached\", \"low_remaining_contract_credit_and_commit_balance_reached\", \"invoice_total_reached\".",
								Computed:    true,
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
							},
							"updated_at": schema.StringAttribute{
								Description: "Timestamp for when the alert was last updated",
								Computed:    true,
								CustomType:  timetypes.RFC3339Type{},
							},
							"credit_grant_type_filters": schema.ListAttribute{
								Description: "An array of strings, representing a way to filter the credit grant this alert applies to, by looking at the credit_grant_type field on the credit grant. This field is only defined for CreditPercentage and CreditBalance alerts",
								Computed:    true,
								CustomType:  customfield.NewListType[types.String](ctx),
								ElementType: types.StringType,
							},
							"credit_type": schema.SingleNestedAttribute{
								Computed:   true,
								CustomType: customfield.NewNestedObjectType[V1CustomerAlertDataAlertCreditTypeDataSourceModel](ctx),
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Computed: true,
									},
									"name": schema.StringAttribute{
										Computed: true,
									},
								},
							},
							"custom_field_filters": schema.ListNestedAttribute{
								Description: "A list of custom field filters for alert types that support advanced filtering",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectListType[V1CustomerAlertDataAlertCustomFieldFiltersDataSourceModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"entity": schema.StringAttribute{
											Description: `Available values: "Contract", "Commit", "ContractCredit".`,
											Computed:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive(
													"Contract",
													"Commit",
													"ContractCredit",
												),
											},
										},
										"key": schema.StringAttribute{
											Computed: true,
										},
										"value": schema.StringAttribute{
											Computed: true,
										},
									},
								},
							},
							"group_key_filter": schema.SingleNestedAttribute{
								Description: "Scopes alert evaluation to a specific presentation group key on individual line items. Only present for spend alerts.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectType[V1CustomerAlertDataAlertGroupKeyFilterDataSourceModel](ctx),
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Computed: true,
									},
									"value": schema.StringAttribute{
										Computed: true,
									},
								},
							},
							"invoice_types_filter": schema.ListAttribute{
								Description: "Only supported for invoice_total_reached alerts. A list of invoice types to evaluate.",
								Computed:    true,
								CustomType:  customfield.NewListType[types.String](ctx),
								ElementType: types.StringType,
							},
							"uniqueness_key": schema.StringAttribute{
								Description: "Prevents the creation of duplicates. If a request to create a record is made with a previously used uniqueness key, a new record will not be created and the request will fail with a 409 error.",
								Computed:    true,
							},
						},
					},
					"customer_status": schema.StringAttribute{
						Description: "The status of the customer alert. If the alert is archived, null will be returned.\nAvailable values: \"ok\", \"in_alarm\", \"evaluating\".",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"ok",
								"in_alarm",
								"evaluating",
							),
						},
					},
					"triggered_by": schema.StringAttribute{
						Description: "If present, indicates the reason the alert was triggered.",
						Computed:    true,
					},
				},
			},
		},
	}
}

func (d *V1CustomerAlertDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *V1CustomerAlertDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_audit_log

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/stainless-sdks/metronome-terraform/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*V1AuditLogsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ending_before": schema.StringAttribute{
				Description: "RFC 3339 timestamp (exclusive). Cannot be used with 'next_page'.",
				Optional:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"resource_id": schema.StringAttribute{
				Description: "Optional parameter that can be used to filter which audit logs are returned. If you specify resource_id, you must also specify resource_type.",
				Optional:    true,
			},
			"resource_type": schema.StringAttribute{
				Description: "Optional parameter that can be used to filter which audit logs are returned. If you specify resource_type, you must also specify resource_id.",
				Optional:    true,
			},
			"sort": schema.StringAttribute{
				Description: "Sort order by timestamp, e.g. date_asc or date_desc. Defaults to date_asc.\nAvailable values: \"date_asc\", \"date_desc\".",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("date_asc", "date_desc"),
				},
			},
			"starting_on": schema.StringAttribute{
				Description: "RFC 3339 timestamp of the earliest audit log to return. Cannot be used with 'next_page'.",
				Optional:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"items": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[V1AuditLogsItemsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed: true,
						},
						"request": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[V1AuditLogsRequestDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Computed: true,
								},
								"ip": schema.StringAttribute{
									Computed: true,
								},
								"user_agent": schema.StringAttribute{
									Computed: true,
								},
							},
						},
						"timestamp": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"action": schema.StringAttribute{
							Computed: true,
						},
						"actor": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[V1AuditLogsActorDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Computed: true,
								},
								"name": schema.StringAttribute{
									Computed: true,
								},
								"email": schema.StringAttribute{
									Computed: true,
								},
							},
						},
						"description": schema.StringAttribute{
							Computed: true,
						},
						"resource_id": schema.StringAttribute{
							Computed: true,
						},
						"resource_type": schema.StringAttribute{
							Computed: true,
						},
						"status": schema.StringAttribute{
							Description: `Available values: "success", "failure", "pending".`,
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"success",
									"failure",
									"pending",
								),
							},
						},
					},
				},
			},
		},
	}
}

func (d *V1AuditLogsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *V1AuditLogsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

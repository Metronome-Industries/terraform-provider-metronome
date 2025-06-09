// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_contract_rate_card_named_schedule

import (
	"context"

	"github.com/Metronome-Industries/terraform-provider-metronome/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*V1ContractRateCardNamedScheduleDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"contract_id": schema.StringAttribute{
				Description: "ID of the contract whose named schedule is to be retrieved",
				Required:    true,
			},
			"customer_id": schema.StringAttribute{
				Description: "ID of the customer whose named schedule is to be retrieved",
				Required:    true,
			},
			"schedule_name": schema.StringAttribute{
				Description: "The identifier for the schedule to be retrieved",
				Required:    true,
			},
			"covering_date": schema.StringAttribute{
				Description: "If provided, at most one schedule segment will be returned (the one that covers this date). If not provided, all segments will be returned.",
				Optional:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"data": schema.ListNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectListType[V1ContractRateCardNamedScheduleDataDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"starting_at": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"value": schema.StringAttribute{
							Computed:   true,
							CustomType: jsontypes.NormalizedType{},
						},
						"ending_before": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
					},
				},
			},
		},
	}
}

func (d *V1ContractRateCardNamedScheduleDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *V1ContractRateCardNamedScheduleDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

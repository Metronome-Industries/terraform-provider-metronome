// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_customer

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/metronome-terraform/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*V1CustomersDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ingest_alias": schema.StringAttribute{
				Description: "Filter the customer list by ingest_alias",
				Optional:    true,
			},
			"only_archived": schema.BoolAttribute{
				Description: "Filter the customer list to only return archived customers. By default, only active customers are returned.",
				Optional:    true,
			},
			"customer_ids": schema.ListAttribute{
				Description: "Filter the customer list by customer_id.  Up to 100 ids can be provided.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"salesforce_account_ids": schema.ListAttribute{
				Description: "Filter the customer list by salesforce_account_id.  Up to 100 ids can be provided.",
				Optional:    true,
				ElementType: types.StringType,
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
				CustomType:  customfield.NewNestedObjectListType[V1CustomersItemsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
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
							CustomType: customfield.NewNestedObjectType[V1CustomersCustomerConfigDataSourceModel](ctx),
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
							CustomType:  customfield.NewNestedObjectType[V1CustomersCurrentBillableStatusDataSourceModel](ctx),
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
		},
	}
}

func (d *V1CustomersDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *V1CustomersDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

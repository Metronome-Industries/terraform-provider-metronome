// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_credit_grant

import (
	"context"
	"fmt"

	"github.com/Metronome-Industries/metronome-go"
	"github.com/Metronome-Industries/terraform-provider-metronome/internal/apijson"
	"github.com/Metronome-Industries/terraform-provider-metronome/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

type V1CreditGrantsDataSource struct {
	client *metronome.Client
}

var _ datasource.DataSourceWithConfigure = (*V1CreditGrantsDataSource)(nil)

func NewV1CreditGrantsDataSource() datasource.DataSource {
	return &V1CreditGrantsDataSource{}
}

func (d *V1CreditGrantsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_v1_credit_grants"
}

func (d *V1CreditGrantsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*metronome.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"unexpected resource configure type",
			fmt.Sprintf("Expected *metronome.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}

func (d *V1CreditGrantsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *V1CreditGrantsDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	params, diags := data.toListParams(ctx)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	env := V1CreditGrantsDataListDataSourceEnvelope{}
	maxItems := int(data.MaxItems.ValueInt64())
	acc := []attr.Value{}
	if maxItems <= 0 {
		maxItems = 1000
	}
	page, err := d.client.V1.CreditGrants.List(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}

	for page != nil && len(page.Data) > 0 {
		bytes := []byte(page.JSON.RawJSON())
		err = apijson.UnmarshalComputed(bytes, &env)
		if err != nil {
			resp.Diagnostics.AddError("failed to unmarshal http request", err.Error())
			return
		}
		acc = append(acc, env.Data.Elements()...)
		if len(acc) >= maxItems {
			break
		}
		page, err = page.GetNextPage()
		if err != nil {
			resp.Diagnostics.AddError("failed to fetch next page", err.Error())
			return
		}
	}

	acc = acc[:min(len(acc), maxItems)]
	result, diags := customfield.NewObjectListFromAttributes[V1CreditGrantsItemsDataSourceModel](ctx, acc)
	resp.Diagnostics.Append(diags...)
	data.Items = result

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_pricing_unit

import (
	"context"

	"github.com/Metronome-Industries/metronome-go"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/metronome-terraform/internal/customfield"
)

type V1PricingUnitsDataListDataSourceEnvelope struct {
	Data customfield.NestedObjectList[V1PricingUnitsItemsDataSourceModel] `json:"data,computed"`
}

type V1PricingUnitsDataSourceModel struct {
	MaxItems types.Int64                                                      `tfsdk:"max_items"`
	Items    customfield.NestedObjectList[V1PricingUnitsItemsDataSourceModel] `tfsdk:"items"`
}

func (m *V1PricingUnitsDataSourceModel) toListParams(_ context.Context) (params metronome.V1PricingUnitListParams, diags diag.Diagnostics) {
	params = metronome.V1PricingUnitListParams{}

	return
}

type V1PricingUnitsItemsDataSourceModel struct {
	ID         types.String `tfsdk:"id" json:"id,computed"`
	IsCurrency types.Bool   `tfsdk:"is_currency" json:"is_currency,computed"`
	Name       types.String `tfsdk:"name" json:"name,computed"`
}

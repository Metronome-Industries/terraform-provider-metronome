// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_plan

import (
	"context"

	"github.com/Metronome-Industries/metronome-go"
	"github.com/Metronome-Industries/terraform-provider-metronome/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type V1PlansDataListDataSourceEnvelope struct {
	Data customfield.NestedObjectList[V1PlansItemsDataSourceModel] `json:"data,computed"`
}

type V1PlansDataSourceModel struct {
	MaxItems types.Int64                                               `tfsdk:"max_items"`
	Items    customfield.NestedObjectList[V1PlansItemsDataSourceModel] `tfsdk:"items"`
}

func (m *V1PlansDataSourceModel) toListParams(_ context.Context) (params metronome.V1PlanListParams, diags diag.Diagnostics) {
	params = metronome.V1PlanListParams{}

	return
}

type V1PlansItemsDataSourceModel struct {
	ID           types.String                  `tfsdk:"id" json:"id,computed"`
	Description  types.String                  `tfsdk:"description" json:"description,computed"`
	Name         types.String                  `tfsdk:"name" json:"name,computed"`
	CustomFields customfield.Map[types.String] `tfsdk:"custom_fields" json:"custom_fields,computed"`
}

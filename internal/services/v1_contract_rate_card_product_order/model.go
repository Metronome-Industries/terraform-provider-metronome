// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_contract_rate_card_product_order

import (
	"github.com/Metronome-Industries/terraform-provider-metronome/internal/apijson"
	"github.com/Metronome-Industries/terraform-provider-metronome/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type V1ContractRateCardProductOrderModel struct {
	RateCardID   types.String                                                      `tfsdk:"rate_card_id" json:"rate_card_id,required"`
	ProductMoves *[]*V1ContractRateCardProductOrderProductMovesModel               `tfsdk:"product_moves" json:"product_moves,required"`
	Data         customfield.NestedObject[V1ContractRateCardProductOrderDataModel] `tfsdk:"data" json:"data,computed"`
}

func (m V1ContractRateCardProductOrderModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m V1ContractRateCardProductOrderModel) MarshalJSONForUpdate(state V1ContractRateCardProductOrderModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type V1ContractRateCardProductOrderProductMovesModel struct {
	Position  types.Float64 `tfsdk:"position" json:"position,required"`
	ProductID types.String  `tfsdk:"product_id" json:"product_id,required"`
}

type V1ContractRateCardProductOrderDataModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

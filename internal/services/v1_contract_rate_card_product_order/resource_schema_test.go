// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_contract_rate_card_product_order_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/metronome-terraform/internal/services/v1_contract_rate_card_product_order"
	"github.com/stainless-sdks/metronome-terraform/internal/test_helpers"
)

func TestV1ContractRateCardProductOrderModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*v1_contract_rate_card_product_order.V1ContractRateCardProductOrderModel)(nil)
	schema := v1_contract_rate_card_product_order.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_contract_rate_card_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/metronome-terraform/internal/services/v1_contract_rate_card"
	"github.com/stainless-sdks/metronome-terraform/internal/test_helpers"
)

func TestV1ContractRateCardModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*v1_contract_rate_card.V1ContractRateCardModel)(nil)
	schema := v1_contract_rate_card.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}

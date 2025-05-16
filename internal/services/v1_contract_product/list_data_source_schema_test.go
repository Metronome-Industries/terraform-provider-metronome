// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_contract_product_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/metronome-terraform/internal/services/v1_contract_product"
	"github.com/stainless-sdks/metronome-terraform/internal/test_helpers"
)

func TestV1ContractProductsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*v1_contract_product.V1ContractProductsDataSourceModel)(nil)
	schema := v1_contract_product.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}

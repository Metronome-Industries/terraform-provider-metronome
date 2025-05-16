// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v2_contract_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/metronome-terraform/internal/services/v2_contract"
	"github.com/stainless-sdks/metronome-terraform/internal/test_helpers"
)

func TestV2ContractDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*v2_contract.V2ContractDataSourceModel)(nil)
	schema := v2_contract.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}

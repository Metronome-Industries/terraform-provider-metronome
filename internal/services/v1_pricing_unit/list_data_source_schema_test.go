// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_pricing_unit_test

import (
	"context"
	"testing"

	"github.com/Metronome-Industries/terraform-provider-metronome/internal/services/v1_pricing_unit"
	"github.com/Metronome-Industries/terraform-provider-metronome/internal/test_helpers"
)

func TestV1PricingUnitsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*v1_pricing_unit.V1PricingUnitsDataSourceModel)(nil)
	schema := v1_pricing_unit.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}

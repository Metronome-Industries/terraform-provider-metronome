// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_customer_plan_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/metronome-terraform/internal/services/v1_customer_plan"
	"github.com/stainless-sdks/metronome-terraform/internal/test_helpers"
)

func TestV1CustomerPlansDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*v1_customer_plan.V1CustomerPlansDataSourceModel)(nil)
	schema := v1_customer_plan.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}

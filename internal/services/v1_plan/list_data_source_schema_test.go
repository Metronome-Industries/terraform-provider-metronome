// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_plan_test

import (
	"context"
	"testing"

	"github.com/Metronome-Industries/terraform-provider-metronome/internal/services/v1_plan"
	"github.com/Metronome-Industries/terraform-provider-metronome/internal/test_helpers"
)

func TestV1PlansDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*v1_plan.V1PlansDataSourceModel)(nil)
	schema := v1_plan.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}

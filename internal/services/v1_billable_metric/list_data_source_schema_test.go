// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_billable_metric_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/metronome-terraform/internal/services/v1_billable_metric"
	"github.com/stainless-sdks/metronome-terraform/internal/test_helpers"
)

func TestV1BillableMetricsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*v1_billable_metric.V1BillableMetricsDataSourceModel)(nil)
	schema := v1_billable_metric.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}

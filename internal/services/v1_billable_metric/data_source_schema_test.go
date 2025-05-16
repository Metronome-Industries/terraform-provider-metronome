// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_billable_metric_test

import (
	"context"
	"testing"

	"github.com/Metronome-Industries/terraform-provider-metronome/internal/services/v1_billable_metric"
	"github.com/Metronome-Industries/terraform-provider-metronome/internal/test_helpers"
)

func TestV1BillableMetricDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*v1_billable_metric.V1BillableMetricDataSourceModel)(nil)
	schema := v1_billable_metric.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}

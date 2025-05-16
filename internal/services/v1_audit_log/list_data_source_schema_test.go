// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v1_audit_log_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/metronome-terraform/internal/services/v1_audit_log"
	"github.com/stainless-sdks/metronome-terraform/internal/test_helpers"
)

func TestV1AuditLogsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*v1_audit_log.V1AuditLogsDataSourceModel)(nil)
	schema := v1_audit_log.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}

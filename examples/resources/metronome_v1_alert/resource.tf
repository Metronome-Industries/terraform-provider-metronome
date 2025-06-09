resource "metronome_v1_alert" "example_v1_alert" {
  alert_type = "low_credit_balance_reached"
  name = "name"
  threshold = 0
  billable_metric_id = "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"
  credit_grant_type_filters = ["string"]
  credit_type_id = "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"
  custom_field_filters = [{
    entity = "Contract"
    key = "key"
    value = "value"
  }]
  customer_id = "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"
  evaluate_on_create = true
  group_key_filter = {
    key = "key"
    value = "value"
  }
  invoice_types_filter = ["PLAN_ARREARS, SCHEDULED, USAGE, CORRECTION, CREDIT_PURCHASE, or SEAT_PURCHASE"]
  plan_id = "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"
  uniqueness_key = "x"
}

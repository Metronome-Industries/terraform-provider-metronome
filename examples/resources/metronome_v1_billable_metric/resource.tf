resource "metronome_v1_billable_metric" "example_v1_billable_metric" {
  name = "name"
  aggregation_key = "aggregation_key"
  aggregation_type = "COUNT"
  custom_fields = {
    foo = "string"
  }
  event_type_filter = {
    in_values = ["string"]
    not_in_values = ["string"]
  }
  group_keys = [["string"]]
  property_filters = [{
    name = "name"
    exists = true
    in_values = ["string"]
    not_in_values = ["string"]
  }]
  sql = "sql"
}

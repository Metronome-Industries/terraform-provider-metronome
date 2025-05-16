data "metronome_v1_audit_logs" "example_v1_audit_logs" {
  ending_before = "2019-12-27T18:11:19.117Z"
  resource_id = "resource_id"
  resource_type = "resource_type"
  sort = "date_asc"
  starting_on = "2019-12-27T18:11:19.117Z"
}

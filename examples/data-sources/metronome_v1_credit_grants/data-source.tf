data "metronome_v1_credit_grants" "example_v1_credit_grants" {
  credit_grant_ids = ["182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"]
  credit_type_ids = ["2714e483-4ff1-48e4-9e25-ac732e8f24f2"]
  customer_ids = ["d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc", "0e5b8609-d901-4992-b394-c3c2e3f37b1c"]
  effective_before = "2022-02-01T00:00:00Z"
  not_expiring_before = "2022-02-01T00:00:00Z"
}

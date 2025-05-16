data "metronome_v1_customers" "example_v1_customers" {
  customer_ids = ["string"]
  ingest_alias = "ingest_alias"
  only_archived = true
  salesforce_account_ids = ["string"]
}

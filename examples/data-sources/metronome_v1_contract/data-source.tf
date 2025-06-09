data "metronome_v1_contract" "example_v1_contract" {
  contract_id = "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"
  customer_id = "13117714-3f05-48e5-a6e9-a66093f13b4d"
  include_balance = true
  include_ledgers = true
}

---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "metronome_v1_customers Data Source - metronome"
subcategory: ""
description: |-
  
---

# metronome_v1_customers (Data Source)



## Example Usage

```terraform
data "metronome_v1_customers" "example_v1_customers" {
  customer_ids = ["string"]
  ingest_alias = "ingest_alias"
  only_archived = true
  salesforce_account_ids = ["string"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `customer_ids` (List of String) Filter the customer list by customer_id.  Up to 100 ids can be provided.
- `ingest_alias` (String) Filter the customer list by ingest_alias
- `max_items` (Number) Max items to fetch, default: 1000
- `only_archived` (Boolean) Filter the customer list to only return archived customers. By default, only active customers are returned.
- `salesforce_account_ids` (List of String) Filter the customer list by salesforce_account_id.  Up to 100 ids can be provided.

### Read-Only

- `items` (Attributes List) The items returned by the data source (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `archived_at` (String) RFC 3339 timestamp indicating when the customer was archived. Null if the customer is active.
- `created_at` (String) RFC 3339 timestamp indicating when the customer was created.
- `current_billable_status` (Attributes) This field's availability is dependent on your client's configuration. (see [below for nested schema](#nestedatt--items--current_billable_status))
- `custom_fields` (Map of String)
- `customer_config` (Attributes) (see [below for nested schema](#nestedatt--items--customer_config))
- `external_id` (String) (deprecated, use ingest_aliases instead) the first ID (Metronome or ingest alias) that can be used in usage events
- `id` (String) the Metronome ID of the customer
- `ingest_aliases` (List of String) aliases for this customer that can be used instead of the Metronome customer ID in usage events
- `name` (String)

<a id="nestedatt--items--current_billable_status"></a>
### Nested Schema for `items.current_billable_status`

Read-Only:

- `effective_at` (String)
- `value` (String) Available values: "billable", "unbillable".


<a id="nestedatt--items--customer_config"></a>
### Nested Schema for `items.customer_config`

Read-Only:

- `salesforce_account_id` (String) The Salesforce account ID for the customer

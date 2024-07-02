---
page_title: "cloudflare_logpush_ownership_challenge Resource - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_logpush_ownership_challenge (Resource)



## Example Usage

```terraform
resource "cloudflare_logpush_ownership_challenge" "example" {
  zone_id          = "0da42c8d2132a9ddaf714f9e7c920711"
  destination_conf = "s3://my-bucket-path?region=us-west-2"
}
```
<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `destination_conf` (String) Uniquely identifies a resource (such as an s3 bucket) where data will be pushed. Additional configuration parameters supported by the destination may be included.

### Optional

- `account_id` (String) The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
- `zone_id` (String) The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.

### Read-Only

- `filename` (String)
- `message` (String)
- `valid` (Boolean)


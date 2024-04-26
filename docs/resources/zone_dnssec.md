---
page_title: "cloudflare_zone_dnssec Resource - Cloudflare"
subcategory: ""
description: |-
  Provides a Cloudflare resource to create and modify zone DNSSEC settings.
---

# cloudflare_zone_dnssec (Resource)

Provides a Cloudflare resource to create and modify zone DNSSEC settings.

## Example Usage

```terraform
resource "cloudflare_zone" "example" {
  zone = "example.com"
}

resource "cloudflare_zone_dnssec" "example" {
  zone_id = cloudflare_zone.example.id
}
```
<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `zone_id` (String) The zone identifier to target for the resource. **Modifying this attribute will force creation of a new resource.**

### Optional

- `modified_on` (String) Zone DNSSEC updated time.

### Read-Only

- `algorithm` (String) Zone DNSSEC algorithm.
- `digest` (String) Zone DNSSEC digest.
- `digest_algorithm` (String) Digest algorithm use for Zone DNSSEC.
- `digest_type` (String) Digest Type for Zone DNSSEC.
- `ds` (String) DS for the Zone DNSSEC.
- `flags` (Number) Zone DNSSEC flags.
- `id` (String) The ID of this resource.
- `key_tag` (Number) Key Tag for the Zone DNSSEC.
- `key_type` (String) Key type used for Zone DNSSEC.
- `public_key` (String) Public Key for the Zone DNSSEC.
- `status` (String) The status of the Zone DNSSEC.

## Import

Import is supported using the following syntax:

```shell
$ terraform import cloudflare_zone_dnssec.example <zone_id>
```
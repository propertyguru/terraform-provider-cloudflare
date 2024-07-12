---
page_title: "cloudflare_access_mutual_tls_certificates Data Source - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_access_mutual_tls_certificates (Data Source)




<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `account_id` (String) The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
- `max_items` (Number) Max items to fetch, default: 1000
- `zone_id` (String) The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.

### Read-Only

- `items` (Attributes List) The items returned by the data source (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Optional:

- `associated_hostnames` (List of String) The hostnames of the applications that will use this certificate.
- `created_at` (String)
- `expires_on` (String)
- `fingerprint` (String) The MD5 fingerprint of the certificate.
- `id` (String) The ID of the application that will use this certificate.
- `name` (String) The name of the certificate.
- `updated_at` (String)


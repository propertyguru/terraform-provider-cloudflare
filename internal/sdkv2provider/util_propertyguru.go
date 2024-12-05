/*
The code in this file is written to suit the needs in PropertyGuru use case.
  - The functions are named with "pg" prefix, indicating PropertyGuru.
  - There are small modifications as well in the other files, mostly to replace the invocation of original function with
    the new function in this file.
  - Any modification to the original code is commented with [PG]......
  - If you want to search what's being modified and why it is being modified, you can search [PG] in the codebase.
    Alternatively, you can also diff the commits.
*/
package sdkv2provider

import (
	"context"
	"sync"

	cloudflare "github.com/cloudflare/cloudflare-go"
)

// pgDNSRecordByZoneIDAndRecordID is a nested map where the first-level key is the Zone ID and the second-level key is the Record ID.
// The value is the DNS Record itself.
var pgDNSRecordByZoneIDAndRecordID = make(map[string]map[string]cloudflare.DNSRecord)

// The mutex is used to lock the operation of fetching DNS records from Cloudflare.
var pgListDNSRecordsMutex sync.Mutex

// The error is used to store the error that occurred when fetching DNS records from Cloudflare.
var pgListDNSRecordsErr error

func pgListDNSRecords(client *cloudflare.API, ctx context.Context, rc *cloudflare.ResourceContainer) {
	pgListDNSRecordsMutex.Lock()
	if _, recordsFound := pgDNSRecordByZoneIDAndRecordID[rc.Identifier]; !recordsFound {
		pgDNSRecordByZoneIDAndRecordID[rc.Identifier] = make(map[string]cloudflare.DNSRecord)
		page, perPage := 1, 5000000
		for {
			param := cloudflare.ListDNSRecordsParams{ResultInfo: cloudflare.ResultInfo{Page: page, PerPage: perPage}}
			records, resultInfo, err := client.ListDNSRecords(ctx, rc, param)
			pgListDNSRecordsErr = err
			for _, record := range records {
				pgDNSRecordByZoneIDAndRecordID[rc.Identifier][record.ID] = record
			}
			if resultInfo.Page == resultInfo.TotalPages {
				break
			}
			page++
		}
	}
	pgListDNSRecordsMutex.Unlock()
}

func pgGetDNSRecord(client *cloudflare.API, ctx context.Context, rc *cloudflare.ResourceContainer, recordID string) (cloudflare.DNSRecord, error) {
	if rc.Identifier == "" {
		return cloudflare.DNSRecord{}, cloudflare.ErrMissingZoneID
	}
	if recordID == "" {
		return cloudflare.DNSRecord{}, cloudflare.ErrMissingDNSRecordID
	}
	pgListDNSRecords(client, ctx, rc)
	if record, recordFound := pgDNSRecordByZoneIDAndRecordID[rc.Identifier][recordID]; recordFound {
		return record, pgListDNSRecordsErr
	}
	return cloudflare.DNSRecord{}, pgListDNSRecordsErr
}

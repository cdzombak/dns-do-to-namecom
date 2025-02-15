package main

import (
	"context"
	"fmt"

	"github.com/digitalocean/godo"
	"github.com/namedotcom/go/namecom"
)

func Migrate(ctx context.Context, doClient *godo.Client, ncClient *namecom.NameCom, domain string, dryRun bool) error {
	sourceRecords, err := getAllDORecords(ctx, doClient, domain)
	if err != nil {
		return err
	}

	fmt.Printf("Found %d records to copy from DigitalOcean\n", len(sourceRecords))

	for _, sourceRecord := range sourceRecords {
		if sourceRecord.Type == "NS" || sourceRecord.Type == "SOA" {
			fmt.Printf("Skipping %s record: %s (%s)\n", sourceRecord.Type, sourceRecord.Name, sourceRecord.Data)
			continue
		}

		if sourceRecord.Type == "SRV" {
			fmt.Printf("Skipping SRV record (not supported by this tool; patches are welcome): %s (%s)\n", sourceRecord.Name, sourceRecord.Data)
			continue
		}

		if !canCreateNamecomRecordOfType(sourceRecord.Type) {
			fmt.Printf("Skipping unsupported %s record (%s)\n", sourceRecord.Type, sourceRecord.Name)
			continue
		}

		fmt.Println("Copying:", sourceRecord.Type, sourceRecord.Name)

		newRec := doToNamecom(sourceRecord)
		newRec.DomainName = domain
		if dryRun {
			fmt.Printf("\twould create record: %+v\n", newRec)
			continue
		}

		resRec, err := ncClient.CreateRecord(&newRec)
		if err != nil {
			return err
		}
		fmt.Printf("\tcreated record %d: %+v\n", resRec.ID, newRec)
	}

	return nil
}

func getAllDORecords(ctx context.Context, client *godo.Client, domain string) ([]godo.DomainRecord, error) {
	var retv []godo.DomainRecord

	page := 0
	more := true
	for more {
		page++
		records, resp, err := client.Domains.Records(ctx, domain, &godo.ListOptions{
			Page: page,
		})
		if err != nil {
			return nil, fmt.Errorf("failed listing DO DNS records (page %d): %w", page, err)
		}
		retv = append(retv, records...)
		if resp.Links == nil || resp.Links.IsLastPage() {
			more = false
		}
	}

	return retv, nil
}

package linode

import (
	"context"
	"strings"

	"github.com/libdns/libdns"
)

type Provider struct {
	Client
	APIToken string `json:"auth_token"`
}

func (p *Provider) unFQDN(fqdn string) string {
	return strings.TrimSuffix(fqdn, ".")
}

func (p *Provider) GetRecords(ctx context.Context, zone string) ([]libdns.Record, error) {
	records, err := p.getDNSEntries(ctx, p.unFQDN(zone))
	if err != nil {
		return nil, err
	}

	return records, nil
}

func (p *Provider) AppendRecords(ctx context.Context, zone string, records []libdns.Record) ([]libdns.Record, error) {
	var appendedRecords []libdns.Record

	for _, record := range records {
		newRecord, err := p.addDNSEntry(ctx, p.unFQDN(zone), record)
		if err != nil {
			return nil, err
		}
		appendedRecords = append(appendedRecords, newRecord)
	}

	return appendedRecords, nil
}

func (p *Provider) DeleteRecords(ctx context.Context, zone string, records []libdns.Record) ([]libdns.Record, error) {
	var deletedRecords []libdns.Record

	for _, record := range records {
		deletedRecord, err := p.removeDNSEntry(ctx, p.unFQDN(zone), record)
		if err != nil {
			return nil, err
		}
		deletedRecords = append(deletedRecords, deletedRecord)
	}

	return deletedRecords, nil
}

func (p *Provider) SetRecords(ctx context.Context, zone string, records []libdns.Record) ([]libdns.Record, error) {
	var setRecords []libdns.Record

	for _, record := range records {
		setRecord, err := p.updateDNSEntry(ctx, p.unFQDN(zone), record)
		if err != nil {
			return setRecords, err
		}
		setRecords = append(setRecords, setRecord)
	}

	return setRecords, nil
}

var (
	_ libdns.RecordGetter   = (*Provider)(nil)
	_ libdns.RecordAppender = (*Provider)(nil)
	_ libdns.RecordSetter   = (*Provider)(nil)
	_ libdns.RecordDeleter  = (*Provider)(nil)
)

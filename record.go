package main

import (
	"fmt"

	"github.com/digitalocean/godo"
	"github.com/namedotcom/go/namecom"
)

func doToNamecom(doRec godo.DomainRecord) namecom.Record {
	retv := namecom.Record{
		Host:     doRec.Name,
		Type:     doRec.Type,
		Answer:   doRec.Data,
		Priority: uint32(doRec.Priority),
		TTL:      uint32(doRec.TTL),
	}

	if doRec.Type == "CAA" {
		retv.Answer = fmt.Sprintf("%d %s \"%s\"", doRec.Flags, doRec.Tag, doRec.Data)
	}

	return retv
}

func canCreateNamecomRecordOfType(t string) bool {
	switch t {
	case "A", "AAAA", "ANAME", "CNAME", "MX", "NS", "SRV", "TXT":
		return true
	default:
		return false
	}
}

package ha

import (
	"testing"

	motan "github.com/weibocom/motan-go/core"
)

func NewFailOver(url *motan.URL) motan.HaStrategy {
	return &FailOverHA{url: url}
}

func TestGetHa(t *testing.T) {
	defaultExtFactory := &motan.DefaultExtensionFactory{}
	defaultExtFactory.Initialize()
	RegistDefaultHa(defaultExtFactory)
	p := make(map[string]string)
	p[motan.Hakey] = FailOver
	url := &motan.URL{Parameters: p}
	ha := defaultExtFactory.GetHa(url)
	if ha.GetName() != "failover" {
		t.Error("GetHa name Error")
	}
	getURL := ha.GetURL()
	if url.GetIdentity() != getURL.GetIdentity() {
		t.Error("GetURL Error")
	}

	p[motan.Hakey] = BackupRequest
	p[motan.FilterKey] = ""
	url = &motan.URL{Parameters: p}
	ha = defaultExtFactory.GetHa(url)
	if ha.GetName() != "backupRequest" {
		t.Error("GetHa name Error")
	}
	getURL = ha.GetURL()
	if url.GetIdentity() != getURL.GetIdentity() {
		t.Error("GetURL Error")
	}
}

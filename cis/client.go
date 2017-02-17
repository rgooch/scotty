package cis

import (
	"fmt"
	"strings"
)

const (
	kCisPath = "packages"
)

var (
	replacer = strings.NewReplacer(".", "_")
)

func (c *Client) write(stats *Stats) error {

	type versionSizeType struct {
		Version string `json:"version"`
		Size    uint64 `json:"size"`
	}

	type packagesType struct {
		PkgMgmtType string                     `json:"pkgMgmtType"`
		Packages    map[string]versionSizeType `json:"packages"`
	}

	jsonToEncode := &packagesType{
		PkgMgmtType: stats.Packages.ManagementType,
		Packages: make(
			map[string]versionSizeType,
			len(stats.Packages.Packages)),
	}

	for _, entry := range stats.Packages.Packages {
		jsonToEncode.Packages[replacer.Replace(entry.Name)] = versionSizeType{
			Version: entry.Version, Size: entry.Size}
	}
	url := fmt.Sprintf("%s/%s/%s", c.endpoint, kCisPath, stats.InstanceId)
	if c.sync != nil {
		return c.sync.Write(url, jsonToEncode)
	} else {
		c.async.Send(url, jsonToEncode)
		return nil
	}
}

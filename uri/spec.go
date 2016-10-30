// Author: lipixun
// Created Time : 二 10/18 22:08:16 2016
//
// File Name: spec.go
// Description:
//	Source code spec structure
package uri

import (
	"strings"
)

const (
	UriTypeKnown     = "unknown"
	UriTypeLocalPath = "localPath"
	UriTypeHttps     = "https"
	UriTypeSSH       = "ssh"
)

func GetUriType(uri string) string {
	uri = strings.ToLower(uri)
	// Check the uri type
	if strings.HasPrefix(uri, "/") || strings.HasPrefix(uri, "file://") {
		return UriTypeLocalPath
	} else if strings.HasPrefix(uri, "https://") {
		return UriTypeHttps
	} else if strings.HasPrefix(uri, "ssh://") || strings.Index(uri, "@") != -1 {
		return UriTypeSSH
	} else {
		return UriTypeKnown
	}
}

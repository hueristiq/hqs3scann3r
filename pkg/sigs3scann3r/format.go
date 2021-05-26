package sigs3scann3r

import (
	"regexp"
	"strings"
)

const (
	s3url  = "^s3://*"
	s3vh   = "(s3.*)\\.amazonaws\\.com$"
	s3path = "^(s3-|s3\\.)?(s3.*)\\.amazonaws\\.com"
)

func Format(bucket, format string) string {
	target := strings.Replace(bucket, "http://", "", 1)
	target = strings.Replace(target, "https://", "", 1)
	target = strings.Replace(target, "s3://", "s3:////", 1)
	target = strings.Replace(target, "//", "", 1)

	var s3name string

	if path, _ := regexp.MatchString(s3path, target); path {
		target = strings.Replace(target, "s3.amazonaws.com/", "", 1)
		target = strings.Split(target, "/")[0]
		s3name = target
	} else if vhost, _ := regexp.MatchString(s3vh, target); vhost {
		target = strings.Replace(target, ".s3.amazonaws.com", "", 1)
		target = strings.Split(target, "/")[0]
		s3name = target
	} else if url, _ := regexp.MatchString(s3url, target); url {
		target = strings.Replace(target, "s3://", "", 1)
		target = strings.Split(target, "/")[0]
		s3name = target
	} else {
		s3name = target
	}

	var result string

	switch format {
	case "path":
		result = "https://s3.amazonaws.com/" + s3name
	case "name":
		result = s3name
	case "url":
		result = "s3://" + s3name
	case "vhost":
		result = s3name + ".s3.amazonaws.com"
	default:
		result = bucket
	}

	return result
}

package config

import (
    util "github.com/devforfu/fastgoing"
)

var RegexURL = util.MustRegexpMap(`https?:\/\/(?P<origin>[\w]+)\.(com|org|io|ru)\/[\w\W]*`)
var RegexMDFile = util.MustRegexpMap(`(?P<year>\d{4})_(?P<month>\d{2})_(?P<day>\d{2})_(?P<name>[\w\W]+)\.md$`)
var RegexJSONPreamble = util.MustRegexpMap("^```json\n(?P<preamble>[\\w\\W]+)```")

const FormatWrappedPostContent = `
{{ define "title" }}%s{{ end }}
{{ define "content" }}
%s
{{ end }}
`

const FormatVerboseDate = "Jan 02, 2006"
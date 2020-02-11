package config

import (
    util "github.com/devforfu/fastgoing"
)

var RegexURL = util.MustRegexpMap(`https?:\/\/(?P<origin>[\w]+)\.(com|org|io|ru)\/[\w\W]*`)
var RegexMDFile = util.MustRegexpMap(`(?P<year>\d{4})-(?P<month>\d{2})-(?P<day>\d{2})-(?P<name>[\w\W]+)\.md$`)
var RegexJSONPreamble = util.MustRegexpMap("^```json\n(?P<preamble>[\\w\\W]+)```")

const FormatWrappedPostContent = `
{{ define "title" }}%s{{ end }}
{{ define "content" }}
<div class="post-page">
<h2 class="post-title">%s</h2>
<div class="post-content">
%s
</div>
</div>
{{ end }}
`

const FormatVerboseDate = "Jan 02, 2006"
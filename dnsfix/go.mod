module dnsfix

require (
	github.com/miekg/dns v1.1.3
	github.com/sirupsen/logrus v1.6.0
	gopkg.in/yaml.v2 v2.2.2
	local.com/abc/game/util v0.0.0
)

replace local.com/abc/game/util => ../util

go 1.14

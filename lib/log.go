package lib

import "github.com/op/go-logging"

var log = logging.MustGetLogger("domain-proxy")

func init() {
	format := logging.MustStringFormatter(
		`Domain-Proxy %{color} %{shortfunc} %{level:.4s} %{shortfile}
%{id:03x}%{color:reset} %{message}`,
	)
	logging.SetFormatter(format)
}

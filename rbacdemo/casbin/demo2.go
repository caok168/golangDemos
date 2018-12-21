package main

import (
	"github.com/casbin/casbin"
	"github.com/casbin/casbin/persist/file-adapter"
)

func main() {
	enforcer := casbin.NewEnforcer()

	adapter := fileadapter.NewFilteredAdapter("examples/rbac_with_domains_policy.csv")
	enforcer.InitWithAdapter("examples/rbac_with_domains_model.conf", adapter)

	filter := &fileadapter.Filter{
		P: []string{"", "domain1"},
		G: []string{"", "", "domain1"},
	}
	enforcer.LoadFilteredPolicy(filter)
}

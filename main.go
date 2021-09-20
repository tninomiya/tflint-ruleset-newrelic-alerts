package main

import (
	"github.com/terraform-linters/tflint-plugin-sdk/plugin"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/tninomiya/tflint-ruleset-newrelic-alerts/rules"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		RuleSet: &tflint.BuiltinRuleSet{
			Name:    "newrelic_nrql_alerts",
			Version: "0.1.0",
			Rules: []tflint.Rule{
			},
		},
	})
}

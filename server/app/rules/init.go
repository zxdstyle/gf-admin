package rules

import "github.com/gogf/gf/v2/util/gvalid"

type ValidatorRule interface {
	Name() string
	Rule() gvalid.RuleFunc
}

var rules = []ValidatorRule{
	existsRule, uniqueRule,
}

func init() {
	for _, rule := range rules {
		gvalid.RegisterRule(rule.Name(), rule.Rule())
	}
}

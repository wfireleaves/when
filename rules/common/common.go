package common

import "github.com/wfireleaves/when/rules"

var All = []rules.Rule{
	SlashDMY(rules.Override),
}

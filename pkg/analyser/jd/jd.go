package jd

import (
	"fmt"
	"github.com/deb-sig/double-entry-generator/pkg/config"
	"github.com/deb-sig/double-entry-generator/pkg/ir"
	"github.com/deb-sig/double-entry-generator/pkg/util"
	"strings"
)

type Jd struct {
}

// GetAllCandidateAccounts returns all accounts defined in config.
func (j Jd) GetAllCandidateAccounts(cfg *config.Config) map[string]bool {
	uniqMap := make(map[string]bool)
	if cfg.Jd == nil || len(cfg.Jd.Rules) == 0 {
		return uniqMap
	}
	for _, r := range cfg.Jd.Rules {
		if r.MethodAccount != nil {
			uniqMap[*r.MethodAccount] = true
		}
		if r.TargetAccount != nil {
			uniqMap[*r.TargetAccount] = true
		}
	}
	uniqMap[cfg.DefaultPlusAccount] = true
	uniqMap[cfg.DefaultMinusAccount] = true
	fmt.Print(uniqMap)
	return uniqMap
}

func (j Jd) GetAccountsAndTags(o *ir.Order, cfg *config.Config, target, provider string) (bool, string, string, map[ir.Account]string, []string) {
	ignore := false

	if cfg.Jd == nil || len(cfg.Jd.Rules) == 0 {
		return ignore, cfg.DefaultMinusAccount, cfg.DefaultPlusAccount, nil, nil
	}
	resMinus := cfg.DefaultMinusAccount
	resPlus := cfg.DefaultPlusAccount
	var extraAccounts map[ir.Account]string
	var tags = make([]string, 0)

	for _, r := range cfg.Jd.Rules {
		match := true
		// get separator
		sep := ","

		matchFunc := util.SplitFindContains
		if r.FullMatch {
			matchFunc = util.SplitFindEquals
		}

		if r.Peer != nil {
			match = matchFunc(*r.Peer, o.Peer, sep, match)
		}
		if r.Type != nil {
			match = matchFunc(*r.Type, o.TypeOriginal, sep, match)
		}
		if r.Item != nil {
			match = matchFunc(*r.Item, o.Item, sep, match)
		}
		if r.Method != nil {
			match = matchFunc(*r.Method, o.Method, sep, match)
		}

		if match {
			if r.Ignore {
				ignore = true
				break
			}

			// Support multiple matches, like one rule matches the
			// minus account, the other rule matches the plus account.
			if r.TargetAccount != nil {
				if o.Type == ir.TypeRecv {
					resMinus = *r.TargetAccount
				} else {
					resPlus = *r.TargetAccount
				}
			}
			if r.MethodAccount != nil {
				if o.Type == ir.TypeRecv {
					resPlus = *r.MethodAccount
				} else {
					resMinus = *r.MethodAccount
				}
			}
		}
	}

	if strings.HasPrefix(o.Item, "退款-") && ir.TypeRecv != o.Type {
		return ignore, resPlus, resMinus, extraAccounts, tags
	}
	return ignore, resMinus, resPlus, extraAccounts, tags
}

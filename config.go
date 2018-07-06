package main

import pf "github.com/ipopov/pricefetch/lib"

var config = []pf.SecurityFetcher{
	pf.VanguardFetcher{
		[]pf.VanguardFund{
			{7555, "bond_cit"},
			{7553, "extended_cit"},
			{7554, "s_p_cit"},
			{7556, "ex_us_cit"},
			{1870, "vtpsx"},
			{569, "vtiax"},
			{585, "vtsax"},
		},
	},
	pf.IexFetcher{
		[]string{
			"vxus",
			"veu",
			"vti",
			"vwo",
			"schb",
			"vt",
			"spdw",
			"goog",
		},
	},
}

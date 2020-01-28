package main

import pf "github.com/ipopov/pricefetch/lib"

var config = []pf.SecurityFetcher{
	pf.VanguardFetcher{
		[]pf.VanguardFund{
			{7555, "bond_cit"},
			{7553, "extended_cit"},
			{7554, "s_p_cit"},
			{1682, "cit_tret"},
			{1870, "vtpsx"},
			{569, "vtiax"},
			{585, "vtsax"},
			{1457, "vdipx"},
			{1865, "vemrx"},
			{570, "vfwax"},
			{31, "vmltx"},
			{76, "vnytx"},
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
			"sptm",
			"goog",
			"onem",
		},
	},
}

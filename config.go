package main

import pf "github.com/ipopov/pricefetch/lib"

var config = []pf.Security{
	pf.VanguardFund{7555, "bond_cit"},
	pf.VanguardFund{7553, "extended_cit"},
	pf.VanguardFund{7554, "s_p_cit"},
	pf.VanguardFund{7556, "ex_us_cit"},
	pf.VanguardFund{1870, "vtpsx"},
	pf.VanguardFund{569, "vtiax"},
	pf.VanguardFund{585, "vtsax"},
	pf.IexStock{"vxus"},
	pf.IexStock{"veu"},
	pf.IexStock{"vti"},
	pf.IexStock{"vwo"},
	pf.IexStock{"schb"},
	pf.IexStock{"vt"},
	pf.IexStock{"spdw"},
	pf.IexStock{"goog"},
}

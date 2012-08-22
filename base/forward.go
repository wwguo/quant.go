// Copyright 2012 Wei-Wei Guo. All rights reserved.
// Use of this source code is governed by a GPL-3
// license that can be found in the LICENSE file.

package quant

// return of forword contract.
func Forward(k, st float64) (long, short float64) {
	long = st - k
	short = k - st
	return
}

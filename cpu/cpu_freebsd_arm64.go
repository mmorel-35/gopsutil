// SPDX-License-Identifier: BSD-3-Clause
package cpu

import "structs"

type cpuTimes struct {
	_    structs.HostLayout
	User uint64
	Nice uint64
	Sys  uint64
	Intr uint64
	Idle uint64
}

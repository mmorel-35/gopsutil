// SPDX-License-Identifier: BSD-3-Clause
package cpu

import "structs"

type cpuTimes struct {
	_    structs.HostLayout
	User uint32
	Nice uint32
	Sys  uint32
	Intr uint32
	Idle uint32
}

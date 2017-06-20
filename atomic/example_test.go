// Copyright 2016 zxfonline@sina.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package atomic_test

import (
	"fmt"

	"github.com/zxfonline/atomic"
)

func Example() {
	// Uint32 is a thin wrapper around the primitive uint32 type.
	var atom atomic.Uint32

	// The wrapper ensures that all operations are atomic.
	atom.Store(42)
	fmt.Println(atom.Inc())
	fmt.Println(atom.CAS(43, 0))
	bb, _ := atom.MarshalBinary()
	atom.UnmarshalBinary(bb)
	fmt.Println(atom.Load())

	// Output:
	// 43
	// true
	// 0
}

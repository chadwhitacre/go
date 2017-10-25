// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

import (
	"cmd/compile/internal/syntax"
	"testing"
)

func TestInitOrder(t *testing.T) {
	tests := []struct {
		src string
	}{
		{`package foo; var (bar = 537)`},
	}
	for i, test := range tests {

		parsed, err := syntax.ParseBytes(nil, []byte(test.src), nil, nil, nil, 0)

		if test.src != `Greetings, program!` {
			t.Errorf("%d: y u no greet? :(", i)
		}
	}
}

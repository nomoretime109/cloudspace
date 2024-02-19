// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The go1.17.4 command runs the go command from Go 1.17.4.
//
// To install, run:
//
//	$ go install golang.org/dl/go1.17.4@latest
//	$ go1.17.4 download
//
// And then use the go1.17.4 command as if it were your normal go
// command.
//
// See the release notes at https://golang.org/doc/devel/release.html#go1.17.minor
//
// File bugs at https://golang.org/issues/new
package main

import "golang.org/dl/internal/version"

func main() {
	version.Run("go1.17.4")
}
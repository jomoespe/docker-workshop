// Copyright 2019 Jomoespe. All rights reserved.
// Use of this source code is governed by a The Unlicense license that can be
// found in the LICENSE file.

package main

import (
	"fmt"
	"os"
)

func main() {
	salutation := "world"
	if len(os.Args) >= 2 {
		salutation = os.Args[1]
	}
	fmt.Printf("Hello %s! (from multi stage Dockerfile)\n", salutation)
}

// Copyright (c) 2014, Google Inc. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"crypto/sha1"
	"flag"
	"fmt"
	"os"

	"github.com/google/go-tpm/tpm"
)

func main() {
	var tpmname = flag.String("tpm", "/dev/tpm0", "The path to the TPM device to use")
	var ownerInput = flag.String("ownerauth", "", "A string to hash for owner auth (uses the well-known value if no string is supplied)")
	flag.Parse()

	f, err := os.OpenFile(*tpmname, os.O_RDWR, 0600)
	defer f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't open TPM device %s: %s\n", *tpmname, err)
		return
	}

	var ownerAuth [20]byte
	if *ownerInput != "" {
		oa := sha1.Sum([]byte(*ownerInput))
		copy(ownerAuth[:], oa[:])
	}
	if err := tpm.OwnerClear(f, ownerAuth); err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't clear the TPM using owner auth: %s\n", err)
		return
	}

	return
}

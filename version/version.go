// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package version

import "fmt"

const Version = "0.1.0"

var (
	Name      string
	GitCommit string

	HumanVersion = fmt.Sprintf("%s v%s (%s)", Name, Version, GitCommit)
)

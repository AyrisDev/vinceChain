// Copyright 2023 Vince Foundation
// This file is part of the Vince Network packages.
//
// Vince is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The Vince packages are distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with The Vince packages. If not, see https://github.com/AyrisDev/vinceChain/blob/main/LICENSE
package version

import (
	"fmt"
	"runtime"
)

var (
	AppVersion = ""
	GitCommit  = ""
	BuildDate  = ""

	GoVersion = ""
	GoArch    = ""
)

func init() {
	if len(AppVersion) == 0 {
		AppVersion = "dev"
	}

	GoVersion = runtime.Version()
	GoArch = runtime.GOARCH
}

func Version() string {
	return fmt.Sprintf(
		"Version %s (%s)\nCompiled at %s using Go %s (%s)",
		AppVersion,
		GitCommit,
		BuildDate,
		GoVersion,
		GoArch,
	)
}

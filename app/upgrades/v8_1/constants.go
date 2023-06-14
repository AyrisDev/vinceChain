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

package v81

const (
	// UpgradeName is the shared upgrade plan name for mainnet and testnet
	UpgradeName = "v8.1.0"
	// MainnetUpgradeHeight defines the vince mainnet block height on which the upgrade will take place
	MainnetUpgradeHeight = 4_500_000
	// TestnetUpgradeHeight defines the vince testnet block height on which the upgrade will take place
	TestnetUpgradeHeight = 5_280_000
	// UpgradeInfo defines the binaries that will be used for the upgrade
	UpgradeInfo = `'{"binaries":{"darwin/arm64":"https://github.com/AyrisDev/vinceChain/releases/download/v8.1.0/vince_8.1.0_Darwin_arm64.tar.gz","darwin/x86_64":"https://github.com/AyrisDev/vinceChain/releases/download/v8.1.0/vince_8.1.0_Darwin_x86_64.tar.gz","linux/arm64":"https://github.com/AyrisDev/vinceChain/releases/download/v8.1.0/vince_8.1.0_Linux_arm64.tar.gz","linux/amd64":"https://github.com/AyrisDev/vinceChain/releases/download/v8.1.0/vince_8.1.0_Linux_x86_64.tar.gz","windows/x86_64":"https://github.com/AyrisDev/vinceChain/releases/download/v8.1.0/vince_8.1.0_Windows_x86_64.zip"}}'`
)

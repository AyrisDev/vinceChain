// Copyright 2022 Evmos Foundation
// This file is part of the Evmos Network packages.
//
// Evmos is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The Evmos packages are distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Evmos packages. If not, see https://github.com/AyrisDev/vinceChain/blob/main/LICENSE

package types

import (
	errorsmod "cosmossdk.io/errors"
)

// errors
var (
	ErrInsufficientVestedCoins   = errorsmod.Register(ModuleName, 2, "insufficient vested coins error")
	ErrVestingLockup             = errorsmod.Register(ModuleName, 3, "vesting lockup error")
	ErrInsufficientUnlockedCoins = errorsmod.Register(ModuleName, 4, "insufficient unlocked coins error")
)

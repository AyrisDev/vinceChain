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

package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	vincetypes "github.com/AyrisDev/vinceChain/v12/types"
)

// CalculateEpochProvisions returns mint provision per epoch
func CalculateEpochMintProvision(
	params Params,
	period uint64,
	epochsPerPeriod int64,
	bondedRatio sdk.Dec,
) sdk.Dec {
	x := period                                              // period
	a := params.ExponentialCalculation.A                     // initial value
	r := params.ExponentialCalculation.R                     // reduction factor
	c := params.ExponentialCalculation.C                     // long term inflation
	bTarget := params.ExponentialCalculation.BondingTarget   // bonding target
	maxVariance := params.ExponentialCalculation.MaxVariance // max percentage that inflation can be increased by

	// exponentialDecay := a * (1 - r) ^ x + c
	decay := sdk.OneDec().Sub(r)
	exponentialDecay := a.Mul(decay.Power(x)).Add(c)

	// bondingIncentive doesn't increase beyond bonding target (0 < b < bonding_target)
	if bondedRatio.GTE(bTarget) {
		bondedRatio = bTarget
	}

	// bondingIncentive = 1 + max_variance - bondingRatio * (max_variance / bonding_target)
	sub := bondedRatio.Mul(maxVariance.Quo(bTarget))
	bondingIncentive := sdk.OneDec().Add(maxVariance).Sub(sub)

	// periodProvision = exponentialDecay * bondingIncentive
	periodProvision := exponentialDecay.Mul(bondingIncentive)

	// epochProvision = periodProvision / epochsPerPeriod
	epochProvision := periodProvision.Quo(sdk.NewDec(epochsPerPeriod))

	// Multiply epochMintProvision with power reduction (10^18 for vince) as the
	// calculation is based on `vince` and the issued tokens need to be given in
	// `avce`
	epochProvision = epochProvision.Mul(sdk.NewDecFromInt(vincetypes.PowerReduction))
	return epochProvision
}

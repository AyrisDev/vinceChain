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
package eip712

import (
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

// createEIP712Domain creates the typed data domain for the given chainID.
func createEIP712Domain(chainID uint64) apitypes.TypedDataDomain {
	domain := apitypes.TypedDataDomain{
		Name:              "Cosmos Web3",
		Version:           "1.0.0",
		ChainId:           math.NewHexOrDecimal256(int64(chainID)), // #nosec G701
		VerifyingContract: "cosmos",
		Salt:              "0",
	}

	return domain
}

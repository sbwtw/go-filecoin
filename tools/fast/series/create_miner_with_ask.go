package series

import (
	"context"
	"math/big"

	"github.com/filecoin-project/specs-actors/actors/abi"

	"github.com/sbwtw/go-filecoin/internal/app/go-filecoin/porcelain"
	"github.com/sbwtw/go-filecoin/tools/fast"
)

// CreateStorageMinerWithAsk setups a miner and sets an ask price. The created ask is
// returned.
func CreateStorageMinerWithAsk(ctx context.Context, miner *fast.Filecoin, collateral *big.Int, price *big.Float, expiry *big.Int, sectorSize abi.SectorSize) (porcelain.Ask, error) {
	// Create miner
	_, err := miner.MinerCreate(ctx, collateral, fast.AOSectorSize(sectorSize), fast.AOPrice(big.NewFloat(1.0)), fast.AOLimit(300))
	if err != nil {
		return porcelain.Ask{}, err
	}

	return SetPriceGetAsk(ctx, miner, price, expiry)
}

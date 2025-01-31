package tortoise

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/common/util"
	"github.com/spacemeshos/go-spacemesh/log/logtest"
)

func TestComputeThreshold(t *testing.T) {
	genesis := types.GetEffectiveGenesis()

	for _, tc := range []struct {
		desc                      string
		config                    Config
		processed, last, verified types.LayerID
		mode                      mode
		epochWeight               map[types.EpochID]util.Weight

		expectedLocal  util.Weight
		expectedGlobal util.Weight
	}{
		{
			desc: "WindowIsNotShorterThanProcessed",
			config: Config{
				LocalThreshold:                  big.NewRat(1, 2),
				GlobalThreshold:                 big.NewRat(1, 2),
				VerifyingModeVerificationWindow: 1,
			},
			processed: genesis.Add(4),
			last:      genesis.Add(4),
			verified:  genesis,
			epochWeight: map[types.EpochID]util.Weight{
				2: util.WeightFromUint64(10),
			},
			expectedLocal:  util.WeightFromUint64(5),
			expectedGlobal: util.WeightFromUint64(20),
		},
		{
			desc: "VerifyingLimitIsUsed",
			config: Config{
				LocalThreshold:                  big.NewRat(1, 2),
				GlobalThreshold:                 big.NewRat(1, 2),
				VerifyingModeVerificationWindow: 2,
			},
			processed: genesis.Add(1),
			last:      genesis.Add(4),
			verified:  genesis,
			epochWeight: map[types.EpochID]util.Weight{
				2: util.WeightFromUint64(10),
			},
			expectedLocal:  util.WeightFromUint64(5),
			expectedGlobal: util.WeightFromUint64(15),
		},
		{
			desc: "FullLimitIsUsed",
			config: Config{
				LocalThreshold:             big.NewRat(1, 2),
				GlobalThreshold:            big.NewRat(1, 2),
				FullModeVerificationWindow: 3,
			},
			processed: genesis.Add(1),
			last:      genesis.Add(4),
			verified:  genesis,
			mode:      mode{}.toggleMode(),
			epochWeight: map[types.EpochID]util.Weight{
				2: util.WeightFromUint64(10),
			},
			expectedLocal:  util.WeightFromUint64(5),
			expectedGlobal: util.WeightFromUint64(20),
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			local, global := computeThresholds(
				logtest.New(t), tc.config, tc.mode,
				tc.verified.Add(1), tc.last, tc.processed, tc.epochWeight,
			)
			require.Equal(t, tc.expectedLocal.String(), local.String())
			require.Equal(t, tc.expectedGlobal.String(), global.String())
		})
	}
}

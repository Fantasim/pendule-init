package pcommon

import (
	"errors"
	"strconv"
	"strings"
)

type AssetType string

var Asset = struct {
	SPOT_PRICE  AssetType
	SPOT_VOLUME AssetType

	FUTURES_PRICE  AssetType
	FUTURES_VOLUME AssetType

	BOOK_DEPTH_P1 AssetType
	BOOK_DEPTH_P2 AssetType
	BOOK_DEPTH_P3 AssetType
	BOOK_DEPTH_P4 AssetType
	BOOK_DEPTH_P5 AssetType

	BOOK_DEPTH_M1 AssetType
	BOOK_DEPTH_M2 AssetType
	BOOK_DEPTH_M3 AssetType
	BOOK_DEPTH_M4 AssetType
	BOOK_DEPTH_M5 AssetType

	METRIC_SUM_OPEN_INTEREST                 AssetType
	METRIC_COUNT_TOP_TRADER_LONG_SHORT_RATIO AssetType
	METRIC_SUM_TOP_TRADER_LONG_SHORT_RATIO   AssetType
	METRIC_COUNT_LONG_SHORT_RATIO            AssetType
	METRIC_SUM_TAKER_LONG_SHORT_VOL_RATIO    AssetType

	CIRCULATING_SUPPLY AssetType
	RSI                AssetType
}{
	SPOT_PRICE:  "spot_price",
	SPOT_VOLUME: "spot_volume",

	FUTURES_PRICE:  "futures_price",
	FUTURES_VOLUME: "futures_volume",

	BOOK_DEPTH_P1: "bd-p1",
	BOOK_DEPTH_P2: "bd-p2",
	BOOK_DEPTH_P3: "bd-p3",
	BOOK_DEPTH_P4: "bd-p4",
	BOOK_DEPTH_P5: "bd-p5",

	BOOK_DEPTH_M1: "bd-m1",
	BOOK_DEPTH_M2: "bd-m2",
	BOOK_DEPTH_M3: "bd-m3",
	BOOK_DEPTH_M4: "bd-m4",
	BOOK_DEPTH_M5: "bd-m5",

	METRIC_SUM_OPEN_INTEREST:                 "metrics_sum_open_interest",
	METRIC_COUNT_TOP_TRADER_LONG_SHORT_RATIO: "metrics_count_toptrader_long_short_ratio",
	METRIC_SUM_TOP_TRADER_LONG_SHORT_RATIO:   "metrics_sum_toptrader_long_short_ratio",
	METRIC_COUNT_LONG_SHORT_RATIO:            "metrics_count_long_short_ratio",
	METRIC_SUM_TAKER_LONG_SHORT_VOL_RATIO:    "metrics_sum_taker_long_short_vol_ratio",

	CIRCULATING_SUPPLY: "circulating_supply",
	RSI:                "rsi",
}

func (asset AssetType) GetBookDepthAssetPercentage() (int, error) {
	if asset == Asset.BOOK_DEPTH_M1 || asset == Asset.BOOK_DEPTH_M2 || asset == Asset.BOOK_DEPTH_M3 || asset == Asset.BOOK_DEPTH_M4 || asset == Asset.BOOK_DEPTH_M5 ||
		asset == Asset.BOOK_DEPTH_P1 || asset == Asset.BOOK_DEPTH_P2 || asset == Asset.BOOK_DEPTH_P3 || asset == Asset.BOOK_DEPTH_P4 || asset == Asset.BOOK_DEPTH_P5 {
		lastChar := asset[len(asset)-1:]
		percent, err := strconv.Atoi(string(lastChar))
		if err != nil {
			return 0, err
		}

		isPlus := strings.HasPrefix(string(asset), "bd-p")
		isMinus := strings.HasPrefix(string(asset), "bd-m")
		if isPlus {
			return percent, nil
		} else if isMinus {
			return -percent, nil
		}
	}
	return 0, errors.New("invalid asset")
}

type AssetStateConfig struct {
	ID       AssetType
	Key      [2]byte
	DataType DataType
}

var DEFAULT_ASSETS = map[AssetType]AssetStateConfig{
	//binance spot trades
	Asset.SPOT_PRICE:  {Asset.SPOT_PRICE, [2]byte{0, 0}, UNIT},
	Asset.SPOT_VOLUME: {Asset.SPOT_VOLUME, [2]byte{0, 1}, QUANTITY},

	//binance book depth
	Asset.BOOK_DEPTH_P1: {Asset.BOOK_DEPTH_P1, [2]byte{0, 2}, UNIT},
	Asset.BOOK_DEPTH_P2: {Asset.BOOK_DEPTH_P2, [2]byte{0, 3}, UNIT},
	Asset.BOOK_DEPTH_P3: {Asset.BOOK_DEPTH_P3, [2]byte{0, 4}, UNIT},
	Asset.BOOK_DEPTH_P4: {Asset.BOOK_DEPTH_P4, [2]byte{0, 5}, UNIT},
	Asset.BOOK_DEPTH_P5: {Asset.BOOK_DEPTH_P5, [2]byte{0, 6}, UNIT},

	Asset.BOOK_DEPTH_M1: {Asset.BOOK_DEPTH_M1, [2]byte{0, 7}, UNIT},
	Asset.BOOK_DEPTH_M2: {Asset.BOOK_DEPTH_M2, [2]byte{0, 8}, UNIT},
	Asset.BOOK_DEPTH_M3: {Asset.BOOK_DEPTH_M3, [2]byte{0, 9}, UNIT},
	Asset.BOOK_DEPTH_M4: {Asset.BOOK_DEPTH_M4, [2]byte{0, 10}, UNIT},
	Asset.BOOK_DEPTH_M5: {Asset.BOOK_DEPTH_M5, [2]byte{0, 11}, UNIT},

	Asset.METRIC_SUM_OPEN_INTEREST:                 {Asset.METRIC_SUM_OPEN_INTEREST, [2]byte{0, 12}, UNIT},
	Asset.METRIC_COUNT_TOP_TRADER_LONG_SHORT_RATIO: {Asset.METRIC_COUNT_TOP_TRADER_LONG_SHORT_RATIO, [2]byte{0, 13}, UNIT},
	Asset.METRIC_SUM_TOP_TRADER_LONG_SHORT_RATIO:   {Asset.METRIC_SUM_TOP_TRADER_LONG_SHORT_RATIO, [2]byte{0, 14}, UNIT},
	Asset.METRIC_COUNT_LONG_SHORT_RATIO:            {Asset.METRIC_COUNT_LONG_SHORT_RATIO, [2]byte{0, 15}, UNIT},
	Asset.METRIC_SUM_TAKER_LONG_SHORT_VOL_RATIO:    {Asset.METRIC_SUM_TAKER_LONG_SHORT_VOL_RATIO, [2]byte{0, 16}, UNIT},

	Asset.CIRCULATING_SUPPLY: {Asset.CIRCULATING_SUPPLY, [2]byte{0, 17}, UNIT},

	Asset.FUTURES_PRICE:  {Asset.FUTURES_PRICE, [2]byte{0, 18}, UNIT},
	Asset.FUTURES_VOLUME: {Asset.FUTURES_VOLUME, [2]byte{0, 19}, QUANTITY},

	Asset.RSI: {Asset.RSI, [2]byte{0, 20}, POINT},
}

var ASSET_LIST_WITHOUT_DEPENDENCIES = []AssetType{
	Asset.SPOT_PRICE,
	Asset.SPOT_VOLUME,
	Asset.FUTURES_PRICE,
	Asset.FUTURES_VOLUME,
	Asset.BOOK_DEPTH_P1,
	Asset.BOOK_DEPTH_P2,
	Asset.BOOK_DEPTH_P3,
	Asset.BOOK_DEPTH_P4,
	Asset.BOOK_DEPTH_P5,
	Asset.BOOK_DEPTH_M1,
	Asset.BOOK_DEPTH_M2,
	Asset.BOOK_DEPTH_M3,
	Asset.BOOK_DEPTH_M4,
	Asset.BOOK_DEPTH_M5,
	Asset.METRIC_SUM_OPEN_INTEREST,
	Asset.METRIC_COUNT_TOP_TRADER_LONG_SHORT_RATIO,
	Asset.METRIC_SUM_TOP_TRADER_LONG_SHORT_RATIO,
	Asset.METRIC_COUNT_LONG_SHORT_RATIO,

	Asset.METRIC_SUM_TAKER_LONG_SHORT_VOL_RATIO,
	Asset.CIRCULATING_SUPPLY,
}

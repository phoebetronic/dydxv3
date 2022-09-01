package config

import (
	"math/big"
)

type Config struct {
	FieldPrime     *big.Int     `json:"FIELD_PRIME"`
	FieldGen       *big.Int     `json:"FIELD_GEN"`
	ECOrder        *big.Int     `json:"EC_ORDER"`
	Alpha          *big.Int     `json:"ALPHA"`
	Beta           *big.Int     `json:"BETA"`
	ConstantPoints [][]*big.Int `json:"CONSTANT_POINTS"`
}

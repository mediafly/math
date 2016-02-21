package math

import (
	"time"
)

func MinByte(a byte, b byte) byte {
	if a < b {
		return a
	} else {
		return b
	}
}

func MaxByte(a byte, b byte) byte {
	if a > b {
		return a
	} else {
		return b
	}
}

func MinDuration(a time.Duration, b time.Duration) time.Duration {
	if a < b {
		return a
	} else {
		return b
	}
}

func MaxDuration(a time.Duration, b time.Duration) time.Duration {
	if a > b {
		return a
	} else {
		return b
	}
}

func MinTime(a time.Time, b time.Time) time.Time {
	if a.Before(b) {
		return a
	} else {
		return b
	}
}

func MaxTime(a time.Time, b time.Time) time.Time {
	if a.After(b) {
		return a
	} else {
		return b
	}
}

func MinFloat32(a float32, b float32) float32 {
	if a < b {
		return a
	} else {
		return b
	}
}

func MaxFloat32(a float32, b float32) float32 {
	if a > b {
		return a
	} else {
		return b
	}
}

func MinFloat64(a float64, b float64) float64 {
	if a < b {
		return a
	} else {
		return b
	}
}

func MaxFloat64(a float64, b float64) float64 {
	if a > b {
		return a
	} else {
		return b
	}
}

func MinInt(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func MaxInt(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func MinInt16(a int16, b int16) int16 {
	if a < b {
		return a
	} else {
		return b
	}
}

func MaxInt16(a int16, b int16) int16 {
	if a > b {
		return a
	} else {
		return b
	}
}

func MinInt32(a int32, b int32) int32 {
	if a < b {
		return a
	} else {
		return b
	}
}

func MaxInt32(a int32, b int32) int32 {
	if a > b {
		return a
	} else {
		return b
	}
}

func MinInt64(a int64, b int64) int64 {
	if a < b {
		return a
	} else {
		return b
	}
}

func MaxInt64(a int64, b int64) int64 {
	if a > b {
		return a
	} else {
		return b
	}
}

func MinInt8(a int8, b int8) int8 {
	if a < b {
		return a
	} else {
		return b
	}
}

func MaxInt8(a int8, b int8) int8 {
	if a > b {
		return a
	} else {
		return b
	}
}

func MinUint16(a uint16, b uint16) uint16 {
	if a < b {
		return a
	} else {
		return b
	}
}

func MaxUint16(a uint16, b uint16) uint16 {
	if a > b {
		return a
	} else {
		return b
	}
}

func MinUint32(a uint32, b uint32) uint32 {
	if a < b {
		return a
	} else {
		return b
	}
}

func MaxUint32(a uint32, b uint32) uint32 {
	if a > b {
		return a
	} else {
		return b
	}
}

func MinUint64(a uint64, b uint64) uint64 {
	if a < b {
		return a
	} else {
		return b
	}
}

func MaxUint64(a uint64, b uint64) uint64 {
	if a > b {
		return a
	} else {
		return b
	}
}

func MinUint8(a uint8, b uint8) uint8 {
	if a < b {
		return a
	} else {
		return b
	}
}

func MaxUint8(a uint8, b uint8) uint8 {
	if a > b {
		return a
	} else {
		return b
	}
}

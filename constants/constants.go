package constants

import (
	m "math"
	"time"
)

const MaxInt = int(m.MaxInt64)
const MinInt = int(m.MinInt64)

const MaxUint8 = uint8(m.MaxUint8)
const MinUint8 = uint8(0)

const MaxUint16 = uint16(m.MaxUint16)
const MinUint16 = uint16(0)

const MaxUint32 = uint32(m.MaxUint32)
const MinUint32 = uint32(0)

const MaxUint64 = uint64(m.MaxUint64)
const MinUint64 = uint64(0)

const MaxInt8 = int8(m.MaxInt8)
const MinInt8 = int8(m.MinInt8)

const MaxInt16 = int16(m.MaxInt16)
const MinInt16 = int16(m.MinInt16)

const MaxInt32 = int32(m.MaxInt32)
const MinInt32 = int32(m.MinInt32)

const MaxInt64 = int64(m.MaxInt64)
const MinInt64 = int64(m.MinInt64)

const MaxDuration = time.Duration(MaxInt64)
const MinDuration = time.Duration(MinInt64)

// http://stackoverflow.com/questions/25065055/what-is-the-maximum-time-time-in-go
var MaxTime = time.Unix(1<<63-62135596801, 999999999)
var MinTime = time.Time{}

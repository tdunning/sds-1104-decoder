package internal

import (
	"fmt"
	"math"
)

const (
	padding = 0x800 - 0x11C
)

type Header struct {
	ChOn              [4]uint32
	VoltsPerDivision  [4]ScaledValue
	VerticalOffset    [4]ScaledValue
	DigitalOn         uint32
	DigitalChannelOn  [16]uint32
	TimePerDivision   ScaledValue
	TimeDelay         ScaledValue
	AnalogSamples     uint32
	AnalogSampleRate  ScaledValue
	DigitalSamples    uint32
	DigitalSampleRate ScaledValue
	Reserved          [padding]uint8
}

type ScaledValue struct {
	Value float64
	Scale MagnitudeType
	Unit  UnitType
}

func (sv ScaledValue) String() string {
	return fmt.Sprintf("%f %v %v", sv.Value, sv.Unit.String(), sv.Scale.String())
}

func (sv ScaledValue) Reduce() float64 {
	return sv.Value * math.Pow10(3*(int(sv.Scale)-8))
}

//go:generate stringer -type=UnitType
type UnitType uint32

const (
	V          = UnitType(0)
	A          = UnitType(1)
	VV         = UnitType(2)
	AA         = UnitType(3)
	OU         = UnitType(4)
	W          = UnitType(5)
	SQRT_V     = UnitType(6)
	SQRT_A     = UnitType(7)
	INTEGRAL_V = UnitType(8)
	INTEGRAL_A = UnitType(9)
	DT_V       = UnitType(10)
	DT_A       = UnitType(11)
	DT_DIV     = UnitType(12)
	Hz         = UnitType(13)
	S          = UnitType(14)
	SA         = UnitType(15)
	PTS        = UnitType(16)
	NULL       = UnitType(17)
	DB         = UnitType(18)
	DBV        = UnitType(19)
	DBA        = UnitType(20)
	VPP        = UnitType(21)
	VDC        = UnitType(22)
	DBM        = UnitType(23)
)

//go:generate stringer -type=MagnitudeType
type MagnitudeType uint32

const (
	YOCTO = MagnitudeType(0)
	ZEPTO = MagnitudeType(1)
	ATTO  = MagnitudeType(2)
	FEMTO = MagnitudeType(3)
	PICO  = MagnitudeType(4)
	NANO  = MagnitudeType(5)
	MICRO = MagnitudeType(6)
	MILLI = MagnitudeType(7)
	IU    = MagnitudeType(8)
	KILO  = MagnitudeType(9)
	MEGA  = MagnitudeType(10)
	GIGA  = MagnitudeType(11)
	TERA  = MagnitudeType(12)
	PETA  = MagnitudeType(13)
)

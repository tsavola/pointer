package pointer

import (
	"testing"
)

func TestBool(t *testing.T) { var x bool; if *Bool(x) != x { t.Fail() } }
func TestByte(t *testing.T) { var x byte; if *Byte(x) != x { t.Fail() } }
func TestComplex128(t *testing.T) { var x complex128; if *Complex128(x) != x { t.Fail() } }
func TestComplex64(t *testing.T) { var x complex64; if *Complex64(x) != x { t.Fail() } }
func TestError(t *testing.T) { var x error; if *Error(x) != x { t.Fail() } }
func TestFloat32(t *testing.T) { var x float32; if *Float32(x) != x { t.Fail() } }
func TestFloat64(t *testing.T) { var x float64; if *Float64(x) != x { t.Fail() } }
func TestInt(t *testing.T) { var x int; if *Int(x) != x { t.Fail() } }
func TestInt16(t *testing.T) { var x int16; if *Int16(x) != x { t.Fail() } }
func TestInt32(t *testing.T) { var x int32; if *Int32(x) != x { t.Fail() } }
func TestInt64(t *testing.T) { var x int64; if *Int64(x) != x { t.Fail() } }
func TestInt8(t *testing.T) { var x int8; if *Int8(x) != x { t.Fail() } }
func TestRune(t *testing.T) { var x rune; if *Rune(x) != x { t.Fail() } }
func TestString(t *testing.T) { var x string; if *String(x) != x { t.Fail() } }
func TestUint(t *testing.T) { var x uint; if *Uint(x) != x { t.Fail() } }
func TestUint16(t *testing.T) { var x uint16; if *Uint16(x) != x { t.Fail() } }
func TestUint32(t *testing.T) { var x uint32; if *Uint32(x) != x { t.Fail() } }
func TestUint64(t *testing.T) { var x uint64; if *Uint64(x) != x { t.Fail() } }
func TestUint8(t *testing.T) { var x uint8; if *Uint8(x) != x { t.Fail() } }
func TestUintptr(t *testing.T) { var x uintptr; if *Uintptr(x) != x { t.Fail() } }

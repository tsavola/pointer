// Package pointer provides constructors for pointers of primitive types.
package pointer

func Bool(x bool) *bool { return &x }
func Byte(x byte) *byte { return &x }
func Complex128(x complex128) *complex128 { return &x }
func Complex64(x complex64) *complex64 { return &x }
func Error(x error) *error { return &x }
func Float32(x float32) *float32 { return &x }
func Float64(x float64) *float64 { return &x }
func Int(x int) *int { return &x }
func Int16(x int16) *int16 { return &x }
func Int32(x int32) *int32 { return &x }
func Int64(x int64) *int64 { return &x }
func Int8(x int8) *int8 { return &x }
func Rune(x rune) *rune { return &x }
func String(x string) *string { return &x }
func Uint(x uint) *uint { return &x }
func Uint16(x uint16) *uint16 { return &x }
func Uint32(x uint32) *uint32 { return &x }
func Uint64(x uint64) *uint64 { return &x }
func Uint8(x uint8) *uint8 { return &x }
func Uintptr(x uintptr) *uintptr { return &x }

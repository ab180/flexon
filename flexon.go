package flexon

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

var (
	AnyToBool   = DefaultAnyToBool
	AnyToInt    = DefaultAnyToInt
	AnyToUint   = DefaultAnyToUint
	AnyToFloat  = DefaultAnyToFloat
	AnyToString = DefaultAnyToString
)

type Bool bool

func (b *Bool) UnmarshalJSON(data []byte) error {
	any := jsoniter.Get(data)
	if err := any.LastError(); err != nil {
		return err
	}

	v, err := AnyToBool(any)
	if err != nil {
		return err
	}

	*b = Bool(v)
	return nil
}

func (b *Bool) Unwrap() bool {
	return bool(*b)
}

type Int int

func (i *Int) UnmarshalJSON(data []byte) error {
	any := jsoniter.Get(data)
	if err := any.LastError(); err != nil {
		return err
	}

	v, err := AnyToInt(any)
	if err != nil {
		return err
	}

	*i = Int(v)
	return nil
}

func (i *Int) Unwrap() int {
	return int(*i)
}

type Int8 int8

func (i *Int8) UnmarshalJSON(data []byte) error {
	any := jsoniter.Get(data)
	if err := any.LastError(); err != nil {
		return err
	}

	v, err := AnyToInt(any)
	if err != nil {
		return err
	}

	*i = Int8(v)
	return nil
}

func (i *Int8) Unwrap() int8 {
	return int8(*i)
}

type Int16 int16

func (i *Int16) UnmarshalJSON(data []byte) error {
	any := jsoniter.Get(data)
	if err := any.LastError(); err != nil {
		return err
	}

	v, err := AnyToInt(any)
	if err != nil {
		return err
	}

	*i = Int16(v)
	return nil
}

func (i *Int16) Unwrap() int16 {
	return int16(*i)
}

type Int32 int32

func (i *Int32) UnmarshalJSON(data []byte) error {
	any := jsoniter.Get(data)
	if err := any.LastError(); err != nil {
		return err
	}

	v, err := AnyToInt(any)
	if err != nil {
		return err
	}

	*i = Int32(v)
	return nil
}

func (i *Int32) Unwrap() int32 {
	return int32(*i)
}

type Int64 int64

func (i *Int64) UnmarshalJSON(data []byte) error {
	any := jsoniter.Get(data)
	if err := any.LastError(); err != nil {
		return err
	}

	v, err := AnyToInt(any)
	if err != nil {
		return err
	}

	*i = Int64(v)
	return nil
}

func (i *Int64) Unwrap() int64 {
	return int64(*i)
}

type Uint uint

func (u *Uint) UnmarshalJSON(data []byte) error {
	any := jsoniter.Get(data)
	if err := any.LastError(); err != nil {
		return err
	}

	v, err := AnyToUint(any)
	if err != nil {
		return err
	}

	*u = Uint(v)
	return nil
}

func (u *Uint) Unwrap() uint {
	return uint(*u)
}

type Uint8 uint8

func (u *Uint8) UnmarshalJSON(data []byte) error {
	any := jsoniter.Get(data)
	if err := any.LastError(); err != nil {
		return err
	}

	v, err := AnyToUint(any)
	if err != nil {
		return err
	}

	*u = Uint8(v)
	return nil
}

func (u *Uint8) Unwrap() uint8 {
	return uint8(*u)
}

type Uint16 uint16

func (u *Uint16) UnmarshalJSON(data []byte) error {
	any := jsoniter.Get(data)
	if err := any.LastError(); err != nil {
		return err
	}

	v, err := AnyToUint(any)
	if err != nil {
		return err
	}

	*u = Uint16(v)
	return nil
}

func (u *Uint16) Unwrap() uint16 {
	return uint16(*u)
}

type Uint32 uint32

func (u *Uint32) UnmarshalJSON(data []byte) error {
	any := jsoniter.Get(data)
	if err := any.LastError(); err != nil {
		return err
	}

	v, err := AnyToUint(any)
	if err != nil {
		return err
	}

	*u = Uint32(v)
	return nil
}

func (u *Uint32) Unwrap() uint32 {
	return uint32(*u)
}

type Uint64 uint64

func (u *Uint64) UnmarshalJSON(data []byte) error {
	any := jsoniter.Get(data)
	if err := any.LastError(); err != nil {
		return err
	}

	v, err := AnyToUint(any)
	if err != nil {
		return err
	}

	*u = Uint64(v)
	return nil
}

func (u *Uint64) Unwrap() uint64 {
	return uint64(*u)
}

type Float32 float32

func (f *Float32) UnmarshalJSON(data []byte) error {
	any := jsoniter.Get(data)
	if err := any.LastError(); err != nil {
		return err
	}

	v, err := AnyToFloat(any)
	if err != nil {
		return err
	}

	*f = Float32(v)
	return nil
}

func (f *Float32) Unwrap() float32 {
	return float32(*f)
}

type Float64 float64

func (f *Float64) UnmarshalJSON(data []byte) error {
	any := jsoniter.Get(data)
	if err := any.LastError(); err != nil {
		return err
	}

	v, err := AnyToFloat(any)
	if err != nil {
		return err
	}

	*f = Float64(v)
	return nil
}

func (f *Float64) Unwrap() float64 {
	return float64(*f)
}

type String string

func (s *String) UnmarshalJSON(data []byte) error {
	any := jsoniter.Get(data)
	if err := any.LastError(); err != nil {
		return err
	}

	v, err := AnyToString(any)
	if err != nil {
		return err
	}

	*s = String(v)
	return nil
}

func (s *String) Unwrap() string {
	return string(*s)
}

func DefaultAnyToBool(any jsoniter.Any) (bool, error) {
	switch any.ValueType() {
	case jsoniter.BoolValue, jsoniter.NilValue, jsoniter.NumberValue:
		return any.ToBool(), nil
	case jsoniter.StringValue:
		str := any.ToString()
		if str == "" {
			break
		}

		b, err := strconv.ParseBool(str)
		if err != nil {
			return false, fmt.Errorf("invalid json bool: %v", err)
		}
		return b, nil
	}

	return false, errors.New("invalid json bool")
}

func DefaultAnyToInt(any jsoniter.Any) (int64, error) {
	switch any.ValueType() {
	case jsoniter.NumberValue, jsoniter.NilValue, jsoniter.BoolValue:
		return any.ToInt64(), nil
	case jsoniter.StringValue:
		str := any.ToString()
		if str == "" {
			break
		}

		if idx := strings.IndexRune(str, '.'); idx != -1 {
			str = str[:idx]
		}

		i, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("invalid json int: %v", err)
		}
		return i, nil
	}

	return 0, errors.New("invalid json int")
}

func DefaultAnyToUint(any jsoniter.Any) (uint64, error) {
	switch any.ValueType() {
	case jsoniter.NilValue, jsoniter.BoolValue:
		return any.ToUint64(), nil
	case jsoniter.NumberValue, jsoniter.StringValue:
		str := any.ToString()
		if str == "" {
			break
		}

		if str[0] == '-' {
			str = str[1:]
		}
		if idx := strings.IndexRune(str, '.'); idx != -1 {
			str = str[:idx]
		}

		u, err := strconv.ParseUint(str, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("invalid json uint: %v", err)
		}
		return u, nil
	}

	return 0, fmt.Errorf("invalid json uint")
}

func DefaultAnyToFloat(any jsoniter.Any) (float64, error) {
	switch any.ValueType() {
	case jsoniter.NumberValue, jsoniter.NilValue, jsoniter.BoolValue:
		return any.ToFloat64(), nil
	case jsoniter.StringValue:
		str := any.ToString()
		if str == "" {
			break
		}

		f, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return 0, fmt.Errorf("invalid json float: %v", err)
		}
		return f, nil
	}

	return 0, errors.New("invalid json float")
}

func DefaultAnyToString(any jsoniter.Any) (string, error) {
	switch any.ValueType() {
	case jsoniter.StringValue, jsoniter.NilValue, jsoniter.NumberValue, jsoniter.BoolValue:
		return any.ToString(), nil
	}

	return "", errors.New("invalid json string")
}

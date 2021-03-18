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

type Float64 int64

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

		f, err := strconv.ParseFloat(str, 10)
		if err != nil {
			return f, fmt.Errorf("invalid json float: %v", err)
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

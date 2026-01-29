package dom

import (
	"syscall/js"
	"time"
)

// ToJSValue 转换Go类型为JS值
// 自定义结构体需要转换为 []any 或者 map[string]any 类型再传入
func ToJSValue(v interface{}) js.Value {
	switch val := v.(type) {
	case nil:
		return js.Null()
	case bool, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, float32, float64, string:
		return js.ValueOf(v)
	case error:
		return js.ValueOf(val.Error())
	case time.Time:
		date := js.Global().Get("Date")
		return date.New(val.Format(time.RFC3339))
	case []interface{}:
		arr := js.Global().Get("Array").New(len(val))
		for i, item := range val {
			arr.SetIndex(i, ToJSValue(item))
		}
		return arr
	case map[string]interface{}:
		obj := js.Global().Get("Object").New()
		for k, item := range val {
			obj.Set(k, ToJSValue(item))
		}
		return obj
	case js.Value:
		return val
	default:
		return js.Undefined()
	}
}

// FromJSValue 将JS值 val 转换为 Go 类型
func FromJSValue(val js.Value) (any, error) {
	if val.IsUndefined() || val.IsNull() {
		return nil, nil
	}
	return jsValueToInterface(val), nil
}

// jsValueToInterface 将JS值转换为Go的interface{}类型
func jsValueToInterface(val js.Value) interface{} {
	if val.IsUndefined() {
		return nil
	}
	if val.IsNull() {
		return nil
	}

	switch val.Type() {
	case js.TypeBoolean:
		return val.Bool()
	case js.TypeNumber:
		f := val.Float()
		if f == float64(int64(f)) {
			return int64(f)
		}
		return f
	case js.TypeString:
		return val.String()
	case js.TypeSymbol:
		return val.String()
	case js.TypeObject:
		if val.InstanceOf(js.Global().Get("Array")) {
			length := val.Length()
			arr := make([]interface{}, length)
			for i := 0; i < length; i++ {
				arr[i] = jsValueToInterface(val.Index(i))
			}
			return arr
		}
		if val.InstanceOf(js.Global().Get("Date")) {
			return val.String()
		}
		// 普通对象
		keys := js.Global().Get("Object").Call("keys", val)
		length := keys.Length()
		obj := make(map[string]interface{}, length)
		for i := 0; i < length; i++ {
			k := keys.Index(i).String()
			obj[k] = jsValueToInterface(val.Get(k))
		}
		return obj
	case js.TypeFunction:
		return val
	default:
		return nil
	}
}

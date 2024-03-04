package gomap

// StringAnyMap cast element value into map[string]interface{}
func (elt *Element) StringAnyMap(defaultValue ...map[string]interface{}) (map[string]interface{}, error) {
	defValue := func() *map[string]interface{} {
		if len(defaultValue) == 0 {
			return nil
		}
		return &defaultValue[0]
	}

	def := defValue()
	if elt.Value == nil {
		if def == nil {
			var v map[string]interface{}
			return v, NewWrongPathError(elt.Path)
		}
		return *def, nil
	}
	switch v := elt.Value.(type) {
	case map[string]interface{}:
		return v, nil
	default:
		return nil, NewWrongTypeError("map[string]interface{}", elt.Value)
	}
}

package gomap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGMap_Basic(t *testing.T) {
	rm := map[string]interface{}{
		"v1": 1,
		"elt": map[string]interface{}{
			"v2": 2,
		},
		"vs": "test",
		"vb": true,
	}
	m := GMap(rm)

	has := m.Has("v1")
	assert.True(t, has)

	has = m.Has("v2")
	assert.False(t, has)

	has = m.Has("elt", "v2")
	assert.True(t, has)

	has = m.Has("elt", "v3")
	assert.False(t, has)

	v, err := m.Get("v1").Int(42)
	assert.NoError(t, err)
	assert.Equal(t, 1, v)

	v, err = m.Get("elt", "v2").Int(42)
	assert.NoError(t, err)
	assert.Equal(t, 2, v)

	v, err = m.Get("a", "b", "v").Int(42)
	assert.NoError(t, err)
	assert.Equal(t, v, 42)

	v, err = m.Get("a", "b", "v").Int()
	assert.Error(t, err)
	assert.Equal(t, "Wrong path a/b/v", err.Error())
	assert.Equal(t, v, 0)

	vs, err := m.Get("vs").String("42")
	assert.NoError(t, err)
	assert.Equal(t, "test", vs)

	vs, err = m.Get("vs2").String("42")
	assert.NoError(t, err)
	assert.Equal(t, "42", vs)

	vs, err = m.Get("vs2").String()
	assert.Error(t, err)
	assert.Equal(t, "Wrong path vs2", err.Error())
	assert.Equal(t, "", vs)

	vb, err := m.Get("vb").Bool()
	assert.NoError(t, err)
	assert.Equal(t, true, vb)
}

func TestGMap_Slices(t *testing.T) {
	m := GMap{
		"v1": GSlice{1, 2},
		"elt": GMap{
			"v2": GSlice{int64(2), int64(3)},
		},
		"vs":    GSlice{"test", "test2"},
		"vui8":  []uint8("uint8"),
		"vui16": []uint16{1},
		"vui32": []uint32{1},
		"vui64": []uint64{1},
		"vb":    []bool{true},
		"vsm": []map[string]interface{}{
			{
				"key1": 1,
				"key2": "2",
			},
		},
	}
	v, err := m.Get("v1").IntSlice([]int{42})
	assert.NoError(t, err)
	assert.EqualValues(t, []int{1, 2}, v)

	v64, err := m.Get("elt", "v2").Int64Slice([]int64{42})
	assert.NoError(t, err)
	assert.Equal(t, []int64{2, 3}, v64)

	v, err = m.Get("elt", "v2").IntSlice([]int{42})
	assert.NoError(t, err)
	assert.Equal(t, []int{42}, v)

	v, err = m.Get("elt", "v2").IntSlice()
	assert.Error(t, err)
	assert.Equal(t, "Wrong type: expected int actual int64", err.Error())

	vs, err := m.Get("vs").StringSlice()
	assert.NoError(t, err)
	assert.Equal(t, []string{"test", "test2"}, vs)

	vui8, err := m.Get("vui8").Uint8Slice()
	assert.NoError(t, err)
	assert.Equal(t, []uint8("uint8"), vui8)

	vui16, err := m.Get("vui16").Uint16Slice()
	assert.NoError(t, err)
	assert.Equal(t, []uint16{1}, vui16)

	vui32, err := m.Get("vui32").Uint32Slice()
	assert.NoError(t, err)
	assert.Equal(t, []uint32{1}, vui32)

	vui64, err := m.Get("vui64").Uint64Slice()
	assert.NoError(t, err)
	assert.Equal(t, []uint64{1}, vui64)

	vb, err := m.Get("vb").BoolSlice()
	assert.NoError(t, err)
	assert.Equal(t, []bool{true}, vb)

	vsm, err := m.Get("vsm").StringAnyMapSlice()
	assert.NoError(t, err)
	assert.Equal(t, []map[string]interface{}{
		{
			"key1": 1,
			"key2": "2",
		},
	}, vsm)
}

func TestGMap_Struct(t *testing.T) {
	type Elt struct {
		V2 int `json:"v2,omitempty"`
	}

	type Elt2 struct {
		V1  int    `json:"v1,omitempty"`
		Elt Elt    `json:"elt,omitempty"`
		VS  string `json:"vs,omitempty"`
	}
	m := GMap{
		"v1": 1,
		"elt": GMap{
			"v2": 2,
		},
		"vs": "test",
	}
	elt := Elt{}
	err := m.Get("elt").Object(&elt)
	assert.NoError(t, err)
	assert.Equal(t, 2, elt.V2)

	elt2 := Elt2{}
	err = m.Get().Object(&elt2)
	assert.NoError(t, err)
	assert.EqualValues(t,
		Elt2{
			V1: 1,
			Elt: Elt{
				V2: 2,
			},
			VS: "test",
		},
		elt2)

}

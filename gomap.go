// Package gomap implements utilities to manipulate generic maps and cast elements into native Go types and native Go structures
//
// The package is really simple to use
package gomap

import (
	"encoding/json"
)

// GMap overload map with utility functions
type GMap map[string]interface{}

// GSlice overload []interface{} to simplify GMap declarations
type GSlice []interface{}

// Element an element of the map
type Element struct {
	Path  []string
	Value interface{}
}

// Get returns map elements along path
func (m GMap) Get(path ...string) *Element {
	if len(path) == 0 {
		return &Element{
			Path:  path,
			Value: m,
		}
	}
	if len(path) == 1 {
		return element(m, path, path[0])
	}
	return element(m, path, path[0]).Get(path[1:]...)
}

// Has returns if map has element on along path
func (m GMap) Has(path ...string) bool {
	if len(path) == 0 {
		return false
	}
	if len(path) == 1 {
		return element(m, path, path[0]).has()
	}
	return element(m, path, path[0]).Get(path[1:]...).has()
}

// Get returns map elements along path
func (elt *Element) Get(path ...string) *Element {
	next := func(path ...string) *Element {
		if elt.Value == nil {
			return &Element{
				Path: elt.Path,
			}
		}
		m, ok := elt.Value.(map[string]interface{})
		if ok {
			return element(m, elt.Path, path[0])
		}
		gm, ok := elt.Value.(GMap)
		if ok {
			return element(gm, elt.Path, path[0])
		}
		return &Element{
			Path: elt.Path,
		}
	}
	if len(path) == 0 {
		return elt
	}
	if len(path) == 1 {
		return next(path...)
	}
	return next(path...).Get(path[1:]...)
}

func (elt *Element) has() bool {
	return elt.Value != nil
}

func element(m GMap, path []string, key string) *Element {
	if v, ok := m[key]; ok {
		return &Element{
			Value: v,
			Path:  path,
		}
	}
	return &Element{
		Path: path,
	}
}

// ToJSON marshal a gmap into json content
func (m GMap) ToJSON() ([]byte, error) {
	return json.Marshal(m)
}

// FromJSON load json content into a GMap
func (m GMap) FromJSON(content []byte) error {
	return json.Unmarshal(content, &m)
}

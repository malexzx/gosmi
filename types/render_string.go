// Code generated by "enumer -type=Render -autotrimprefix -json"; DO NOT EDIT

package types

import (
	"encoding/json"
	"fmt"
)

const (
	_Render_name_0 = "NumericName"
	_Render_name_1 = "Qualified"
	_Render_name_2 = "Format"
	_Render_name_3 = "Printable"
	_Render_name_4 = "Unknown"
	_Render_name_5 = "All"
)

var (
	_Render_index_0 = [...]uint8{0, 7, 11}
	_Render_index_1 = [...]uint8{0, 9}
	_Render_index_2 = [...]uint8{0, 6}
	_Render_index_3 = [...]uint8{0, 9}
	_Render_index_4 = [...]uint8{0, 7}
	_Render_index_5 = [...]uint8{0, 3}
)

func (i Render) String() string {
	switch {
	case 1 <= i && i <= 2:
		i -= 1
		return _Render_name_0[_Render_index_0[i]:_Render_index_0[i+1]]
	case i == 4:
		return _Render_name_1
	case i == 8:
		return _Render_name_2
	case i == 16:
		return _Render_name_3
	case i == 32:
		return _Render_name_4
	case i == 255:
		return _Render_name_5
	default:
		return fmt.Sprintf("Render(%d)", i)
	}
}

var _RenderNameToValue_map = map[string]Render{
	_Render_name_0[0:7]:  1,
	_Render_name_0[7:11]: 2,
	_Render_name_1[0:9]:  4,
	_Render_name_2[0:6]:  8,
	_Render_name_3[0:9]:  16,
	_Render_name_4[0:7]:  32,
	_Render_name_5[0:3]:  255,
}

func RenderFromString(s string) (Render, error) {
	if val, ok := _RenderNameToValue_map[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Render values", s)
}

func RenderAsList() []Render {
	list := make([]Render, len(_RenderNameToValue_map))
	idx := 0
	for _, v := range _RenderNameToValue_map {
		list[idx] = v
		idx++
	}
	return list
}

func RenderAsListString() []string {
	list := make([]string, len(_RenderNameToValue_map))
	idx := 0
	for k := range _RenderNameToValue_map {
		list[idx] = k
		idx++
	}
	return list
}

func RenderIsValid(t Render) bool {
	for _, v := range RenderAsList() {
		if t == v {
			return true
		}
	}
	return false
}

func (i Render) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

func (i *Render) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Render should be a string, got %s", data)
	}

	var err error
	*i, err = RenderFromString(s)
	return err
}

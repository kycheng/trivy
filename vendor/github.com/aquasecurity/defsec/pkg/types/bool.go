package types

import (
	"encoding/json"
)

type BoolValue struct {
	BaseAttribute
	value bool
}

func (b BoolValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"value":    b.value,
		"metadata": b.metadata,
	})
}

func (b *BoolValue) UnmarshalJSON(data []byte) error {
	var keys map[string]interface{}
	if err := json.Unmarshal(data, &keys); err != nil {
		return err
	}
	if keys["value"] != nil {
		b.value = keys["value"].(bool)
	}
	if keys["metadata"] != nil {
		raw, err := json.Marshal(keys["metadata"])
		if err != nil {
			return err
		}
		var m Metadata
		if err := json.Unmarshal(raw, &m); err != nil {
			return err
		}
		b.metadata = m
	}
	return nil
}

func Bool(value bool, metadata Metadata) BoolValue {
	return BoolValue{
		value:         value,
		BaseAttribute: BaseAttribute{metadata: metadata},
	}
}

func BoolDefault(value bool, metadata Metadata) BoolValue {
	b := Bool(value, metadata)
	b.BaseAttribute.metadata.isDefault = true
	return b
}

func BoolUnresolvable(m Metadata) BoolValue {
	b := Bool(false, m)
	b.BaseAttribute.metadata.isUnresolvable = true
	return b
}

func BoolExplicit(value bool, metadata Metadata) BoolValue {
	b := Bool(value, metadata)
	b.BaseAttribute.metadata.isExplicit = true
	return b
}

func (b BoolValue) Value() bool {
	return b.value
}

func (b BoolValue) GetRawValue() interface{} {
	return b.value
}

func (b BoolValue) IsTrue() bool {
	if b.metadata.isUnresolvable {
		return false
	}
	return b.Value()
}

func (b BoolValue) IsFalse() bool {
	if b.metadata.isUnresolvable {
		return false
	}
	return !b.Value()
}

func (s BoolValue) ToRego() interface{} {
	return map[string]interface{}{
		"filepath":  s.metadata.Range().GetFilename(),
		"startline": s.metadata.Range().GetStartLine(),
		"endline":   s.metadata.Range().GetEndLine(),
		"managed":   s.metadata.isManaged,
		"explicit":  s.metadata.isExplicit,
		"value":     s.Value(),
		"fskey":     CreateFSKey(s.metadata.Range().GetFS()),
		"resource":  s.metadata.Reference(),
	}
}

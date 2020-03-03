package libra

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
)

// Libra Canonical Serialization (LCS)
// https://github.com/libra/libra/blob/master/common/lcs/src/lib.rs
type CanonicalSerializer struct {
	reader *bytes.Reader
}

// LCS handler
func NewCanonicalSerializer(blob []byte) *CanonicalSerializer {
	return &CanonicalSerializer{reader: bytes.NewReader(blob)}
}

// LCS read int32
func (cs *CanonicalSerializer) ReadInt32() (int32, error) {
	var val int32
	err := binary.Read(cs.reader, binary.LittleEndian, &val)
	return val, err
}

// LCS read uint64
func (cs *CanonicalSerializer) ReadUint64() (uint64, error) {
	var val uint64
	err := binary.Read(cs.reader, binary.LittleEndian, &val)
	return val, err
}

// LCS read bool
func (cs *CanonicalSerializer) ReadBool() (bool, error) {
	var val bool
	err := binary.Read(cs.reader, binary.LittleEndian, &val)
	return val, err
}

// LCS read string (len + data)
func (cs *CanonicalSerializer) ReadString() (string, error) {
	var len uint32
	err := binary.Read(cs.reader, binary.LittleEndian, &len)
	if err != nil {
		return "", err
	}

	val := make([]byte, len)
	err = binary.Read(cs.reader, binary.LittleEndian, &val)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(val), nil
}

// LCS read string (data[])
func (cs *CanonicalSerializer) ReadXString(len uint32) (string, error) {
	val := make([]byte, len)
	err := binary.Read(cs.reader, binary.LittleEndian, &val)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(val), nil
}

// LCS read map struct
func (cs *CanonicalSerializer) ReadMap() (map[string][]byte, error) {
	var mapEntryCount uint32

	err := binary.Read(cs.reader, binary.LittleEndian, &mapEntryCount)
	if err != nil {
		return nil, err
	}

	m := make(map[string][]byte, mapEntryCount)

	var i uint32
	for i = 0; i < mapEntryCount; i++ {
		var mapKeyLen uint32
		err := binary.Read(cs.reader, binary.LittleEndian, &mapKeyLen)
		if err != nil {
			return nil, err
		}
		mapKey := make([]byte, mapKeyLen)
		err = binary.Read(cs.reader, binary.LittleEndian, &mapKey)
		if err != nil {
			return nil, err
		}

		var mapValLength uint32
		err = binary.Read(cs.reader, binary.LittleEndian, &mapValLength)
		if err != nil {
			return nil, err
		}
		mapVal := make([]byte, mapValLength)
		err = binary.Read(cs.reader, binary.LittleEndian, &mapVal)
		if err != nil {
			return nil, err
		}
		mapKeyAsString := hex.EncodeToString(mapKey)
		m[mapKeyAsString] = mapVal
	}
	return m, nil
}

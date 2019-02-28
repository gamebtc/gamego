
package msg

func (m *GameFrame) Marshal() (dAtA []byte, err error) {
	return m.Data, nil
}

func (m *GameFrame) MarshalTo(dAtA []byte) (int, error) {
	i:= len(m.Data)
	i = copy(dAtA[:i], m.Data)
	return i, nil
}

func (m *GameFrame) Size() (n int) {
	if m == nil {
		return 0
	}
	return len(m.Data)
}

func (m *GameFrame) Unmarshal(dAtA []byte) error {
	m.Data = dAtA
	return nil
}


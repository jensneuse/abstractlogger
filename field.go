package abstractlogger

type Field struct {
	Kind           FieldKind
	Key            string
	StringValue    string
	IntValue       int64
	ByteValue      []byte
}

func (f Field) InterfaceValue () interface{} {
	switch f.Kind {
	case StringField:
		return f.StringValue
	case IntField:
		return f.IntValue
	case BoolField:
		if f.IntValue == 0 {
			return false
		}
		return true
	case ByteStringField:
		return f.ByteValue
	default:
		return nil
	}
}

type FieldKind int

const (
	StringField FieldKind = iota + 1
	IntField
	BoolField
	ByteStringField
)

func String(key, value string) Field {
	return Field{
		Kind:        StringField,
		Key:         key,
		StringValue: value,
	}
}

func Int(key string, value int) Field {
	return Field{
		Kind:     IntField,
		Key:      key,
		IntValue: int64(value),
	}
}

func Bool(key string, value bool) Field {
	var integer int64
	if value {
		integer = 1
	}
	return Field{
		Kind:     BoolField,
		Key:      key,
		IntValue: integer,
	}
}

func ByteString(key string, value []byte) Field {
	return Field{
		Kind:      ByteStringField,
		Key:       key,
		ByteValue: value,
	}
}

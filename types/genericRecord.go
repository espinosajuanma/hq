package types

import S "github.com/espinosajuanma/slingr-go"

type GenericRecord interface {
	S.RecordReference | string | interface{}
}

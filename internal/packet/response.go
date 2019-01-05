package packet

import (
	"encoding/json"
	"fmt"
)

type Res struct {
	Code        string                 `json:"code"`
	State       string                 `json:"state"`
	Message     string                 `json:"message"`
	MessageType string                 `json:"messageType"`
	Data        map[string]interface{} `json:"data"`
}

func (o *Res) AddData(key string, value interface{}) {
	if o.Data == nil {
		o.Data = make(map[string]interface{})
	}

	o.Data[key] = value
}

func (o *Res) GetData(key string) (interface{}, bool) {
	if o.Data == nil {
		return nil, false
	}

	val, ok := o.Data[key]
	return val, ok
}

func (o *Res) DeleteData(key string) {
	delete(o.Data, key)
}

func (o *Res) ToJson() string {
	if o.Data == nil {
		o.Data = make(map[string]interface{})
	}

	result, _ := json.Marshal(o)
	return string(result)
}

func (o *Res) ToString() string {
	return fmt.Sprintf("code : %s, state : %s, message type : %s, message : %s", o.Code, o.State, o.MessageType, o.Message)
}

func Response(code string) *Res {
	return &Res{Code: code}
}

func ResponseWithCustomMessage(code string, messageType string, message string) *Res {
	return &Res{Code: code, MessageType: messageType, Message: message}
}

func ResponseString(code string) string {
	res := Res{Code: code}
	return res.ToJson()
}

func ResponseStringWithCustomMessage(code string, messageType string, message string) string {
	res := Res{Code: code, MessageType: messageType, Message: message}
	return res.ToJson()
}

package messenger

import "encoding/json"

func UnmarshalPassThreadControl(data []byte) (PassThreadControl, error) {
	var r PassThreadControl
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PassThreadControl) marshal() ([]byte, error) {
	return json.Marshal(r)
}

type PassThreadControl struct {
	Recipient     Recipient `json:"recipient"`
	TargetAppID   int64     `json:"target_app_id"`
	Metadata      string    `json:"metadata"`
	NewOwnerAppID int64     `json:"new_owner_app_id"`
}

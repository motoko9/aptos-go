package rpcmodule

type ViewRequest struct {
	Function      string        `json:"function,omitempty"`
	TypeArguments []string      `json:"type_arguments"`
	Arguments     []interface{} `json:"arguments"`
}

package web

var (
	HdrContentType    = "Content-Type"
	ApplicationJSON   = "application/json"
	PhemexAPIEndpoint = "https://api.phemex.com"
)

type weberr struct {
	Code    int    `json:"code"`
	Message string `json:"msg,omitempty"`
}

type Response struct {
	Data interface{} `json:"data,omitempty"`
	weberr
}

type PhemexAPIResponse struct {
	Error string
	Id    int32
}

type PhemexSpotTickerRsp struct {
	PhemexAPIResponse
	Result []tickerData
}

type tickerData struct {
	AskEp      int64  `json:"askEp"`
	BidEp      int64  `json:"bidEp"`
	HighEp     int64  `json:"highEp"`
	IndexEp    int64  `json:"indexEp"`
	LastEp     int64  `json:"lastEp"`
	LowEp      int64  `json:"lowEp"`
	OpenEp     int64  `json:"openEp"`
	Symbol     string `json:"symbol"`
	Timestamp  int64  `json:"timestamp"`
	TurnoverEv int64  `json:"turnoverEv"`
	VolumeEv   int64  `json:"volumeEv"`
}

func NewErrorResponse(e Error) Response {
	return Response{
		Data: nil,
		weberr: weberr{
			Code:    e.Code(),
			Message: e.Message(),
		},
	}
}

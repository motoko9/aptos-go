package hello_blockchain

type MessageChangeEvents struct {
    Counter string `json:"counter"`
    Guid    struct {
        Id struct {
            Addr        string `json:"addr"`
            CreationNum string `json:"creation_num"`
        } `json:"id"`
    } `json:"guid"`
}

type MessageHolder struct {
    Message             string              `json:"message"`
    MessageChangeEvents MessageChangeEvents `json:"message_change_events"`
}

package main

type Event struct {
	UUid string `json:"uuid""`
	Src  string `json:"src" binding:"required"`
	Ip   string `json:"ip" `
}

func (ev *Event) GetUUid() string {
	return ev.UUid
}

func (ev *Event) GetSrc() string {
	return ev.Src
}

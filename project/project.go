package project

type Client struct {
	ID   uint32
	Name string
}

type Project struct {
	ID     uint32
	Name   string
	Client Client
}

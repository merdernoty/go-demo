package main

import "fmt"

type HTTPStatus int

const (
	StatusOk HTTPStatus = 200 + iota
	StatusCreate
	StatauAccepted
	StatusNoContent
)

const (
	StatusBadReq HTTPStatus = 400 + iota
	StatusUnauthorized
	StatusForbidden
	StatusNotFound
)

func (s HTTPStatus) String() string {
	switch s {
	case StatusOk:
		return "OK"
	case StatusCreate:
		return "Created"
	case StatusBadReq:
		return "Bad Req"
	case StatusNotFound:
		return "Not Found"
	default:
		return fmt.Sprintf("Status %d", s)
	}
}


// func main() {
// 	fmt.Printf("Success: %s (%d)\n", StatusOk, StatusOk)
// 	fmt.Printf("Failed: %s (%d)\n", StatusBadReq, StatusBadReq )
// }
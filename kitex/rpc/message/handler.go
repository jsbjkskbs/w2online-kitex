package main

import (
	"context"
	message "work/kitex_gen/message"
)

// MessageServiceImpl implements the last service interface defined in the IDL.
type MessageServiceImpl struct{}

// InsertMessage implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) InsertMessage(ctx context.Context, request *message.InsertMessageRequest) (resp *message.InsertMessageResponse, err error) {
	// TODO: Your code here...
	return
}

// PopMessage implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) PopMessage(ctx context.Context, request *message.PopMessageRequest) (resp *message.PopMessageResponse, err error) {
	// TODO: Your code here...
	return
}

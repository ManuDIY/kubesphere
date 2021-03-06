package spdystream

import (
	"io"
	"net/http"
)

// MirrorStreamHandler mirrors all streams.
func MirrorStreamHandler(stream *Stream) {
	replyErr := stream.SendReply(http.Header{}, false)
	if replyErr != nil {
		return
	}

	go func() {
		io.Copy(stream, stream)
		stream.Close()
	}()
	go func() {
		for {
			header, receiveErr := stream.ReceiveHeader()
			if receiveErr != nil {
				return
			}
			sendErr := stream.SendHeader(header, false)
			if sendErr != nil {
				return
			}
		}
	}()
}

// NoopStreamHandler does nothing when stream connects.
func NoOpStreamHandler(stream *Stream) {
	stream.SendReply(http.Header{}, false)
}

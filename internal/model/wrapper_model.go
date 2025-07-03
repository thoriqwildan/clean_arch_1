package model

type ChannelResponseWrapper struct {
	WebResponse[ChannelResponse]
}

type ChannelsResponseWrapper struct {
	WebResponse[[]ChannelResponse]
}

type ErrorWrapper struct {
	WebResponse[any]
}
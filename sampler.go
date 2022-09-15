package eduid_amapi

import (
	"context"
	"net/http"
)

type samplerService struct {
	client     *Client
	subBaseURL string
}
type SamplerRequest struct {
	Periodicity         float64
	DurationOfExecution float64
	CleanedType         string
}

type SamplerReply struct {
	Status string
}

func (s *samplerService) Get(ctx context.Context, body *SamplerRequest) (*SamplerReply, *http.Response, error) {
	reply := &SamplerReply{}
	resp, err := s.client.call(ctx, http.MethodPost, "/sampler", body, reply)
	if err != nil {
		return nil, resp, err
	}

	return reply, resp, nil
}

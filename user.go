package eduid_amapi

import (
	"context"
	"net/http"
	"net/url"

	"github.com/masv3971/eduid_amapi/amapi_types"
)

type userService struct {
	client     *Client
	subBaseURL string
}

func (s *userService) url(eppn, path string) (string, error) {
	if path == "" {
		return url.JoinPath(s.subBaseURL, eppn)
	}
	return url.JoinPath(s.subBaseURL, eppn, path)
}

type BaseRequest struct {
	Reason string `validate:"required"`
	Source string `validate:"required"`
	Eppn   string `validate:"required"`
}

type UpdateReply struct {
	status bool
	diff   string
}

type UpdateNameRequest struct {
	*BaseRequest
	GivenName   string `json:"given_name"`
	DisplayName string `json:"display_name"`
	Surname     string `json:"surname"`
}

func (s *userService) UpdateName(ctx context.Context, body *UpdateNameRequest) (*UpdateReply, *http.Response, error) {
	reply := &UpdateReply{}
	endpointURL, err := s.url(body.Eppn, "name")
	if err != nil {
		return nil, nil, err
	}
	resp, err := s.client.call(ctx, http.MethodPut, endpointURL, body, reply)
	if err != nil {
		return nil, resp, err
	}
	return reply, resp, nil
}

type UpdateMetaRequest struct {
	*BaseRequest
	Meta *amapi_types.Meta `json:"meta"`
}

func (s *userService) UpdateMeta(ctx context.Context, body *UpdateMetaRequest) (*UpdateReply, *http.Response, error) {
	reply := &UpdateReply{}
	endpointURL, err := s.url(body.Eppn, "meta")
	if err != nil {
		return nil, nil, err
	}
	resp, err := s.client.call(ctx, http.MethodPut, endpointURL, body, reply)
	if err != nil {
		return nil, resp, err
	}
	return reply, resp, nil
}

type UpdateEmailRequest struct {
	*BaseRequest
	MailAddresses *amapi_types.MailAddresses `json:"mail_addresses"`
}

func (s *userService) UpdateEmail(ctx context.Context, body *UpdateEmailRequest) (*UpdateReply, *http.Response, error) {
	reply := &UpdateReply{}
	endpointURL, err := s.url(body.Eppn, "email")
	if err != nil {
		return nil, nil, err
	}
	resp, err := s.client.call(ctx, http.MethodPut, endpointURL, body, reply)
	if err != nil {
		return nil, resp, err
	}
	return reply, resp, nil
}

type UpdateLanguageRequest struct {
	*BaseRequest
	Language string `json:"language"`
}

func (s *userService) UpdateLanguage(ctx context.Context, body *UpdateLanguageRequest) (*UpdateReply, *http.Response, error) {
	reply := &UpdateReply{}
	endpointURL, err := s.url(body.Eppn, "language")
	if err != nil {
		return nil, nil, err
	}
	resp, err := s.client.call(ctx, http.MethodPut, endpointURL, body, reply)
	if err != nil {
		return nil, resp, err
	}
	return reply, resp, nil
}

type UpdatePhoneRequest struct {
	*BaseRequest
	PhoneNumbers amapi_types.PhoneNumbers `json:"phone_numbers"`
}

func (s *userService) UpdatePhone(ctx context.Context, body *UpdatePhoneRequest) (*UpdateReply, *http.Response, error) {
	reply := &UpdateReply{}
	endpointURL, err := s.url(body.Eppn, "phone")
	if err != nil {
		return nil, nil, err
	}
	resp, err := s.client.call(ctx, http.MethodPut, endpointURL, body, reply)
	if err != nil {
		return nil, resp, err
	}
	return reply, resp, nil
}

type TerminateRequest struct {
	*BaseRequest
}

func (s *userService) Terminate(ctx context.Context, body *TerminateRequest) (*UpdateReply, *http.Response, error) {
	reply := &UpdateReply{}
	endpointURL, err := s.url(body.Eppn, "")
	if err != nil {
		return nil, nil, err
	}
	resp, err := s.client.call(ctx, http.MethodDelete, endpointURL, body, reply)
	if err != nil {
		return nil, resp, err
	}
	return reply, resp, nil
}

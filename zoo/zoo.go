package zoo

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/go-resty/resty/v2"
)

type ZooService struct {
	Client *resty.Client
}

func CreateZooApi(baseUrl string) *ZooService {
	client := resty.New()
	client.SetBaseURL(baseUrl)
	client.SetTimeout(5 * time.Second)
	return &ZooService{
		Client: client,
	}
}

func (s *ZooService) PostGateAuth(body PostGateAuthBody) (*PostGateAuthResponse, error) {
	bodyj, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("error when json marshal. error: %w", err)
	}

	start := time.Now()

	resp, err := s.Client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(bodyj).
		SetResult(&PostGateAuthResponse{}).
		Post("webhook/device/gate/auth")

	elapsed := time.Since(start)
	log.Printf("Request gate auth took %s", elapsed)

	if err != nil {
		return nil, fmt.Errorf("error when post gate auth. error: %w", err)
	}

	return resp.Result().(*PostGateAuthResponse), nil
}

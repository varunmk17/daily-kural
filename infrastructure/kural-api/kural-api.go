package kuralapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const KURAL_API_URL = "https://getthirukkural.appspot.com/api/3.0/kural/%d?appid=op7r7rdthtmmu&format=json"

type KuralApiResponse struct {
	Number      int    `json:"number"`
	Paal        string `json:"paal"`
	Iyal        string `json:"iyal"`
	Athigaram   string `json:"athigaram"`
	Line1       string `json:"line1"`
	Line2       string `json:"line2"`
	Urai1       string `json:"urai1"`
	Urai1Author string `json:"urai1Author"`
	Urai2       string `json:"urai2"`
	Urai2Author string `json:"urai2Author"`
	Urai3       string `json:"urai3"`
	Urai3Author string `json:"urai3Author"`
	Translation string `json:"translation"`
	En		    string `json:"en"`
}

func GetKuralByNumber(kuralNumber int) (*KuralApiResponse, error) {
	kuralUrl := fmt.Sprintf(KURAL_API_URL, kuralNumber)

	resp, err := http.Get(kuralUrl)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if string(body) == "BAD REQUEST" {
		return nil, fmt.Errorf("failed to get kural with number %v", kuralNumber)
	}

	kuralApiResponse := &KuralApiResponse{}
	if err := json.Unmarshal(body, kuralApiResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal kural with number %v, %w", kuralNumber, err)
	}

	return kuralApiResponse, err
}

package kuralapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const KURAL_API_URL = "https://api-thirukkural.vercel.app/api?num=%d"

type KuralApiResponse struct {
	Number                int    `json:"number"`
	SectionInTamil        string `json:"sect_tam"`
	ChapterGroupInTamil   string `json:"chapgrp_tam"`
	ChapterInTamil        string `json:"chap_tam"`
	Line1InTamil          string `json:"line1"`
	Line2InTamil          string `json:"line2"`
	ExplanationInTamil    string `json:"tam_exp"`
	SectionInEnglish      string `json:"sect_eng"`
	ChapterGroupInEnglish string `json:"chapgrp_eng"`
	ChapterInEnglish      string `json:"chap_eng"`
	KuralInEnglish        string `json:"eng"`
	ExplanationInEnglish  string `json:"eng_exp"`
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

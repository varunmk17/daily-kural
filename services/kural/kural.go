package kural

import (
	"fmt"
	"strings"

	"local.dev.com/config"
	kuralapi "local.dev.com/infrastructure/kural-api"
	"local.dev.com/utils"
)

type Kural struct {
	Number       int
	Section      string
	ChapterGroup string
	Chapter      string
	Kural        string
	Urai         []Urai
	Language     string
	Translation string // poetic translation (from API: "translation")
	En          string // prose explanation (from API: "en")
	Headers      *KuralNotificationHelper
}

type Urai struct {
	Explanation string
	Author      string
}

type KuralNotificationHelper struct {
	Language          string
	HeaderKural       string
	HeaderExplanation string
}

var SUPPORTED_LANGUAGES = []string{"tamil", "english"}

func GetDailyKural(appSettings *config.Config, number int, language string) (*Kural, error) {
	if number < 1 || number > 1330 {
		return nil, fmt.Errorf("kural number should be between 1 and 1330")
	}

	if !utils.StringArrayContains(SUPPORTED_LANGUAGES, language) {
		return nil, fmt.Errorf("language %s is currently not supported", language)
	}

	kuralApiResponse, err := kuralapi.GetKuralByNumber(number)
	if err != nil {
		return nil, fmt.Errorf("failed to get daily kural, %w", err)
	}

	var kural *Kural

	switch strings.ToLower(language) {
	case "tamil":
		{
			urai := []Urai{}
			if kuralApiResponse.Urai1 != "" && kuralApiResponse.Urai1Author != "" {
				urai = append(urai, Urai{kuralApiResponse.Urai1, kuralApiResponse.Urai1Author})
			}

			if kuralApiResponse.Urai2 != "" && kuralApiResponse.Urai2Author != "" {
				urai = append(urai, Urai{kuralApiResponse.Urai2, kuralApiResponse.Urai2Author})
			}

			if kuralApiResponse.Urai3 != "" && kuralApiResponse.Urai3Author != "" {
				urai = append(urai, Urai{kuralApiResponse.Urai3, kuralApiResponse.Urai3Author})
			}

			kural = &Kural{
				Number:       kuralApiResponse.Number,
				Section:      kuralApiResponse.Paal,
				ChapterGroup: kuralApiResponse.Iyal,
				Chapter:      kuralApiResponse.Athigaram,
				Kural:        fmt.Sprintf("%s %v %s", kuralApiResponse.Line1, "\n <br/>", kuralApiResponse.Line2),
				Urai:         urai,
				Translation:  kuralApiResponse.Translation, 
				En:           kuralApiResponse.En,          
				Language:     language,
				Headers: &KuralNotificationHelper{
					Language:          language,
					HeaderKural:       "திருக்குறள்",
					HeaderExplanation: "பொருள்",
				},
			}
		}
	}

	return kural, nil
}

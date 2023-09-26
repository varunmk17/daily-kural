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
	Explanation  string
	Language     string
	Headers      *KuralNotificationHelper
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
			kural = &Kural{
				Number:       kuralApiResponse.Number,
				Section:      kuralApiResponse.SectionInTamil,
				ChapterGroup: kuralApiResponse.ChapterGroupInTamil,
				Chapter:      kuralApiResponse.ChapterInTamil,
				Kural:        fmt.Sprintf("%s %v %s", kuralApiResponse.Line1InTamil, "\n <br/>", kuralApiResponse.Line2InTamil),
				Explanation:  kuralApiResponse.ExplanationInTamil,
				Language:     language,
				Headers: &KuralNotificationHelper{
					Language:          language,
					HeaderKural:       "திருக்குறள்",
					HeaderExplanation: "பொருள்",
				},
			}
		}
	case "english":
		{
			kural = &Kural{
				Number:       kuralApiResponse.Number,
				Section:      kuralApiResponse.SectionInEnglish,
				ChapterGroup: kuralApiResponse.ChapterGroupInEnglish,
				Chapter:      kuralApiResponse.ChapterInEnglish,
				Kural:        kuralApiResponse.KuralInEnglish,
				Explanation:  kuralApiResponse.ExplanationInEnglish,
				Language:     language,
				Headers: &KuralNotificationHelper{
					Language:          language,
					HeaderKural:       "Thirukkural",
					HeaderExplanation: "Explanation",
				},
			}
		}
	}

	return kural, nil
}

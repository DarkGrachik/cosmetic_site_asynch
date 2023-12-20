package models

type ClinicalTrialRequest struct {
	AccessToken int64  `json:"access_token"`
	ClinicalTrial int `json:"clinical_trial"`
}

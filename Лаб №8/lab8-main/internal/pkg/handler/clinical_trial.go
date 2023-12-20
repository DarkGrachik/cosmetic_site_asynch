package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"lab8/internal/models"
	"math/rand"
	"net/http"
	"time"
)


func (h *Handler) issueClinicalTrial(c *gin.Context) {
	var input models.Request
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println("handler.issueClinicalTrial:", input)

	c.Status(http.StatusOK)

	go func() {
		time.Sleep(3 * time.Second)
		sendClinicalTrialRequest(input)
	}()
}

func sendClinicalTrialRequest(request models.Request) {

	var clinicalTrial = 0
	if rand.Intn(10) % 10 >= 3 {
	 clinicalTrial = rand.Intn(2)
	}

	answer := models.ClinicalTrialRequest{
		AccessToken: 123,
		ClinicalTrial: clinicalTrial,
	}

	client := &http.Client{}

	jsonAnswer, _ := json.Marshal(answer)
	bodyReader := bytes.NewReader(jsonAnswer)

	requestURL := fmt.Sprintf("http://127.0.0.1:8000/api/cosmetics/%d/update_clinical_trial/", request.CosmeticId)

	req, _ := http.NewRequest(http.MethodPut, requestURL, bodyReader)

	req.Header.Set("Content-Type", "application/json")

	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending PUT request:", err)
		return
	}

	defer response.Body.Close()

	fmt.Println("PUT Request Status:", response.Status)
}

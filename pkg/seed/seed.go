package seed

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/jorgepuerta00/accountapi-master/pkg/model"
	"github.com/sirupsen/logrus"
)

type Seed interface {
	Seeding(int) ([]model.Account, error)
}

type Seeding struct {
	baseURL          string
	httpClient       http.Client
	logger           logrus.FieldLogger
	testAccountsData []model.Account
}

func NewSeeding(logger logrus.FieldLogger, baseURL string) *Seeding {
	return &Seeding{
		baseURL:          baseURL,
		logger:           logger,
		httpClient:       http.Client{},
		testAccountsData: []model.Account{},
	}
}

func (c Seeding) Seeding(seeding int) ([]model.Account, error) {

	for i := 1; i <= seeding; i++ {
		newAccount := model.Account{
			Type:           "accounts",
			ID:             uuid.NewString(),
			OrganisationID: uuid.NewString(),
			Version:        0,
			Attributes: model.AccountAttributes{
				Country:                 "GB",
				BaseCurrency:            "GBP",
				AccountNumber:           "41426819",
				BankID:                  "400300",
				BankIDCode:              "GBDSC",
				Bic:                     "NWBKGB22",
				Iban:                    "GB11NWBK40030041426819",
				Name:                    []string{"Samantha Holder"},
				AlternativeNames:        []string{"Sam Holder"},
				AccountClassification:   "Personal",
				JointAccount:            false,
				AccountMatchingOptOut:   false,
				SecondaryIdentification: "A1B2C3D4",
				Switched:                false,
			},
		}

		body := struct {
			Data model.Account `json:"data"`
		}{
			Data: newAccount,
		}

		payload := new(bytes.Buffer)
		json.NewEncoder(payload).Encode(&body)

		_, err := c.httpClient.Post(c.baseURL, "application/json", payload)
		if err != nil {
			c.logger.Error("Seeding.Seed", "error:", err)
			return []model.Account{}, err
		}

		a := append(c.testAccountsData, newAccount)
		c.testAccountsData = a
	}

	return c.testAccountsData, nil
}

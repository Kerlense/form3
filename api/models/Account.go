package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type AccountData struct {
	Attributes     *AccountAttributes `json:"attributes,omitempty"`
	ID             string             `json:"id,omitempty"`
	OrganisationID string             `json:"organisation_id,omitempty"`
	Type           string             `json:"type,omitempty"`
	Version        *int64             `json:"version,omitempty"`
}

type AccountAttributes struct {
	AccountClassification   *string  `json:"account_classification,omitempty"`
	AccountMatchingOptOut   *bool    `json:"account_matching_opt_out,omitempty"`
	AccountNumber           string   `json:"account_number,omitempty"`
	AlternativeNames        []string `json:"alternative_names,omitempty"`
	BankID                  string   `json:"bank_id,omitempty"`
	BankIDCode              string   `json:"bank_id_code,omitempty"`
	BaseCurrency            string   `json:"base_currency,omitempty"`
	Bic                     string   `json:"bic,omitempty"`
	Country                 *string  `json:"country,omitempty"`
	Iban                    string   `json:"iban,omitempty"`
	JointAccount            *bool    `json:"joint_account,omitempty"`
	Name                    []string `json:"name,omitempty"`
	SecondaryIdentification string   `json:"secondary_identification,omitempty"`
	Status                  *string  `json:"status,omitempty"`
	Switched                *bool    `json:"switched,omitempty"`
}

func (a *AccountData) CreateAccount(db *gorm.DB) (*AccountData, error) {
	var err error
	err = db.Debug().Model(&AccountData{}).Create(&a).Error
	if err != nil {
		return &AccountData{}, err
	}

	return a, nil
}

func (a *AccountData) GetAccount(db *gorm.DB, pid uint64) (*AccountData, error) {
	var err error
	err = db.Debug().Model(&AccountData{}).Where("id = ?", pid).Take(&a).Error
	if err != nil {
		return &AccountData{}, err
	}
	return a, nil
}

func (a *AccountData) DeleteAccount(db *gorm.DB, pid uint64, uid uint32) (int64, error) {

	db = db.Debug().Model(&AccountData{}).Where("id = ? and author_id = ?", pid, uid).Take(&AccountData{}).Delete(&AccountData{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Account not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

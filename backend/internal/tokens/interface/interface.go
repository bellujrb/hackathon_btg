package inter

import "time"

type TokenCreateInput struct {
	Traits       *map[string]string `json:"traits,omitempty"`
	Name         string             `json:"name"`
	AssetValue   float64            `json:"assetValue"`
	ExpireDate   string             `json:"expireDate"`
	Description  string             `json:"description"`
	MaxSupply    int                `json:"maxSupply"`
	ImageURL     *string            `json:"imageUrl"`
	ContractId   string             `json:"contractId"`
	AssetID      string             `json:"assetId"`      // Adicionado AssetID
	NominalValue float64            `json:"nominalValue"` // Adicionado NominalValue
	Guarantees   string             `json:"guarantees"`   // Adicionado Guarantees
}

type TokenCreateOutput struct {
	ID            string            `json:"id"`
	URINumber     int               `json:"uriNumber"`
	ContractID    string            `json:"contractId"`
	Name          string            `json:"name"`
	Description   string            `json:"description"`
	MaxSupply     int               `json:"maxSupply"`
	CurrentSupply int               `json:"currentSupply"`
	Traits        map[string]string `json:"traits,omitempty"`
	ImageURL      string            `json:"imageUrl"`
	MetadataURL   string            `json:"metadataUrl"`
	CreatedAt     string            `json:"createdAt"`
	UpdatedAt     string            `json:"updatedAt"`
}

type GetToken struct {
	ContractID string `json:"contractId"`
	UriNumber  string `json:"uriNumber"`
}

type AllTokensOutput struct {
	Tokens []TokenCreateOutput `json:"tokens"`
}
type GainTokenController struct {
	WalletID string `json:"walletId"`
}

type Gimmetoken struct {
	ContractID string `json:"contractID"`
	WalletID   string `json:"walletID"`
	Quantity   int    `json:"quantity"`
	URINumber  int    `json:"uriNumber"`
}

type TransactionRes struct {
	ID              string            `json:"id"`
	WalletID        string            `json:"walletId"`
	Status          string            `json:"status"`
	TransactionHash string            `json:"transactionHash"`
	Result          map[string]string `json:"result"`
	Request         struct {
		ContractID string `json:"contractId"`
		Quantity   int    `json:"quantity"`
	} `json:"request"`
	Type        string `json:"type"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
	CompletedAt string `json:"completedAt"`
}

type PriceInfo struct {
	Price        string
	PriceConvert string
	Token        string
	TokenConvert string
	Date         time.Time
}

type SaveToken struct {
	UriNumber    uint64  `gorm:"column:uriNumber" json:"uriNumber"`
	ContractID   string  `gorm:"column:contract_id" json:"contract_id"`
	AssetID      string  `gorm:"column:asset_id" json:"asset_id"`
	ExpireDate   string  `gorm:"column:expire_date" json:"expireDate"`
	NominalValue float64 `gorm:"column:nominal_value" json:"nominal_value"`
	Guarantees   string  `gorm:"column:guarantees" json:"guarantees"`
}

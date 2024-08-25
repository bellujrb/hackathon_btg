package db

import "time"

type User struct {
	ID        uint      `gorm:"column:id;primaryKey" json:"id"`
	Email     string    `gorm:"column:email;unique;not null" json:"email"`
	Password  string    `gorm:"column:password;not null" json:"password"`
	Wallet    string    `gorm:"column:wallet" json:"wallet"`
	WalletId  string    `gorm:"column:walletId" json:"walletId"`
	ProjectId string    `gorm:"column:projectId" json:"projectId"`
	Login     Login     `gorm:"foreignKey:LoginID" json:"login"`
	LoginID   uint      `gorm:"column:Login_idLogin;not null" json:"Login_idLogin"`
	CreateAt  time.Time `gorm:"column:create_at;not null" json:"create_at"`
	UpdateAt  time.Time `gorm:"column:update_at;not null" json:"update_at"`
}

type Login struct {
	ID       uint      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Email    string    `gorm:"column:email" json:"email"`
	IsLogged bool      `gorm:"column:isLogged;not null" json:"isLogged"`
	CreateAt time.Time `gorm:"column:create_at;not null" json:"create_at"`
	UpdateAt time.Time `gorm:"column:update_at;not null" json:"update_at"`
}

type Token struct {
	ID           uint      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UriNumber    uint64    `gorm:"column:uriNumber" json:"uriNumber"`
	ContractID   string    `gorm:"column:contract_id" json:"contract_id"`
	AssetID      string    `gorm:"column:asset_id" json:"asset_id"`
	ExpireDate   string    `gorm:"column:expireDate" json:"expireDate"`
	NominalValue float64   `gorm:"column:nominal_value" json:"nominal_value"`
	Guarantees   string    `gorm:"column:guarantees" json:"guarantees"`
	Quantity     int       `gorm:"column:quantity" json:"quantity"`
	InvestorID   string    `gorm:"column:investor_id" json:"investor_id"`
	CreatedAt    time.Time `gorm:"column:create_at;not null" json:"create_at"`
	UpdatedAt    time.Time `gorm:"column:update_at;not null" json:"update_at"`
}

type Receivable struct {
	ID                   uint    `gorm:"primaryKey" json:"id"`
	EmpresaNome          string  `json:"empresa_nome"`
	EmpresaCNPJ          string  `json:"empresa_cnpj"`
	EmpresaEndereco      string  `json:"empresa_endereco"`
	EmpresaTelefone      string  `json:"empresa_telefone"`
	EmpresaEmail         string  `json:"empresa_email"`
	DevedorNome          string  `json:"devedor_nome"`
	DevedorCNPJ          string  `json:"devedor_cnpj"`
	Valor                float64 `json:"valor"`
	DataEmissaoTitulo    string  `json:"data_emissao_titulo"`
	DataVencimentoTitulo string  `json:"data_vencimento_titulo"`
	Garantias            string  `json:"garantias"`
	DescontoAntecipacao  float64 `json:"desconto_antecipacao"`
	Banco                string  `json:"banco"`
	Agencia              string  `json:"agencia"`
	ContaCorrente        string  `json:"conta_corrente"`
	ChavePix             string  `json:"chave_pix"`
}

type Contract struct {
	ContractID  string    `gorm:"primaryKey"`
	Paid        bool      `gorm:"default:false"`
	PaymentDate time.Time `gorm:"default:null"`
	TotalTokens int
}

type Distribution struct {
	ID         uint   `gorm:"primaryKey"`
	InvestorID string `gorm:"index"`
	Amount     float64
	ContractID string `gorm:"index"`
}

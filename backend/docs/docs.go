// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/all-contract": {
            "get": {
                "description": "Get All contracts",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Contract"
                ],
                "summary": "Get All contracts",
                "responses": {
                    "200": {
                        "description": "New User Created successfully",
                        "schema": {
                            "$ref": "#/definitions/inter.UserOutputController"
                        }
                    },
                    "500": {
                        "description": "Unable to store data in database",
                        "schema": {
                            "$ref": "#/definitions/erros.InternalServerError"
                        }
                    }
                }
            }
        },
        "/api/all-token": {
            "get": {
                "description": "Create a new user in db",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Token"
                ],
                "summary": "Get all Tokens",
                "parameters": [
                    {
                        "type": "string",
                        "description": "contract id",
                        "name": "ContractId",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "New User Created successfully",
                        "schema": {
                            "$ref": "#/definitions/inter.AllTokensOutput"
                        }
                    },
                    "500": {
                        "description": "Unable to store data in database",
                        "schema": {
                            "$ref": "#/definitions/erros.InternalServerError"
                        }
                    }
                }
            }
        },
        "/api/contract": {
            "get": {
                "description": "Pull a contract from blockchain",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Contract"
                ],
                "summary": "Pull a contract",
                "parameters": [
                    {
                        "type": "string",
                        "description": "contract Id",
                        "name": "ContractId",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "New User Created successfully",
                        "schema": {
                            "$ref": "#/definitions/inter.UserOutputController"
                        }
                    },
                    "500": {
                        "description": "Unable to store data in database",
                        "schema": {
                            "$ref": "#/definitions/erros.InternalServerError"
                        }
                    }
                }
            }
        },
        "/api/create-contract": {
            "post": {
                "description": "Create a new contract",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Contract"
                ],
                "summary": "Create Contract",
                "parameters": [
                    {
                        "description": "Data for make a new contract",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/inter.ContractController"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "New Contract Created successfully",
                        "schema": {
                            "$ref": "#/definitions/inter.UserOutputController"
                        }
                    },
                    "500": {
                        "description": "Error for make a new contract",
                        "schema": {
                            "$ref": "#/definitions/erros.InternalServerError"
                        }
                    }
                }
            }
        },
        "/api/create-token": {
            "post": {
                "description": "Create a new user in db",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Token"
                ],
                "summary": "Make Tokens",
                "parameters": [
                    {
                        "description": "Data for make a new token",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/inter.TokenCreateInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "New Token Created successfully",
                        "schema": {
                            "$ref": "#/definitions/inter.TokenCreateOutput"
                        }
                    },
                    "500": {
                        "description": "Unable to created a new token",
                        "schema": {
                            "$ref": "#/definitions/erros.InternalServerError"
                        }
                    }
                }
            }
        },
        "/api/deploy-contract": {
            "post": {
                "description": "Deploy new contract in the blockchain",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Contract"
                ],
                "summary": "Deploy new contract",
                "parameters": [
                    {
                        "description": "ContractId for deploy a new contract",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/inter.DeployController"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Deploy its a sucessfully",
                        "schema": {
                            "$ref": "#/definitions/inter.ContractDeploymentResponse"
                        }
                    },
                    "500": {
                        "description": "Unable to make deploy",
                        "schema": {
                            "$ref": "#/definitions/erros.InternalServerError"
                        }
                    }
                }
            }
        },
        "/api/get-token": {
            "get": {
                "description": "Create a new user in db",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Token"
                ],
                "summary": "Get specify token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "E-mail from user",
                        "name": "ContractId",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Its a uri content in tokens",
                        "name": "UriToken",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "New User Created successfully",
                        "schema": {
                            "$ref": "#/definitions/inter.UserOutputController"
                        }
                    },
                    "500": {
                        "description": "Unable to store data in database",
                        "schema": {
                            "$ref": "#/definitions/erros.InternalServerError"
                        }
                    }
                }
            }
        },
        "/api/loggout": {
            "put": {
                "description": "pull user and wallet for db",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Loggout User",
                "parameters": [
                    {
                        "type": "string",
                        "description": "E-mail from user",
                        "name": "Email",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "New User Created successfully",
                        "schema": {
                            "$ref": "#/definitions/inter.UserOutputController"
                        }
                    },
                    "500": {
                        "description": "Unable to store data in database",
                        "schema": {
                            "$ref": "#/definitions/erros.InternalServerError"
                        }
                    }
                }
            }
        },
        "/api/user": {
            "get": {
                "description": "pull user and wallet for db",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "E-mail from user",
                        "name": "Email",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "New User Created successfully",
                        "schema": {
                            "$ref": "#/definitions/inter.UserOutputController"
                        }
                    },
                    "500": {
                        "description": "Unable to store data in database",
                        "schema": {
                            "$ref": "#/definitions/erros.InternalServerError"
                        }
                    }
                }
            }
        },
        "/create-user": {
            "post": {
                "description": "Create a new user in db",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Create user",
                "parameters": [
                    {
                        "description": "Data for make a new user",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/inter.UserController"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "New User Created successfully",
                        "schema": {
                            "$ref": "#/definitions/inter.UserOutputController"
                        }
                    },
                    "500": {
                        "description": "Unable to store data in database",
                        "schema": {
                            "$ref": "#/definitions/erros.InternalServerError"
                        }
                    }
                }
            }
        },
        "/receivables": {
            "get": {
                "description": "Returns all registered receivables",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Receivable"
                ],
                "summary": "Get all receivables",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/db.Receivable"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/receivable.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Creates a new receivable from the provided data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Receivable"
                ],
                "summary": "Creates a new receivable",
                "parameters": [
                    {
                        "description": "Receivable data",
                        "name": "receivable",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/db.Receivable"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/db.Receivable"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/receivable.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/receivables/{id}": {
            "get": {
                "description": "Fetches a receivable by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Receivable"
                ],
                "summary": "Get a receivable by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Receivable ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/db.Receivable"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/receivable.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/receivables/{id}/early_payment": {
            "get": {
                "description": "Returns options for anticipating the receivable's payment",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Receivable"
                ],
                "summary": "Fetches early payment options for a receivable",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Receivable ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/db.Receivable"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/receivable.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/receivables/{id}/payment_date": {
            "put": {
                "description": "Changes the payment date of the given receivable",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Receivable"
                ],
                "summary": "Updates the payment date of a receivable",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Receivable ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "New Payment Date",
                        "name": "payment_date",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/db.Receivable"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/receivable.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "base.BaseReq": {
            "type": "object",
            "properties": {
                "LogData": {
                    "$ref": "#/definitions/base.LogData"
                },
                "ResponseData": {
                    "$ref": "#/definitions/base.ResponseData"
                }
            }
        },
        "base.LogData": {
            "type": "object",
            "properties": {
                "ExecutionTime": {
                    "type": "string"
                },
                "HttpStatusCode": {
                    "type": "integer"
                },
                "IP": {
                    "type": "string"
                },
                "Route": {
                    "type": "string"
                }
            }
        },
        "base.ResponseData": {
            "type": "object",
            "properties": {
                "Response": {}
            }
        },
        "db.Receivable": {
            "type": "object",
            "properties": {
                "agencia": {
                    "type": "string"
                },
                "banco": {
                    "type": "string"
                },
                "chave_pix": {
                    "type": "string"
                },
                "conta_corrente": {
                    "type": "string"
                },
                "data_emissao_titulo": {
                    "type": "string"
                },
                "data_vencimento_titulo": {
                    "type": "string"
                },
                "desconto_antecipacao": {
                    "type": "number"
                },
                "devedor_cnpj": {
                    "type": "string"
                },
                "devedor_nome": {
                    "type": "string"
                },
                "empresa_cnpj": {
                    "type": "string"
                },
                "empresa_email": {
                    "type": "string"
                },
                "empresa_endereco": {
                    "type": "string"
                },
                "empresa_nome": {
                    "type": "string"
                },
                "empresa_telefone": {
                    "type": "string"
                },
                "garantias": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "valor": {
                    "type": "number"
                }
            }
        },
        "erros.InternalServerError": {
            "type": "object",
            "properties": {
                "BaseReq": {
                    "$ref": "#/definitions/base.BaseReq"
                },
                "InternalServerError": {
                    "$ref": "#/definitions/erros.Message"
                },
                "Response": {}
            }
        },
        "erros.Message": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "inter.AllTokensOutput": {
            "type": "object",
            "properties": {
                "tokens": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/inter.TokenCreateOutput"
                    }
                }
            }
        },
        "inter.ContractController": {
            "type": "object",
            "properties": {
                "contractType": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "symbol": {
                    "type": "string"
                }
            }
        },
        "inter.ContractDeploymentResponse": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "blockchainName": {
                    "type": "string"
                },
                "blockscanUrl": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "deployedAt": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "symbol": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "inter.DeployController": {
            "type": "object",
            "properties": {
                "contractId": {
                    "type": "string"
                }
            }
        },
        "inter.TokenCreateInput": {
            "type": "object",
            "properties": {
                "assetId": {
                    "description": "Adicionado AssetID",
                    "type": "string"
                },
                "assetValue": {
                    "type": "number"
                },
                "contractId": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "expireDate": {
                    "type": "string"
                },
                "guarantees": {
                    "description": "Adicionado Guarantees",
                    "type": "string"
                },
                "imageUrl": {
                    "type": "string"
                },
                "maxSupply": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "nominalValue": {
                    "description": "Adicionado NominalValue",
                    "type": "number"
                },
                "traits": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                }
            }
        },
        "inter.TokenCreateOutput": {
            "type": "object",
            "properties": {
                "contractId": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "currentSupply": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "imageUrl": {
                    "type": "string"
                },
                "maxSupply": {
                    "type": "integer"
                },
                "metadataUrl": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "traits": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "updatedAt": {
                    "type": "string"
                },
                "uriNumber": {
                    "type": "integer"
                }
            }
        },
        "inter.UserController": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "inter.UserOutputController": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "projectId": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "receivable.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    },
    "externalDocs": {
        "description": "OpenAPI",
        "url": "https://swagger.io/resources/open-api/"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "3.145.127.73:8080",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Brasa",
	Description:      "This is a server for app.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

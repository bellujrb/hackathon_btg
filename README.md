## Hackathon BTG (BrasaHacks)

<div align="center">
    <img src="https://cdn.discordapp.com/attachments/1235359156743962746/1237891322015121438/image.png?ex=663d4ba2&is=663bfa22&hm=24435224343d05b2b227e2e6fc3c9f9ae639ea5b4b44b3d323970f805dc777ae&" alt="Logo">
</div>

---

# LUMQI: Powered by BTG Pactual

> _TEAM LUMQI: Hackathon BTG (BrasaHacks)

![Build Status](https://img.shields.io/badge/Build-Passing-brightgreen)
![Platform](https://img.shields.io/badge/Platform-Web-blue)
![License](https://img.shields.io/badge/License-MIT-green)

---

> _Swagger API Documentation: 

> _Documentation PROJECT:
> 
> _Prototyping: https://www.figma.com/design/imspxcdxenZWKVmady3MPz/BrasaHacks---UI%2FUX?node-id=37-770&t=029wiMmdBkYynDVu-1

> _Video Demo Application and API:

---

## 🌐 Introduction

A LUMQI tokenizes receivables using the Polygon blockchain and integrates 100% with the Lumx API, which manages all blockchain operations in the backend. Our goal is to democratize access to the capital markets by simplifying investments in tokenized receivables.


## Application Programming Interface Lumx

We use 100% of the Lumx API to perform all blockchain-related functions within our application.

## Link Contract

https://amoy.polygonscan.com/address/0xf37C923B98d931fff26E944F65113a1f8F66BBF7

## Tokenization Project

```json
{
  "id": "36defbb5-ec8f-4f30-88c5-21cda524218a",
  "name": "BTG Hacks",
  "blockchain": {
    "name": "Polygon",
    "decimalChainId": 80002
  },
  "createdAt": "2024-08-25T23:01:45.268Z",
  "updatedAt": "2024-08-25T23:01:45.268Z",
  "apiKey": "secret"
}
```

---

## Contract Information

```json
{
  "id": "2514cce9-3ad3-475c-9513-070dedf30ff6",
  "address": "0xf37C923B98d931fff26E944F65113a1f8F66BBF7",
  "type": "fungible",
  "name": "BTG Credit 1",
  "symbol": "BTG1",
  "description": "Recebivel",
  "blockchainName": "Polygon",
  "blockExplorerUrl": "https://amoy.polygonscan.com/address/0xf37C923B98d931fff26E944F65113a1f8F66BBF7",
  "baseUri": "https://storage.googleapis.com/lumx-nft-data-tmp/36defbb5-ec8f-4f30-88c5-21cda524218a/2514cce9-3ad3-475c-9513-070dedf30ff6/",
  "createdAt": "2024-08-25T23:02:28.614Z",
  "deployedAt": "2024-08-25T23:06:39.976Z",
  "updatedAt": "2024-08-25T23:06:39.978Z"
}
```

## Mocks in Web

Due to time constraints, some features are mocked on the website, but the backend is fully functional and you can test it yourself via Swagger.

---

## 🛠 Installation (Mobile)

1. **Pre-requisites**
    - Make sure you have Dart and Flutter installed on your machine.

2. **Clone the Repository**

    ```bash
    git clone https://github.com/bellujrb/hackathon_btg/frontend
    ```

3. **Install Dependencies**

    ```bash
    flutter pub get
    ```

4. **Run the App**

    ```bash
    flutter run
    ```

---

## 🛠 Installation (Backend)

1. **Pre-requisites**
    - Make sure you have MSI installed on your machine.

    ```bash
    https://go.dev/doc/install
    ```

2. **Clone the Repository**

    ```bash
    git clone https://github.com/bellujrb/hackathon_btg/backend
    ```

3. **Install Dependencies**

    ```bash
    go get
    ```

4. **Run the App**

    ```bash
    go run main
    ```

---

## 📂 Project File Tree
    
```
hackathon_btg
├── frontend
│   └── ...
├── back-end
│   └── ...
├── README.MD
│   └── ...
```
---

#### `hackathon_btg`

- `front-end`
    - Frontend Application
- `back-end`
    - Back-end Application using API Lumx for Blockchain.
- `README.md`
    - Documentation Project

---

## 🛠 Tech Stack Mobile

### Design Patterns (Mobile)
- Singleton

### External Packages (Mobile)
- Flutter Modular
- HTTP
- Logger

### Architecture (Mobile)
- Modular

---

## 🛠 Tech Stack (Backend)

### API
- Lumx 

### External Packages (Backend)
- Go
- Gin Tonic
- Gin Swagger and Swaggo

### Architecture (Backend)
- MVC

---

## 🙏 Acknowledgments

Special thanks to BTG and Brasa for this ambitious opportunity.

---

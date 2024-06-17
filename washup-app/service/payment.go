package service

import (
	"github.com/fabianMendez/mercadopago"
	"fmt"
	"bytes"
    "encoding/json"
    "net/http"
)

type PaymentParams struct {
	Title string `json:"title"`
	Price float64 `json:"price"`
}

type PreferenceHandler struct {
	mpClient *mercadopago.Client
}

func NewPreferenceHandler() *PreferenceHandler {
	client := mercadopago.NewClient("https://api.mercadopago.com/v1", "TEST-76053192-d831-4b11-8281-3270e7a283f1", "TEST-3705400319827255-110219-cc509cb730e4aa70fb23f32cae71acee-1532920685")
	return &PreferenceHandler{mpClient: &client}
}

func (h *PreferenceHandler) CreatePreference(params PaymentParams) (string, error) {
    preferenceData := map[string]interface{}{
        "items": []map[string]interface{}{
            {
                "title":       params.Title,
                "quantity":    1,
                "currency_id": "UYU",
                "unit_price":  params.Price,
            },
        },
        "back_urls": map[string]string{
            "success": "http://localhost:3000/success",
            "failure": "http://localhost:3000/failure",
            "pending": "http://localhost:3000/pending",
        },
        "auto_return": "approved",
    }

    data, err := json.Marshal(preferenceData)
    if err != nil {
        return "", err
    }

    // Utiliza la URL y el access token directamente aquí
    url := "https://api.mercadopago.com/checkout/preferences"
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
    if err != nil {
        return "", err
    }
    
    // Agrega el access token al header de la solicitud
    req.Header.Add("Authorization", "Bearer " + "TEST-3705400319827255-110219-cc509cb730e4aa70fb23f32cae71acee-1532920685")
    req.Header.Add("Content-Type", "application/json")
    
    // Realiza la solicitud HTTP
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()
    
    if resp.StatusCode != http.StatusCreated {
        return "", fmt.Errorf("failed to create preference, status code: %d", resp.StatusCode)
    }

    var result map[string]interface{}
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return "", err
    }

    initPoint, ok := result["init_point"].(string)
    if !ok {
        return "", fmt.Errorf("init_point not found in MercadoPago response")
    }

    return initPoint, nil
}
package qr

import (
    "fmt"
    "github.com/skip2/go-qrcode" 
)

type Service struct{}

func NewService() *Service {
    return &Service{}
}

func (s *Service) GenerateQR(text string) ([]byte, error) {
    png, err := qrcode.Encode(text, qrcode.Medium, 256)
    if err != nil {
        return nil, fmt.Errorf("failed to generate QR code: %w", err)
    }
    return png, nil
}
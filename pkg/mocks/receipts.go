package mocks

import (
	"fetch_receipt_processor/pkg/models"
)

// ReceiptService is a mock implementation of the service layer for receipt processing.
type ReceiptService struct {
	CalculatePointsFunc func(receipt models.Receipt) int
	ProcessReceiptFunc  func(receipt models.Receipt) (string, error)
	GetPointsFunc       func(id string) (int, error)
}

// CalculatePoints calculates the points for a given receipt.
func (s *ReceiptService) CalculatePoints(receipt models.Receipt) int {
	return s.CalculatePointsFunc(receipt)
}

// ProcessReceipt processes a receipt and returns the ID of the receipt.
func (s *ReceiptService) ProcessReceipt(receipt models.Receipt) (string, error) {
	return s.ProcessReceiptFunc(receipt)
}

// GetPoints retrieves points associated with a receipt ID.
func (s *ReceiptService) GetPoints(id string) (int, error) {
	return s.GetPointsFunc(id)
}

// NewReceiptService creates a new ReceiptService instance.
func NewReceiptService() *ReceiptService {
	return &ReceiptService{}
}

// Rule is a mock implementation of the Rule interface.
type Rule struct {
	CalculateFunc func(receipt models.Receipt) int
}



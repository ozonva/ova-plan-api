package kafka

import (
	"encoding/json"
	"fmt"
	"github.com/ozonva/ova-plan-api/internal/models"
)

type planMessage struct {
	encodedMessage []byte
}

func (p *planMessage) GetEncoded() []byte {
	return p.encodedMessage
}

func NewPlanMessage(plan *models.Plan) (Message, error) {
	encoded, err := json.Marshal(plan)
	if err != nil {
		return nil, err
	}
	return &planMessage{encodedMessage: encoded}, nil
}

func NewRemovePlanMessage(planId uint64) (Message, error) {
	encoded := fmt.Sprintf(`{"id":"%d"}`, planId)
	return &planMessage{encodedMessage: []byte(encoded)}, nil
}

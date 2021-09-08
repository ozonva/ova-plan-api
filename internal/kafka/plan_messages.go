package kafka

import "github.com/ozonva/ova-plan-api/internal/models"

const (
	CreatePlanTopic Topic = "create_plan"
	UpdatePlanTopic Topic = "update_plan"
	RemovePlanTopic Topic = "remove_plan"
)

type planMessages struct {
	messages []Message
	topic    Topic
}

func (p *planMessages) GetTopic() Topic {
	return p.topic
}

func (p *planMessages) GetMessages() []Message {
	return p.messages
}

func NewCreatePlanMessages(plans []models.Plan) (Messages, error) {
	return generatePlanMsgs(plans, CreatePlanTopic)
}

func NewUpdatePlanMessages(plans []models.Plan) (Messages, error) {
	return generatePlanMsgs(plans, UpdatePlanTopic)
}

func NewRemovePlanMessages(planIds []uint64) (Messages, error) {
	msgs := make([]Message, 0, len(planIds))
	for _, id := range planIds {
		removeMsg, err := NewRemovePlanMessage(id)
		if err != nil {
			return nil, err
		}
		msgs = append(msgs, removeMsg)
	}
	return &planMessages{messages: msgs, topic: RemovePlanTopic}, nil
}

func generatePlanMsgs(plans []models.Plan, topic Topic) (Messages, error) {
	planMsgs := make([]Message, 0, len(plans))
	for _, plan := range plans {
		planMsg, err := NewPlanMessage(&plan)
		if err != nil {
			return nil, err
		}
		planMsgs = append(planMsgs, planMsg)
	}

	return &planMessages{messages: planMsgs, topic: topic}, nil
}

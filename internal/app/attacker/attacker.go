package attacker

import "vegeta-server/internal/app/server/models"

// IAttacker is the attacker interface
type IAttacker interface {
	Submit(models.Attack) models.AttackResponse
}

type attacker struct {
	// scheduler schedules an attack to run
	scheduler IScheduler
}

// NewAttacker returns an implementation of the
// IAttacker interface
func NewAttacker(scheduler IScheduler) *attacker {
	return &attacker{
	scheduler,
	}
}

func (a *attacker) Submit(attack models.Attack) models.AttackResponse {
	resp := a.scheduler.Schedule(attack)
	return resp
}

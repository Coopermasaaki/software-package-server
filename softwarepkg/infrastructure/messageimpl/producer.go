package messageimpl

import (
	"github.com/opensourceways/software-package-server/common/infrastructure/kafka"
	"github.com/opensourceways/software-package-server/softwarepkg/domain/message"
)

var producerInstance *producer

func Producer() *producer {
	return producerInstance
}

type producer struct {
	topics Topics
}

func (p *producer) NotifyPkgApplied(e message.EventMessage) error {
	return send(p.topics.NewSoftwarePkg, e)
}

func (p *producer) NotifyPkgToRerunCI(e message.EventMessage) error {
	return send(p.topics.NewSoftwarePkg, e)
}

func (p *producer) NotifyPkgApproved(e message.EventMessage) error {
	return send(p.topics.ApprovedSoftwarePkg, e)
}

func (p *producer) NotifyPkgRejected(e message.EventMessage) error {
	return send(p.topics.RejectedSoftwarePkg, e)
}

func (p *producer) NotifyPkgAbandoned(e message.EventMessage) error {
	return send(p.topics.AbandonedSoftwarePkg, e)
}

func (p *producer) NotifyPkgAlreadyExisted(e message.EventMessage) error {
	return send(p.topics.AlreadyExistedSoftwarePkg, e)
}

func send(topic string, v message.EventMessage) error {
	body, err := v.Message()
	if err != nil {
		return err
	}

	return kafka.Publish(topic, body)
}

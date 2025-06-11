package main

func (a *analytics) handleEmailBounce(em email) error {
	err := em.recipient.updateStatus(em.status)
	if err != nil {
		return err
	}
	err = a.track(em.status)
	if err != nil {
		return err
	}
	return nil
}

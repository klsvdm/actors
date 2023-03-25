package actor

import "context"

type actorMock struct {
	id *ID
}

func (a *actorMock) ID() *ID {
	return a.id
}

func (a *actorMock) Invoke(_ *ID, _ any) {
}

func (a *actorMock) Shutdown(_ context.Context) error {
	return nil
}
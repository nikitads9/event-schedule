package sender

import (
	"context"
	"sync"
)

type App struct {
	serviceProvider *serviceProvider
	pathConfig      string
}

// NewApp ...
func NewApp(ctx context.Context, pathConfig string) (*App, error) {
	a := &App{
		pathConfig: pathConfig,
	}
	err := a.initDeps(ctx)

	return a, err
}

func (a *App) initDeps(ctx context.Context) error {

	inits := []func(context.Context) error{
		a.initServiceProvider,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider(a.pathConfig)

	return nil
}

// Run ...
func (a *App) Run(ctx context.Context) error {
	defer func() {
		a.serviceProvider.rabbitConsumer.Close()
	}()

	wg := &sync.WaitGroup{}
	wg.Add(1)
	err := a.runSenderService(ctx, wg)
	if err != nil {
		return err
	}
	wg.Wait()
	a.serviceProvider.GetLogger().Info("sender service stopped")
	return nil
}

func (a *App) runSenderService(ctx context.Context, wg *sync.WaitGroup) error {
	go func() {
		defer wg.Done()

		a.serviceProvider.GetLogger().Info("attempting to run sender service")
		a.serviceProvider.GetSenderService(ctx).Run(ctx)

	}()

	return nil
}

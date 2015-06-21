package shadow

import (
	"github.com/dropbox/godropbox/errors"
)

type ContextItem interface {
	GetName() string
	Init(*Application) error
}

type ContextItemRunner interface {
	Run() error
}

type Resource interface {
	ContextItem
}

type Service interface {
	ContextItem
}

type Application struct {
	resources []Resource
	services  []Service

	Version string
	Build   string
}

func NewApplication(resources []Resource, services []Service, version string, build string) (*Application, error) {
	application := &Application{
		resources: []Resource{},
		services:  []Service{},
		Version:   version,
		Build:     build,
	}

	for i := range resources {
		if err := application.RegisterResource(resources[i]); err != nil {
			return nil, err
		}
	}

	for i := range services {
		if err := application.RegisterService(services[i]); err != nil {
			return nil, err
		}
	}

	return application, nil
}

func (a *Application) Run() error {
	// Resources
	resources := a.GetResources()

	for i := range resources {
		if err := resources[i].Init(a); err != nil {
			return err
		}
	}

	for i := range resources {
		if runner, ok := resources[i].(ContextItemRunner); ok {
			runner.Run()
		}
	}

	// Services
	services := a.GetServices()

	for i := range services {
		if err := services[i].Init(a); err != nil {
			return err
		}
	}

	for i := range services {
		if runner, ok := services[i].(ContextItemRunner); ok {
			runner.Run()
		}
	}

	select {}

	return nil
}

func (a *Application) GetResource(name string) (Resource, error) {
	for i := range a.resources {
		if a.resources[i].GetName() == name {
			return a.resources[i], nil
		}
	}

	return nil, errors.Newf("Resource \"%s\" not found", name)
}

func (a *Application) GetResources() []Resource {
	return a.resources
}

func (a *Application) HasResource(name string) bool {
	_, err := a.GetResource(name)
	return err == nil
}

func (a *Application) RegisterResource(resource Resource) error {
	if _, err := a.GetResource(resource.GetName()); err == nil {
		return errors.Newf("Resource \"%s\" already exists", resource.GetName())
	}

	a.resources = append(a.resources, resource)
	return nil
}

func (a *Application) GetService(name string) (Service, error) {
	for i := range a.services {
		if a.services[i].GetName() == name {
			return a.services[i], nil
		}
	}

	return nil, errors.Newf("Service \"%s\" not found", name)
}

func (a *Application) GetServices() []Service {
	return a.services
}

func (a *Application) HasService(name string) bool {
	_, err := a.GetService(name)
	return err == nil
}

func (a *Application) RegisterService(service Service) error {
	if _, err := a.GetService(service.GetName()); err == nil {
		return errors.Newf("Service \"%s\" already exists", service.GetName())
	}

	a.services = append(a.services, service)
	return nil
}

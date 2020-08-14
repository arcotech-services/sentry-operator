package controllers

import "github.com/jace-ys/sentry-operator/pkg/sentry"

type Sentry struct {
	Organization string
	Client       *SentryClient
}

type SentryClient struct {
	Projects SentryProjects
	Teams    SentryTeams
}

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . SentryProjects
type SentryProjects interface {
	Update(organizationSlug, projectSlug string, params *sentry.UpdateProjectParams) (*sentry.Project, *sentry.Response, error)
	Delete(organizationSlug, projectSlug string) (*sentry.Response, error)
}

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . SentryTeams
type SentryTeams interface {
	ListProjects(organizationSlug, teamSlug string) ([]sentry.Project, *sentry.Response, error)
	CreateProject(organizationSlug, teamSlug string, params *sentry.CreateProjectParams) (*sentry.Project, *sentry.Response, error)
}

func removeFinalizer(finalizers []string, name string) []string {
	var result []string
	for _, item := range finalizers {
		if item == name {
			continue
		}

		result = append(result, item)
	}

	return result
}

func containsFinalizer(finalizers []string, name string) bool {
	for _, finalizer := range finalizers {
		if finalizer == name {
			return true
		}
	}

	return false
}

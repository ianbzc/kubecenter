package pods

import (
	"context"
)

const (
	AppName = "pods"
)

type Service interface {
	GetPods(ctx context.Context) (*Pods, error)
	GetNamespaceList(ctx context.Context) (*NamespaceList, error)
	GetPodsListUnderNamespaceWithKeyword(ctx context.Context, namespace string, keyword string) (*PodsList, error)
}

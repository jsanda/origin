// This file was automatically generated by informer-gen

package v1

import (
	authorization_v1 "github.com/openshift/api/authorization/v1"
	clientset "github.com/openshift/origin/pkg/authorization/generated/clientset"
	internalinterfaces "github.com/openshift/origin/pkg/authorization/generated/informers/externalversions/internalinterfaces"
	v1 "github.com/openshift/origin/pkg/authorization/generated/listers/authorization/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// ClusterRoleInformer provides access to a shared informer and lister for
// ClusterRoles.
type ClusterRoleInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ClusterRoleLister
}

type clusterRoleInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

// NewClusterRoleInformer constructs a new informer for ClusterRole type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterRoleInformer(client clientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
				return client.AuthorizationV1().ClusterRoles().List(options)
			},
			WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
				return client.AuthorizationV1().ClusterRoles().Watch(options)
			},
		},
		&authorization_v1.ClusterRole{},
		resyncPeriod,
		indexers,
	)
}

func defaultClusterRoleInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewClusterRoleInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc})
}

func (f *clusterRoleInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&authorization_v1.ClusterRole{}, defaultClusterRoleInformer)
}

func (f *clusterRoleInformer) Lister() v1.ClusterRoleLister {
	return v1.NewClusterRoleLister(f.Informer().GetIndexer())
}

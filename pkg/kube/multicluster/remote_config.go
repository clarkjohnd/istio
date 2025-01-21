package multicluster

import (
	"google.golang.org/protobuf/runtime/protoimpl"
	meshapi "istio.io/api/mesh/v1alpha1"
)

type RemoteSecret struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	remoteConfig *remoteMeshConfig `protobuf:"bytes,1,rep,name=remote_mesh_config,json=remoteMeshConfig,proto3" json:"remote_mesh_config,omitempty"`
}

type remoteMeshConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// discoverySelectors specific to the remote cluster
	DiscoverySelectors []*meshapi.LabelSelector `protobuf:"bytes,2,rep,name=discovery_selectors,json=discoverySelectors,proto3" json:"discovery_selectors,omitempty"`
}

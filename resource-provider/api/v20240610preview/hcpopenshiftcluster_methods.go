package v20240610preview

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	api2 "github.com/Azure/ARO-HCP/toolkit/api"
	"slices"
)

func (v *version) NewHCPOpenShiftCluster(from *api2.HCPOpenShiftCluster) api2.VersionedHCPOpenShiftCluster {
	out := &HCPOpenShiftCluster{
		Properties: HCPOpenShiftClusterProperties{
			ProvisioningState: from.Properties.ProvisioningState,
			ClusterProfile: ClusterProfile{
				ControlPlaneVersion: from.Properties.ClusterProfile.ControlPlaneVersion,
				SubnetID:            from.Properties.ClusterProfile.SubnetID,
			},
			APIProfile: APIProfile{
				URL:        from.Properties.APIProfile.URL,
				IP:         from.Properties.APIProfile.IP,
				Visibility: Visibility(from.Properties.APIProfile.Visibility),
			},
			ConsoleProfile: ConsoleProfile{
				URL: from.Properties.ConsoleProfile.URL,
			},
			IngressProfiles: make([]IngressProfile, 0, len(from.Properties.IngressProfiles)),
			NetworkProfile: NetworkProfile{
				PodCIDR:           from.Properties.NetworkProfile.PodCIDR,
				ServiceCIDR:       from.Properties.NetworkProfile.ServiceCIDR,
				MachineCIDR:       from.Properties.NetworkProfile.MachineCIDR,
				HostPrefix:        from.Properties.NetworkProfile.HostPrefix,
				OutboundType:      OutboundType(from.Properties.NetworkProfile.OutboundType),
				PreconfiguredNSGs: PreconfiguredNSGs(from.Properties.NetworkProfile.PreconfiguredNSGs),
			},
			NodePoolProfiles: make([]NodePoolProfile, 0, len(from.Properties.NodePoolProfiles)),
			EtcdEncryption: EtcdEncryptionProfile{
				DiscEncryptionSetID: from.Properties.EtcdEncryption.DiscEncryptionSetID,
			},
			AutoRepair:    from.Properties.AutoRepair,
			Labels:        slices.Clone(from.Properties.Labels),
			OIDCIssuerURL: from.Properties.OIDCIssuerURL,
		},
	}

	out.TrackedResource.Copy(&from.TrackedResource)

	for _, item := range from.Properties.IngressProfiles {
		out.Properties.IngressProfiles = append(
			out.Properties.IngressProfiles,
			IngressProfile{
				IP:         item.IP,
				Name:       item.Name,
				Visibility: Visibility(item.Visibility),
			})
	}

	for _, item := range from.Properties.NodePoolProfiles {
		out.Properties.NodePoolProfiles = append(
			out.Properties.NodePoolProfiles,
			NodePoolProfile{
				Name:             item.Name,
				Replicas:         item.Replicas,
				SubnetID:         item.SubnetID,
				EncryptionAtHost: item.EncryptionAtHost,
				VMSize:           item.VMSize,
				Autoscaling: NodePoolAutoscaling{
					MinReplicas: item.Autoscaling.MinReplicas,
					MaxReplicas: item.Autoscaling.MaxReplicas,
				},
			})
	}

	return out
}

func (c *HCPOpenShiftCluster) Normalize(out *api2.HCPOpenShiftCluster) {
	c.TrackedResource.Copy(&out.TrackedResource)
	out.Properties.ProvisioningState = c.Properties.ProvisioningState
	out.Properties.ClusterProfile.ControlPlaneVersion = c.Properties.ClusterProfile.ControlPlaneVersion
	out.Properties.ClusterProfile.SubnetID = c.Properties.ClusterProfile.SubnetID
	out.Properties.APIProfile.URL = c.Properties.APIProfile.URL
	out.Properties.APIProfile.IP = c.Properties.APIProfile.IP
	out.Properties.APIProfile.Visibility = api2.Visibility(c.Properties.APIProfile.Visibility)
	out.Properties.ConsoleProfile.URL = c.Properties.ConsoleProfile.URL
	out.Properties.IngressProfiles = make([]api2.IngressProfile, 0, len(c.Properties.IngressProfiles))
	for _, item := range c.Properties.IngressProfiles {
		out.Properties.IngressProfiles = append(
			out.Properties.IngressProfiles,
			api2.IngressProfile{
				IP:         item.IP,
				Name:       item.Name,
				Visibility: api2.Visibility(item.Visibility),
			})
	}
	out.Properties.NetworkProfile.PodCIDR = c.Properties.NetworkProfile.PodCIDR
	out.Properties.NetworkProfile.ServiceCIDR = c.Properties.NetworkProfile.ServiceCIDR
	out.Properties.NetworkProfile.HostPrefix = c.Properties.NetworkProfile.HostPrefix
	out.Properties.NetworkProfile.OutboundType = api2.OutboundType(c.Properties.NetworkProfile.OutboundType)
	out.Properties.NetworkProfile.PreconfiguredNSGs = api2.PreconfiguredNSGs(c.Properties.NetworkProfile.PreconfiguredNSGs)
	out.Properties.NodePoolProfiles = make([]api2.NodePoolProfile, 0, len(c.Properties.NodePoolProfiles))
	for _, item := range c.Properties.NodePoolProfiles {
		out.Properties.NodePoolProfiles = append(
			out.Properties.NodePoolProfiles,
			api2.NodePoolProfile{
				Name:             item.Name,
				Replicas:         item.Replicas,
				SubnetID:         item.SubnetID,
				EncryptionAtHost: item.EncryptionAtHost,
				VMSize:           item.VMSize,
				Autoscaling: api2.NodePoolAutoscaling{
					MinReplicas: item.Autoscaling.MinReplicas,
					MaxReplicas: item.Autoscaling.MaxReplicas,
				},
			})
	}
	out.Properties.EtcdEncryption = api2.EtcdEncryptionProfile{
		DiscEncryptionSetID: c.Properties.EtcdEncryption.DiscEncryptionSetID,
	}
	out.Properties.AutoRepair = c.Properties.AutoRepair
	out.Properties.Labels = slices.Clone(c.Properties.Labels)
	out.Properties.OIDCIssuerURL = c.Properties.OIDCIssuerURL
}

func (c *HCPOpenShiftCluster) ValidateStatic() error {
	return nil
}

package util

import (
	v1 "k8s.io/api/core/v1"
)

type PodSpec struct {
	NS                  string                      `default:""`
	Pvclaims            []*v1.PersistentVolumeClaim `default:"nil"`
	InlineVolumeSources []*v1.VolumeSource          `default:"nil"`
	IsPrivileged        bool                        `default:"false"`
	HostNetwork         bool                        `default:"false"`
	Command             string                      `default:""`
	HostIPC             bool                        `default:"false"`
	HostPID             bool                        `default:"false"`
	seLinuxLabel        *v1.SELinuxOptions
	fsGroup             *int64
	RunAsNonRoot        bool               `default:"true"`
	Capability          []v1.Capability    `default:"{}"`
	Ports               []v1.ContainerPort `default:"nil"`
}

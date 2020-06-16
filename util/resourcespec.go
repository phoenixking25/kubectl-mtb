package util

import (
	"fmt"

	"github.com/creasty/defaults"
	v1 "k8s.io/api/core/v1"
)

// https://github.com/creasty/defaults#usage
type PodSpec struct {
	NS                       string                      `default:""`
	Pvclaims                 []*v1.PersistentVolumeClaim `default:"-"`
	InlineVolumeSources      []*v1.VolumeSource          `default:"-"`
	HostNetwork              bool                        `default:"false"`
	Command                  string                      `default:""`
	HostIPC                  bool                        `default:"false"`
	HostPID                  bool                        `default:"false"`
	seLinuxLabel             *v1.SELinuxOptions
	fsGroup                  *int64
	RunAsNonRoot             bool               `default:"-"`
	IsPrivileged             bool               `default:"false"`
	Capability               []v1.Capability    `default:"-"`
	Ports                    []v1.ContainerPort `default:"-"`
	AllowPrivilegeEscalation bool               `default:"false"`
}

func (p *PodSpec) SetDefaults() error {
	if err := defaults.Set(p); err != nil {
		return fmt.Errorf("it should not return an error: %v", err)
	}
	return nil
}

type ServiceConfig struct {
	Type     v1.ServiceType
	Selector map[string]string
}

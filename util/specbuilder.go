package util

import (
	"fmt"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/uuid"
	imageutils "k8s.io/kubernetes/test/utils/image"
)

func MakeSecPod(ns string, pvclaims []*v1.PersistentVolumeClaim, inlineVolumeSources []*v1.VolumeSource, isPrivileged bool, command string, hostIPC bool, hostPID bool, seLinuxLabel *v1.SELinuxOptions, fsGroup *int64, runAsNonRoot bool, capability v1.Capability) *v1.Pod {
	if len(command) == 0 {
		command = "trap exit TERM; while true; do sleep 1; done"
	}
	podName := "security-context-" + string(uuid.NewUUID())
	if fsGroup == nil {
		fsGroup = func(i int64) *int64 {
			return &i
		}(1000)
	}
	podSpec := &v1.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      podName,
			Namespace: ns,
		},
		Spec: v1.PodSpec{
			HostIPC: hostIPC,
			HostPID: hostPID,
			SecurityContext: &v1.PodSecurityContext{
				FSGroup:      fsGroup,
				RunAsNonRoot: &runAsNonRoot,
			},
			Containers: []v1.Container{
				{
					Name:    "write-pod",
					Image:   imageutils.GetE2EImage(imageutils.BusyBox),
					Command: []string{"/bin/sh"},
					Args:    []string{"-c", command},
					SecurityContext: &v1.SecurityContext{
						Privileged: &isPrivileged,
						Capabilities: &v1.Capabilities{
							Add: []v1.Capability{
								capability,
							},
						},
					},
				},
			},
			RestartPolicy: v1.RestartPolicyOnFailure,
		},
	}
	var volumeMounts = make([]v1.VolumeMount, 0)
	var volumeDevices = make([]v1.VolumeDevice, 0)
	var volumes = make([]v1.Volume, len(pvclaims)+len(inlineVolumeSources))
	volumeIndex := 0
	for _, pvclaim := range pvclaims {
		volumename := fmt.Sprintf("volume%v", volumeIndex+1)
		if pvclaim.Spec.VolumeMode != nil && *pvclaim.Spec.VolumeMode == v1.PersistentVolumeBlock {
			volumeDevices = append(volumeDevices, v1.VolumeDevice{Name: volumename, DevicePath: "/mnt/" + volumename})
		} else {
			volumeMounts = append(volumeMounts, v1.VolumeMount{Name: volumename, MountPath: "/mnt/" + volumename})
		}

		volumes[volumeIndex] = v1.Volume{Name: volumename, VolumeSource: v1.VolumeSource{PersistentVolumeClaim: &v1.PersistentVolumeClaimVolumeSource{ClaimName: pvclaim.Name, ReadOnly: false}}}
		volumeIndex++
	}
	for _, src := range inlineVolumeSources {
		volumename := fmt.Sprintf("volume%v", volumeIndex+1)
		// In-line volumes can be only filesystem, not block.
		volumeMounts = append(volumeMounts, v1.VolumeMount{Name: volumename, MountPath: "/mnt/" + volumename})
		volumes[volumeIndex] = v1.Volume{Name: volumename, VolumeSource: *src}
		volumeIndex++
	}

	podSpec.Spec.Containers[0].VolumeMounts = volumeMounts
	podSpec.Spec.Containers[0].VolumeDevices = volumeDevices
	podSpec.Spec.Volumes = volumes
	podSpec.Spec.SecurityContext.SELinuxOptions = seLinuxLabel
	return podSpec
}

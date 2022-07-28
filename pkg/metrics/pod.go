package metrics

/*

# HELP kube_pod_completion_time Completion time in unix timestamp for a pod.
# Type kube_pod_completion_time gauge

# HELP kube_pod_container_info Information about a container in a pod.
# Type kube_pod_container_info gauge

# HELP kube_pod_container_resource_limits The number of requested limit resource by a container.
# Type kube_pod_container_resource_limits gauge

# HELP kube_pod_container_resource_requests The number of requested request resource by a container.
# Type kube_pod_container_resource_requests gauge

# HELP kube_pod_container_state_started Start time in unix timestamp for a pod container.
# Type kube_pod_container_state_started gauge

# HELP kube_pod_container_status_last_terminated_reason Describes the last reason the container was in terminated state.
# Type kube_pod_container_status_last_terminated_reason gauge

# HELP kube_pod_container_status_ready Describes whether the containers readiness check succeeded.
# Type kube_pod_container_status_ready gauge

# HELP kube_pod_container_status_restarts_total The number of container restarts per container.
# Type kube_pod_container_status_restarts_total counter

# HELP kube_pod_container_status_running Describes whether the container is currently in running state.
# Type kube_pod_container_status_running gauge

# HELP kube_pod_container_status_terminated Describes whether the container is currently in terminated state.
# Type kube_pod_container_status_terminated gauge

# HELP kube_pod_container_status_terminated_reason Describes the reason the container is currently in terminated state.
# Type kube_pod_container_status_terminated_reason gauge

# HELP kube_pod_container_status_waiting Describes whether the container is currently in waiting state.
# Type kube_pod_container_status_waiting gauge

# HELP kube_pod_container_status_waiting_reason Describes the reason the container is currently in waiting state.
# Type kube_pod_container_status_waiting_reason gauge

# HELP kube_pod_created Unix creation timestamp
# Type kube_pod_created gauge

# HELP kube_pod_deletion_timestamp Unix deletion timestamp
# Type kube_pod_deletion_timestamp gauge

# HELP kube_pod_info Information about pod.
# Type kube_pod_info gauge

# HELP kube_pod_init_container_info Information about an init container in a pod.
# Type kube_pod_init_container_info gauge

# HELP kube_pod_init_container_resource_limits The number of requested limit resource by an init container.
# Type kube_pod_init_container_resource_limits gauge

# HELP kube_pod_init_container_resource_requests The number of requested request resource by an init container.
# Type kube_pod_init_container_resource_requests gauge

# HELP kube_pod_init_container_status_last_terminated_reason Describes the last reason the init container was in terminated state.
# Type kube_pod_init_container_status_last_terminated_reason gauge

# HELP kube_pod_init_container_status_ready Describes whether the init containers readiness check succeeded.
# Type kube_pod_init_container_status_ready gauge

# HELP kube_pod_init_container_status_restarts_total The number of restarts for the init container.
# Type kube_pod_init_container_status_restarts_total counter

# HELP kube_pod_init_container_status_running Describes whether the init container is currently in running state.
# Type kube_pod_init_container_status_running gauge

# HELP kube_pod_init_container_status_terminated Describes whether the init container is currently in terminated state.
# Type kube_pod_init_container_status_terminated gauge

# HELP kube_pod_init_container_status_terminated_reason Describes the reason the init container is currently in terminated state.
# Type kube_pod_init_container_status_terminated_reason gauge

# HELP kube_pod_init_container_status_waiting Describes whether the init container is currently in waiting state.
# Type kube_pod_init_container_status_waiting gauge

# HELP kube_pod_init_container_status_waiting_reason Describes the reason the init container is currently in waiting state.
# Type kube_pod_init_container_status_waiting_reason gauge

# HELP kube_pod_annotations Kubernetes annotations converted to Prometheus labels.
# Type kube_pod_annotations gauge

# HELP kube_pod_labels Kubernetes labels converted to Prometheus labels.
# Type kube_pod_labels gauge

# HELP kube_pod_overhead_cpu_cores The pod overhead in regards to cpu cores associated with running a pod.
# Type kube_pod_overhead_cpu_cores gauge

# HELP kube_pod_overhead_memory_bytes The pod overhead in regards to memory associated with running a pod.
# Type kube_pod_overhead_memory_bytes gauge

# HELP kube_pod_owner Information about the Pod's owner.
# Type kube_pod_owner gauge

# HELP kube_pod_restart_policy Describes the restart policy in use by this pod.
# Type kube_pod_restart_policy gauge

# HELP kube_pod_runtimeclass_name_info The runtimeclass associated with the pod.
# Type kube_pod_runtimeclass_name_info gauge

# HELP kube_pod_spec_volumes_persistentvolumeclaims_info Information about persistentvolumeclaim volumes in a pod.
# Type kube_pod_spec_volumes_persistentvolumeclaims_info gauge

# HELP kube_pod_spec_volumes_persistentvolumeclaims_readonly Describes whether a persistentvolumeclaim is mounted read only.
# Type kube_pod_spec_volumes_persistentvolumeclaims_readonly gauge

# HELP kube_pod_start_time Start time in unix timestamp for a pod.
# Type kube_pod_start_time gauge

# HELP kube_pod_status_phase The pods current phase.
# Type kube_pod_status_phase gauge

# HELP kube_pod_status_ready Describes whether the pod is ready to serve requests.
# Type kube_pod_status_ready gauge

# HELP kube_pod_status_reason The pod status reasons
# Type kube_pod_status_reason gauge

# HELP kube_pod_status_scheduled Describes the status of the scheduling process for the pod.
# Type kube_pod_status_scheduled gauge

# HELP kube_pod_status_scheduled_time Unix timestamp when pod moved into scheduled status
# Type kube_pod_status_scheduled_time gauge

# HELP kube_pod_status_unschedulable Describes the unschedulable status for the pod.
# Type kube_pod_status_unschedulable gauge

*/

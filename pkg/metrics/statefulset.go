package metrics

/*
# HELP kube_statefulset_created Unix creation timestamp
# Type kube_statefulset_created gauge

# HELP kube_statefulset_status_replicas The number of replicas per StatefulSet.
# Type kube_statefulset_status_replicas gauge

# HELP kube_statefulset_status_replicas_available The number of available replicas per StatefulSet.
# Type kube_statefulset_status_replicas_available gauge

# HELP kube_statefulset_status_replicas_current The number of current replicas per StatefulSet.
# Type kube_statefulset_status_replicas_current gauge

# HELP kube_statefulset_status_replicas_ready The number of ready replicas per StatefulSet.
# Type kube_statefulset_status_replicas_ready gauge

# HELP kube_statefulset_status_replicas_updated The number of updated replicas per StatefulSet.
# Type kube_statefulset_status_replicas_updated gauge

# HELP kube_statefulset_status_observed_generation The generation observed by the StatefulSet controller.
# Type kube_statefulset_status_observed_generation gauge

# HELP kube_statefulset_replicas Number of desired pods for a StatefulSet.
# Type kube_statefulset_replicas gauge

# HELP kube_statefulset_metadata_generation Sequence number representing a specific generation of the desired state for the StatefulSet.
# Type kube_statefulset_metadata_generation gauge

# HELP kube_statefulset_annotations Kubernetes annotations converted to Prometheus labels.
# Type kube_statefulset_annotations gauge

# HELP kube_statefulset_labels Kubernetes labels converted to Prometheus labels.
# Type kube_statefulset_labels gauge

# HELP kube_statefulset_status_current_revision Indicates the version of the StatefulSet used to generate Pods in the sequence [0,currentReplicas).
# Type kube_statefulset_status_current_revision gauge

# HELP kube_statefulset_status_update_revision Indicates the version of the StatefulSet used to generate Pods in the sequence [replicas-updatedReplicas,replicas)
# Type kube_statefulset_status_update_revision gauge
*/

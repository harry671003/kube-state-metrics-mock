package metrics

/*
# HELP kube_replicaset_created Unix creation timestamp
# Type kube_replicaset_created gauge

# HELP kube_replicaset_status_replicas The number of replicas per ReplicaSet.
# Type kube_replicaset_status_replicas gauge

# HELP kube_replicaset_status_fully_labeled_replicas The number of fully labeled replicas per ReplicaSet.
# Type kube_replicaset_status_fully_labeled_replicas gauge

# HELP kube_replicaset_status_ready_replicas The number of ready replicas per ReplicaSet.
# Type kube_replicaset_status_ready_replicas gauge

# HELP kube_replicaset_status_observed_generation The generation observed by the ReplicaSet controller.
# Type kube_replicaset_status_observed_generation gauge

# HELP kube_replicaset_spec_replicas Number of desired pods for a ReplicaSet.
# Type kube_replicaset_spec_replicas gauge

# HELP kube_replicaset_metadata_generation Sequence number representing a specific generation of the desired state.
# Type kube_replicaset_metadata_generation gauge

# HELP kube_replicaset_owner Information about the ReplicaSet's owner.
# Type kube_replicaset_owner gauge

# HELP kube_replicaset_annotations Kubernetes annotations converted to Prometheus labels.
# Type kube_replicaset_annotations gauge

# HELP kube_replicaset_labels Kubernetes labels converted to Prometheus labels.
# Type kube_replicaset_labels gauge
*/

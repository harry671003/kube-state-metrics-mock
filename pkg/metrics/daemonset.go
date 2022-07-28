package metrics

/*

# HELP kube_daemonset_created Unix creation timestamp
# Type kube_daemonset_created gauge

# HELP kube_daemonset_status_current_number_scheduled The number of nodes running at least one daemon pod and are supposed to.
# Type kube_daemonset_status_current_number_scheduled gauge

# HELP kube_daemonset_status_desired_number_scheduled The number of nodes that should be running the daemon pod.
# Type kube_daemonset_status_desired_number_scheduled gauge

# HELP kube_daemonset_status_number_available The number of nodes that should be running the daemon pod and have one or more of the daemon pod running and available
# Type kube_daemonset_status_number_available gauge

# HELP kube_daemonset_status_number_misscheduled The number of nodes running a daemon pod but are not supposed to.
# Type kube_daemonset_status_number_misscheduled gauge

# HELP kube_daemonset_status_number_ready The number of nodes that should be running the daemon pod and have one or more of the daemon pod running and ready.
# Type kube_daemonset_status_number_ready gauge

# HELP kube_daemonset_status_number_unavailable The number of nodes that should be running the daemon pod and have none of the daemon pod running and available
# Type kube_daemonset_status_number_unavailable gauge

# HELP kube_daemonset_status_observed_generation The most recent generation observed by the daemon set controller.
# Type kube_daemonset_status_observed_generation gauge

# HELP kube_daemonset_status_updated_number_scheduled The total number of nodes that are running updated daemon pod
# Type kube_daemonset_status_updated_number_scheduled gauge

# HELP kube_daemonset_metadata_generation Sequence number representing a specific generation of the desired state.
# Type kube_daemonset_metadata_generation gauge

# HELP kube_daemonset_annotations Kubernetes annotations converted to Prometheus labels.
# Type kube_daemonset_annotations gauge

# HELP kube_daemonset_labels Kubernetes labels converted to Prometheus labels.
# Type kube_daemonset_labels gauge

*/

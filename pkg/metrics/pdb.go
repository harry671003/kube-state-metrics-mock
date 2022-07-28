package metrics

/*
# HELP kube_poddisruptionbudget_annotations Kubernetes annotations converted to Prometheus labels.
# Type kube_poddisruptionbudget_annotations gauge

# HELP kube_poddisruptionbudget_labels Kubernetes labels converted to Prometheus labels.
# Type kube_poddisruptionbudget_labels gauge

# HELP kube_poddisruptionbudget_created Unix creation timestamp
# Type kube_poddisruptionbudget_created gauge

# HELP kube_poddisruptionbudget_status_current_healthy Current number of healthy pods
# Type kube_poddisruptionbudget_status_current_healthy gauge

# HELP kube_poddisruptionbudget_status_desired_healthy Minimum desired number of healthy pods
# Type kube_poddisruptionbudget_status_desired_healthy gauge

# HELP kube_poddisruptionbudget_status_pod_disruptions_allowed Number of pod disruptions that are currently allowed
# Type kube_poddisruptionbudget_status_pod_disruptions_allowed gauge

# HELP kube_poddisruptionbudget_status_expected_pods Total number of pods counted by this disruption budget
# Type kube_poddisruptionbudget_status_expected_pods gauge

# HELP kube_poddisruptionbudget_status_observed_generation Most recent generation observed when updating this PDB status
# Type kube_poddisruptionbudget_status_observed_generation gauge
*/

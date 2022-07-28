package metrics

/*

# HELP kube_deployment_created Unix creation timestamp
# Type kube_deployment_created gauge

# HELP kube_deployment_status_replicas The number of replicas per deployment.
# Type kube_deployment_status_replicas gauge

# HELP kube_deployment_status_replicas_ready The number of ready replicas per deployment.
# Type kube_deployment_status_replicas_ready gauge

# HELP kube_deployment_status_replicas_available The number of available replicas per deployment.
# Type kube_deployment_status_replicas_available gauge

# HELP kube_deployment_status_replicas_unavailable The number of unavailable replicas per deployment.
# Type kube_deployment_status_replicas_unavailable gauge

# HELP kube_deployment_status_replicas_updated The number of updated replicas per deployment.
# Type kube_deployment_status_replicas_updated gauge

# HELP kube_deployment_status_observed_generation The generation observed by the deployment controller.
# Type kube_deployment_status_observed_generation gauge

# HELP kube_deployment_status_condition The current status conditions of a deployment.
# Type kube_deployment_status_condition gauge

# HELP kube_deployment_spec_replicas Number of desired pods for a deployment.
# Type kube_deployment_spec_replicas gauge

# HELP kube_deployment_spec_paused Whether the deployment is paused and will not be processed by the deployment controller.
# Type kube_deployment_spec_paused gauge

# HELP kube_deployment_spec_strategy_rollingupdate_max_unavailable Maximum number of unavailable replicas during a rolling update of a deployment.
# Type kube_deployment_spec_strategy_rollingupdate_max_unavailable gauge

# HELP kube_deployment_spec_strategy_rollingupdate_max_surge Maximum number of replicas that can be scheduled above the desired number of replicas during a rolling update of a deployment.
# Type kube_deployment_spec_strategy_rollingupdate_max_surge gauge

# HELP kube_deployment_metadata_generation Sequence number representing a specific generation of the desired state.
# Type kube_deployment_metadata_generation gauge

# HELP kube_deployment_annotations Kubernetes annotations converted to Prometheus labels.
# Type kube_deployment_annotations gauge

# HELP kube_deployment_labels Kubernetes labels converted to Prometheus labels.
# Type kube_deployment_labels gauge

*/

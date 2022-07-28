package metrics

/*

# HELP kube_horizontalpodautoscaler_info Information about this autoscaler.
# Type kube_horizontalpodautoscaler_info gauge

# HELP kube_horizontalpodautoscaler_metadata_generation The generation observed by the HorizontalPodAutoscaler controller.
# Type kube_horizontalpodautoscaler_metadata_generation gauge

# HELP kube_horizontalpodautoscaler_spec_max_replicas Upper limit for the number of pods that can be set by the autoscaler; cannot be smaller than MinReplicas.
# Type kube_horizontalpodautoscaler_spec_max_replicas gauge

# HELP kube_horizontalpodautoscaler_spec_min_replicas Lower limit for the number of pods that can be set by the autoscaler, default 1.
# Type kube_horizontalpodautoscaler_spec_min_replicas gauge

# HELP kube_horizontalpodautoscaler_spec_target_metric The metric specifications used by this autoscaler when calculating the desired replica count.
# Type kube_horizontalpodautoscaler_spec_target_metric gauge

# HELP kube_horizontalpodautoscaler_status_current_replicas Current number of replicas of pods managed by this autoscaler.
# Type kube_horizontalpodautoscaler_status_current_replicas gauge

# HELP kube_horizontalpodautoscaler_status_desired_replicas Desired number of replicas of pods managed by this autoscaler.
# Type kube_horizontalpodautoscaler_status_desired_replicas gauge

# HELP kube_horizontalpodautoscaler_annotations Kubernetes annotations converted to Prometheus labels.
# Type kube_horizontalpodautoscaler_annotations gauge

# HELP kube_horizontalpodautoscaler_labels Kubernetes labels converted to Prometheus labels.
# Type kube_horizontalpodautoscaler_labels gauge

# HELP kube_horizontalpodautoscaler_status_condition The condition of this autoscaler.
# Type kube_horizontalpodautoscaler_status_condition gauge

*/

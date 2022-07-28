package metrics

/*

# HELP kube_job_annotations Kubernetes annotations converted to Prometheus labels.
# Type kube_job_annotations gauge

# HELP kube_job_labels Kubernetes labels converted to Prometheus labels.
# Type kube_job_labels gauge

# HELP kube_job_info Information about job.
# Type kube_job_info gauge

# HELP kube_job_created Unix creation timestamp
# Type kube_job_created gauge

# HELP kube_job_spec_parallelism The maximum desired number of pods the job should run at any given time.
# Type kube_job_spec_parallelism gauge

# HELP kube_job_spec_completions The desired number of successfully finished pods the job should be run with.
# Type kube_job_spec_completions gauge

# HELP kube_job_spec_active_deadline_seconds The duration in seconds relative to the startTime that the job may be active before the system tries to terminate it.
# Type kube_job_spec_active_deadline_seconds gauge

# HELP kube_job_status_succeeded The number of pods which reached Phase Succeeded.
# Type kube_job_status_succeeded gauge

# HELP kube_job_status_failed The number of pods which reached Phase Failed and the reason for failure.
# Type kube_job_status_failed gauge

# HELP kube_job_status_active The number of actively running pods.
# Type kube_job_status_active gauge

# HELP kube_job_complete The job has completed its execution.
# Type kube_job_complete gauge

# HELP kube_job_failed The job has failed its execution.
# Type kube_job_failed gauge

# HELP kube_job_status_start_time StartTime represents time when the job was acknowledged by the Job Manager.
# Type kube_job_status_start_time gauge

# HELP kube_job_status_completion_time CompletionTime represents time when the job was completed.
# Type kube_job_status_completion_time gauge

# HELP kube_job_owner Information about the Job's owner.
# Type kube_job_owner gauge

*/

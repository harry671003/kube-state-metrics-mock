package metrics

/*
# HELP kube_cronjob_annotations Kubernetes annotations converted to Prometheus labels.
# Type kube_cronjob_annotations gauge

# HELP kube_cronjob_labels Kubernetes labels converted to Prometheus labels.
# Type kube_cronjob_labels gauge

# HELP kube_cronjob_info Info about cronjob.
# Type kube_cronjob_info gauge

# HELP kube_cronjob_created Unix creation timestamp
# Type kube_cronjob_created gauge

# HELP kube_cronjob_status_active Active holds pointers to currently running jobs.
# Type kube_cronjob_status_active gauge

# HELP kube_cronjob_status_last_schedule_time LastScheduleTime keeps information of when was the last time the job was successfully scheduled.
# Type kube_cronjob_status_last_schedule_time gauge

# HELP kube_cronjob_spec_suspend Suspend flag tells the controller to suspend subsequent executions.
# Type kube_cronjob_spec_suspend gauge

# HELP kube_cronjob_spec_starting_deadline_seconds Deadline in seconds for starting the job if it misses scheduled time for any reason.
# Type kube_cronjob_spec_starting_deadline_seconds gauge

# HELP kube_cronjob_next_schedule_time Next time the cronjob should be scheduled. The time after lastScheduleTime, or after the cron job's creation time if it's never been scheduled. Use this to determine if the job is delayed.
# Type kube_cronjob_next_schedule_time gauge

# HELP kube_cronjob_metadata_resource_version Resource version representing a specific version of the cronjob.
# Type kube_cronjob_metadata_resource_version gauge

# HELP kube_cronjob_spec_successful_job_history_limit Successful job history limit tells the controller how many completed jobs should be preserved.
# Type kube_cronjob_spec_successful_job_history_limit gauge

# HELP kube_cronjob_spec_failed_job_history_limit Failed job history limit tells the controller how many failed jobs should be preserved.
# Type kube_cronjob_spec_failed_job_history_limit gauge

*/

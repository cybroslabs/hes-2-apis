# Model: io.clbs.openhes.models.acquisition.JobStatus

Sub-message containing job status info

## Fields

| Field | Type | Description |
| --- | --- | --- |
| status | [io.clbs.openhes.models.acquisition.JobStatusCode](model-io-clbs-openhes-models-acquisition-jobstatuscode.md) | The status of the job. |
| code | [io.clbs.openhes.models.acquisition.JobErrorCode](model-io-clbs-openhes-models-acquisition-joberrorcode.md) | The error code of the job. |
| results | [io.clbs.openhes.models.acquisition.ActionResult](model-io-clbs-openhes-models-acquisition-actionresult.md) | The result data for all job actions. |
| createdAt | google.protobuf.Timestamp | The creation timestamp of the job. |
| startedAt | google.protobuf.Timestamp | The start timestamp of the job. |
| finishedAt | google.protobuf.Timestamp | The finish timestamp of the job. |
| attemptsDone | int32 | The number of attempts already done. |
| deviceInfo | [io.clbs.openhes.models.acquisition.DeviceInfo](model-io-clbs-openhes-models-acquisition-deviceinfo.md) | The device info. It contains the data from the action ACTION_TYPE_GET_DEVICE_INFO. |
| queueId | int64 | The internal queue identifier set by the Taskmaster when the job is queued to process. |


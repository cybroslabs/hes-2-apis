# Model: io.clbs.openhes.models.acquisition.BulkSpec

## Fields

| Field | Type | Description |
| --- | --- | --- |
| correlationId | string | @gqltype: UUID<br><br>The correlation identifier, e.g. to define relation to non-homogenous group. |
| devices | [io.clbs.openhes.models.acquisition.ListOfJobDeviceId](model-io-clbs-openhes-models-acquisition-listofjobdeviceid.md) | The list of devices in the bulk. |
| deviceGroupId | string | The device group identifier. |
| settings | [io.clbs.openhes.models.acquisition.JobSettings](model-io-clbs-openhes-models-acquisition-jobsettings.md) | The bulk-shared job settings. |
| actions | [io.clbs.openhes.models.acquisition.JobActionSet](model-io-clbs-openhes-models-acquisition-jobactionset.md) | The list actions to be executed. |
| webhookUrl | string | The webhook URL to call when the bulk is completed. |


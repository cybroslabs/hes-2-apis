# ApiService - Bulks

## CreateBulk

Starts a new bulk of jobs.

```proto
CreateBulk(io.clbs.openhes.models.acquisition.CreateBulkRequest) returns (google.protobuf.StringValue)
```

Input: [io.clbs.openhes.models.acquisition.CreateBulkRequest](model-io-clbs-openhes-models-acquisition-createbulkrequest.md)
Output: [google.protobuf.StringValue](model-google-protobuf-stringvalue.md)

## ListBulks

Retrieves the list of bulks.

```proto
ListBulks(io.clbs.openhes.models.common.ListSelector) returns (io.clbs.openhes.models.acquisition.ListOfBulk)
```

Input: [io.clbs.openhes.models.common.ListSelector](model-io-clbs-openhes-models-common-listselector.md)
Output: [io.clbs.openhes.models.acquisition.ListOfBulk](model-io-clbs-openhes-models-acquisition-listofbulk.md)

## GetBulk

Retrieves the bulk info and status.

```proto
GetBulk(google.protobuf.StringValue) returns (io.clbs.openhes.models.acquisition.Bulk)
```

Output: 

## CancelBulk

Cancels the bulk of jobs.

```proto
CancelBulk(google.protobuf.StringValue)
```

## GetBulkJob

Retrieves the job status.

```proto
GetBulkJob(google.protobuf.StringValue) returns (io.clbs.openhes.models.acquisition.BulkJob)
```

Output: 


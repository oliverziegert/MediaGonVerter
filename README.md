# MediaService

## Bucket Lifecycle Policy
As we don't have expiration logic implemented in the service, we expect the S3 backend to automatically expire and remove uploaded files.

One must ensure that the following lifecycle policy gets applied to each tenant bucket:
```xml
<LifecycleConfiguration>
    <Rule>
        <Status>Enabled</Status>
        <Filter>
            <Prefix>media/</Prefix>
        </Filter>
        <Expiration>
            <Days>180</Days>
        </Expiration>
    </Rule>
</LifecycleConfiguration>
```
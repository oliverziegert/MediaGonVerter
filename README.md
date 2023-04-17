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

## Contributors ✨

Thanks go to my colleagues:
- [@InterstellarFox](https://github.com/InterstellarFox)
- [@mischcon](https://github.com/mischcon)

## Copyright and License

Copyright Oliver Ziegert. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in
compliance with the License. You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is
distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
implied. See the License for the specific language governing permissions and limitations under the
License.

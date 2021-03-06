# Scotty
Scotty is a high performance, scalable collector for the
[tricorder](https://github.com/Symantec/tricorder) metrics publishing library.
Scotty provides a RESTful API to grab the latest polled metrics, and it can
push metrics to various persistent stores.

Please see the
[design document](https://docs.google.com/document/d/142Llj30LplgxWhOLOprqH59hS01EJ9iC1THV3no5oy0/pub)
and the
[online code documentation](https://godoc.org/github.com/Symantec/scotty)
for more information.

## Contributions

Prior to receiving information from any contributor, Symantec requires
that all contributors complete, sign, and submit Symantec Personal
Contributor Agreement (SPCA).  The purpose of the SPCA is to clearly
define the terms under which intellectual property has been
contributed to the project and thereby allow Symantec to defend the
project should there be a legal dispute regarding the software at some
future time. A signed SPCA is required to be on file before an
individual is given commit privileges to the Symantec open source
project.  Please note that the privilege to commit to the project is
conditional and may be revoked by Symantec.

If you are employed by a corporation, a Symantec Corporate Contributor
Agreement (SCCA) is also required before you may contribute to the
project.  If you are employed by a company, you may have signed an
employment agreement that assigns intellectual property ownership in
certain of your ideas or code to your company.  We require a SCCA to
make sure that the intellectual property in your contribution is
clearly contributed to the Symantec open source project, even if that
intellectual property had previously been assigned by you.

Please complete the SPCA and, if required, the SCCA and return to
Symantec at:

Symantec Corporation
Legal Department
Attention:  Product Legal Support Team
350 Ellis Street
Mountain View, CA 94043

Please be sure to keep a signed copy for your records.

## LICENSE

Copyright 2015 Symantec Corporation.

Licensed under the Apache License, Version 2.0 (the “License”); you
may not use this file except in compliance with the License.

You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0 Unless required by
applicable law or agreed to in writing, software distributed under the
License is distributed on an “AS IS” BASIS, WITHOUT WARRANTIES OR
CONDITIONS OF ANY KIND, either express or implied. See the License for
the specific language governing permissions and limitations under the
License.

## Building and Running Tests

When building scotty for the very first time, perform the following steps to
install the correct dependencies.

```
go get github.com/Symantec/scotty
cd $GOPATH/src/github.com/Symantec/scotty
make getdeps
```


From the top level directory of the scotty project:

To run all the tests

```
go test -v ./...
```

To rebuild after doing code changes

```
go install ./...
```


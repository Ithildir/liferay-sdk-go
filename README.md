# Liferay Go SDK

<img align="right" width="128" height="128" src="https://raw.githubusercontent.com/Ithildir/liferay-sdk-go/master/logo.png">

This repository contains a version of the [Liferay Mobile SDK](https://github.com/liferay/liferay-mobile-sdk) 
written in the [Go](http://golang.org) language.

The Liferay Go SDK is an accurate port of the [Android](https://github.com/liferay/liferay-mobile-sdk/tree/master/android) 
version and can be considered feature-complete. Nevertheless, it was created 
just for fun and basically as an experiment. Your mileage may vary.

## Setup

To install the Liferay Go SDK, simply run:

```bash
go get github.com/ithildir/liferay-sdk-go/liferay
```

Alternatively, the source code of the SDK for Liferay 6.2 is available in the
[*liferay*](https://github.com/Ithildir/liferay-sdk-go/tree/master/liferay)
folder. The [*tests*](https://github.com/Ithildir/liferay-sdk-go/tree/master/tests) 
folder contains the same unit tests of the [Liferay Android SDK](https://github.com/liferay/liferay-mobile-sdk/tree/master/android) 
and can be a good help to understand how to use the SDK.

The SDK has been developed and tested with Go versions 1.3 and 1.4.

## Build

The source code of the SDK for Liferay 6.2 is already provided. This repository
also contains the SDK builder, so you can easily generate a new SDK for a
different version of Liferay or for your plugins:

```bash
go get github.com/ithildir/liferay-sdk-go/liferay
cd $GOPATH/src/github.com/ithildir/liferay-sdk-go
./gradlew generate
```

Have a look at [this page](https://github.com/liferay/liferay-mobile-sdk/tree/master/builder#properties)
to find the various configuration options available in `gradle.properties`.

## Use

The Liferay Go SDK is an almost line-by-line port of the Android version, so
more detailed information about usage can be found at [this page](https://github.com/Ithildir/liferay-mobile-sdk/tree/master/android#use).

1. Create a `Session` with user credentials:

	```go
	import "github.com/ithildir/liferay-sdk-go/liferay"

	session := liferay.NewSession("http://localhost:8080", "test@liferay.com", "test")
	```

2. Import the service package you need, for example:

	```go
	import "github.com/ithildir/liferay-sdk-go/liferay/service/v62/blogsentry"
	```

3. Create a `Service` object and make a service call:

	```go
	service := blogsentry.NewService(session)

	entries, err := service.GetGroupEntries(10184, 0, 10)

	if err != nil {
		log.Fatal(err)
	}
	```

	Go lacks support for method overloading (here's [why](https://golang.org/doc/faq#overloading)), 
	so overloaded service functions are renamed in `GetGroupEntries2`, 
	`GetGroupEntries3`, etc.

	Compared to my [other](https://github.com/Ithildir/liferay-sdk-builder-windows) 
	SDK, the Liferay Go SDK is a bit more rough in regards to return types, so 
	some type assertions will be required in order to obtain the correct data 
	types:

	```go
	for _, e := range entries {
		e := e.(map[string]interface{})

		log.Println(e["title"])
	}
	```

### Batch

The SDK allows sending requests using batch processing, which can be much more
efficient in some cases. Read [dlapp_test.go](https://github.com/Ithildir/liferay-sdk-go/blob/master/tests/dlapp_test.go)
to find an example on how to create a batch of calls and send them all together.

### Non-primitive arguments

There are some special cases in which service methods arguments are not
primitives. In these cases, you should use `ObjectWrapper`. Read
[orderbycomparator_test.go](https://github.com/Ithildir/liferay-sdk-go/blob/master/tests/orderbycomparator_test.go)
and [servicecontext_test.go](https://github.com/Ithildir/liferay-sdk-go/blob/master/tests/servicecontext_test.go)
to find examples on how to use non-primitive arguments in service methods.

### Binaries

Some Liferay services require argument types such as byte slices (`[]byte`) and
Readers (`io.Reader`). Read [dlapp_test.go](https://github.com/Ithildir/liferay-sdk-go/blob/master/tests/dlapp_test.go)
to find an example on how to send binary data to the portal.

## Thanks

A huge huge huge thanks to [Pier](https://github.com/yuchi/) for the great logo!

## License

This library, *Liferay Go SDK*, is free software ("Licensed Software");
you can redistribute it and/or modify it under the terms of the [GNU Lesser General Public License](http://www.gnu.org/licenses/lgpl-2.1.html) 
as published by the Free Software Foundation; either version 2.1 of the License, 
or (at your option) any later version.

This library is distributed in the hope that it will be useful, but WITHOUT ANY
WARRANTY; including but not limited to, the implied warranty of MERCHANTABILITY,
NONINFRINGEMENT, or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Lesser General
Public License for more details.

You should have received a copy of the [GNU Lesser General Public License](http://www.gnu.org/licenses/lgpl-2.1.html) 
along with this library; if not, write to the Free Software Foundation, Inc., 51 
Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA

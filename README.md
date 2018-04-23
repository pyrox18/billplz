# pyrox18/billplz

A Go wrapper for the [Billplz REST API](https://www.billplz.com/api).

## Installation

Assuming a correctly configured Go toolchain, run

```bash
$ go get -u github.com/pyrox18/billplz
```

This project also uses [dep](https://github.com/golang/dep), so if dep is installed, `dep ensure` can be run to install the dependencies in the project folder.

## Usage

Example:

```go
import "github.com/pyrox18/billplz"

func main() {
  // Initialise a new client
  c, err := billplz.NewClient(nil, "BILLPLZ_API_KEY_HERE", true);

  // Get a set of collections
  collections, err := c.GetCollectionIndex(1, "");

  // Do whatever with the collections
}
```

Refer to the [documentation](https://godoc.org/github.com/pyrox18/billplz) for details on available types and functions.

## Billplz API Version Support

This package makes requests to version 3 of the Billplz REST API. Features in version 4 are not supported.

## License

This package is licensed under the MIT license. See the LICENSE file for details.
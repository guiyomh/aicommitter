# Package conventionalcommit

This package provides a parser for the [Conventional Commits](https://www.conventionalcommits.org) specification.

## Installation

```bash
go get github.com/your-username/aicommiter2/pkg/conventionalcommit
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/your-username/aicommiter2/pkg/conventionalcommit"
)

func main() {
    message := `feat(api): add new endpoint

Detailed description of the changes.

BREAKING CHANGE: This change breaks compatibility with the old API`

    commit, err := conventionalcommit.Parse(message)
    if err != nil {
        panic(err)
    }

    fmt.Printf("Type: %s\n", commit.Type)
    fmt.Printf("Scope: %s\n", commit.Scope)
    fmt.Printf("Description: %s\n", commit.Description)
    fmt.Printf("Breaking: %v\n", commit.Breaking)
    fmt.Printf("Body: %s\n", commit.Body)
    fmt.Printf("Footer: %s\n", commit.Footer)

    // Convert back to string
    fmt.Printf("Complete message: %s\n", commit.String())
}
```

## Structure

The package defines the following types:

### Type

Standard supported commit types:

- `TypeFeat`: New features
- `TypeFix`: Bug fixes
- `TypeDocs`: Documentation changes
- `TypeStyle`: Code style changes
- `TypeRefactor`: Code refactoring
- `TypePerf`: Performance improvements
- `TypeTest`: Adding or modifying tests
- `TypeChore`: Maintenance tasks

### Commit

The `Commit` structure represents a parsed commit message:

```go
type Commit struct {
    Type        Type   // Commit type (feat, fix, etc.)
    Scope       string // Change scope (optional)
    Description string // Short description of the change
    Body        string // Message body (optional)
    Footer      string // Message footer (optional)
    Breaking    bool   // Indicates if this is a breaking change
}
```

## Functions

### Parse

```go
func Parse(message string) (*Commit, error)
```

Parses a commit message and returns a `Commit` structure. Returns an error if the message format is invalid.

### String

```go
func (c *Commit) String() string
```

Converts a `Commit` back to a string formatted according to the Conventional Commits specification.

## Tests

To run the tests:

```bash
go test -v
```

## License

[Your License]

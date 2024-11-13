# sc-interns-2024

The technical take home for 2024 internship applications.

## Getting started

Requires `Go` >= `1.23`

follow the official install instruction: [Golang Installation](https://go.dev/doc/install)

To run the code on your local machine

```
  go run main.go
```

## Folder structure

```
| go.mod
| README.md
| main.go
| folder
    | get_folder.go
    | get_folder_test.go
    | move_folder.go
    | static.go
    | sample.json
```

## Instructions

- This technical assessment consists of 2 components:
- Component 1:

  - within `get_folder.go`.
    - We would like you to read through, and run, the code.
    - Implement `GetAllChildFolders` method in `get_folder.go` that returns all child folders of a given folder.
    - Write up some unit tests in `get_folder_test.go` for all methods in `get_folder.go`.

- Component 2:
  - within `move_folder.go`.
    - Implement `MoveFolder` method in `move_folder.go` that moves a folder from one parent to another. (more details under component 2 section)
    - Write up some unit tests in `move_folder_test.go` for the `MoveFolder` method.

## Path Structure

You are given a hierarchical tree where each node in the tree is represented by a path similar to `ltree` paths in PostgreSQL.

The tree structure is represented as a series of paths, where each path is folder name separated by dots (e.g., `"alpha.bravo.charlie"`). Each name in the path represents a node, and the full path represents that node’s position in the hierarchy.

we use `ltree` path for our site directory structure as well as our documents folder structure within the SC platform. This allow us to easily store and manipulate our folder structure using psql.

## Component 1

You will need to implement the following:

1. A method to get all child folders of a given folder.
2. The method should return a list of all child folders.
3. Implement any necessary error handling (e.g. invalid orgID, invalid paths, etc).

### Example Scenario

```go
folders := [
  {
    name: "alpha",
    path: "alpha",
    orgID: "org1",
  },
  {
    name: "bravo",
    path: "alpha.bravo",
    orgID: "org1",
  },
  {
    name: "charlie",
    path : "alpha.bravo.charlie",
    orgID: "org1",
  },
  {
    name: "delta",
    path: "alpha.delta",
    orgID: "org1",
  },
  {
    name: "echo",
    path: "echo",
    orgID: "org1",
  },
  {
    name: "foxtrot",
    path: "foxtrot",
    orgID: "org2",
  },
]

getAllChildFolders("org1", "alpha")
// Expected output
[
   {
    name: "bravo",
    path: "alpha.bravo",
    orgID: "org1",
  },
  {
    name: "charlie",
    path : "alpha.bravo.charlie",
    orgID: "org1",
  },
  {
    name: "delta",
    path: "alpha.delta",
    orgID: "org1",
  },
]

getAllChildFolders("org1", "bravo")
// Expected output
[
  {
    name: "charlie",
    path : "alpha.bravo.charlie",
    orgID: "org1",
  },
]

getAllChildFolders("org1", "charlie")
// Expected output
[]

getAllChildFolders("org1", "echo")
// Expected output
[]

getAllChildFolders("org1", "invalid_folder")
// Error: Folder does not exist

getAllChildFolders("org1", "foxtrot")
// Error: Folder does not exist in the specified organization
```

## Component 2

You will need to implement the following:

1. A method to move a subtree from one parent node to another, while maintaining the order of the children.
2. The method should return the new folder structure once the move has occurred.
3. Implement any necessary error handling (e.g. invalid paths, moving a node to a child of itself, moving folders to a different orgID, etc).
4. There is no need to persist state, we can assume each method call will be independent of the previous one.

### Example Scenario

```go

folders := [
  {
    name: "alpha",
    path: "alpha",
    orgID: "org1",
  },
  {
    name: "bravo",
    path: "alpha.bravo",
    orgID: "org1",
  },
  {
    name: "charlie",
    path: "alpha.bravo.charlie",
    orgID: "org1",
  },
  {
    name: "delta",
    path: "alpha.delta",
    orgID: "org1",
  },
  {
    name: "echo",
    path: "alpha.delta.echo",
    orgID: "org1",
  },
  {
    name: "foxtrot",
    path: "foxtrot",
    orgID: "org2",
  }
  {
    name: "golf",
    path: "golf",
    orgID: "org1",
  }
]

moveFolder("bravo", "delta")
// Expected output
[
  {
    name: "alpha",
    path: "alpha",
    orgID: "org1",
  },
  {
    name: "bravo",
    path: "alpha.delta.bravo",
    orgID: "org1",
  },
  {
    name: "charlie",
    path: "alpha.delta.bravo.charlie",
    orgID: "org1",
  },
  {
    name: "delta",
    path: "alpha.delta",
    orgID: "org1",
  },
  {
    name: "echo",
    path: "alpha.delta.echo",
    orgID: "org1",
  },
  {
    name: "foxtrot",
    path: "foxtrot",
    orgID: "org2",
  }
  {
    name: "golf",
    path: "golf",
    orgID: "org1",
  }
]

moveFolder("bravo", "golf")
// Expected output
[
  {
    name: "alpha",
    path: "alpha",
    orgID: "org1",
  },
  {
    name: "bravo",
    path: "golf.bravo",
    orgID: "org1",
  },
  {
    name: "charlie",
    path: "golf.bravo.charlie",
    orgID: "org1",
  },
  {
    name: "delta",
    path: "alpha.delta",
    orgID: "org1",
  },
  {
    name: "echo",
    path: "alpha.delta.echo",
    orgID: "org1",
  },
  {
    name: "foxtrot",
    pa th: "foxtrot",
    orgID: "org2",
  },
  {
    name: "golf",
    path: "golf",
    orgID: "org1",
  }
]

moveFolder("bravo", "charlie")
// Error: Cannot move a folder to a child of itself

moveFolder("bravo", "bravo")
// Error: Cannot move a folder to itself

moveFolder("bravo", "foxtrot")
// Error: Cannot move a folder to a different organization

moveFolder("invalid_folder", "delta")
// Error: Source folder does not exist

moveFolder("bravo", "invalid_folder")
// Error: Destination folder does not exist

```

### Sample Data

a pre-populated `sample.json` file is provided for you to use as a sample data. You can use this data to test your implementation. You can also tweak the data to test different scenarios by changing the config within `static.go` and running the code.

Copy and paste the code snippet below into `main.go` and running `go run main.go`.

```go
  package main

  import (
    "github.com/georgechieng-sc/interns-2022/folders"
  )

  func main() {
    res := folders.GenerateData()

    folders.PrettyPrint(res)

    folders.WriteSampleData(res)
  }
```

## FAQ

- Can I use external libraries?
  - Yes, you can use external libraries.
- Can I use a different programming language?
  - No, you must use Go.
- Can I use a different testing framework?
  - Yes, you can use a different testing framework, be prepared to explain why you chose it.
- Can I use a different data structure?
  - Yes, you can use a different data structure.
- Does the folder result in component 2 need to be sorted?
  - No, the order of the folders does not matter.

## Submission

Create a repo in your chosen git repository (make sure it is public so we can access it) and reply with the link to your code. We recommend using GitHub.

## Contact

If you have any questions feel free to contact us at: interns@safetyculture.io

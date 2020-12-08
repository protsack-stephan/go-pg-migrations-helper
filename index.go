package pgmigrations

import (
	"fmt"
	"strings"
)

// IndexUnique making index unique
const IndexUnique = "UNIQUE"

// IndexConcurrently making index concurent
const IndexConcurrently = "CONCURRENTLY"

// IndexMethod what type of index we want to create
type IndexMethod string

// IndexMethodBTree crete B-tree index
const IndexMethodBTree IndexMethod = "BTREE"

// IndexMethodHash create hash index
const IndexMethodHash IndexMethod = "HASH"

// IndexMethodGist creat gist type of index
const IndexMethodGist IndexMethod = "GIST"

// IndexMethodSpGist create spgist type of index
const IndexMethodSpGist IndexMethod = "SPGIST"

// IndexMethodGin create gin type of index
const IndexMethodGin IndexMethod = "GIN"

// Index struct to represent sql index
type Index struct {
	Table        string
	Method       IndexMethod
	Unique       bool
	Concurrently bool
	Columns      []string
}

// Name get index name
func (idx *Index) Name() string {
	unique := strings.ToLower(string(IndexUnique))

	if !idx.Unique {
		unique = fmt.Sprintf("not%s", unique)
	}

	if len(idx.Method) <= 0 {
		idx.Method = IndexMethodBTree
	}

	return fmt.Sprintf("%s_%s_%s_%s",
		idx.Table,
		strings.Join(idx.Columns, "_"),
		unique,
		strings.ToLower(string(idx.Method)))
}

// Create generate index sql
func (idx *Index) Create() string {
	unique := " "

	if idx.Unique {
		unique = fmt.Sprintf(" %s ", IndexUnique)
	}

	return fmt.Sprintf("CREATE%sINDEX%s %s ON %s USING %s (%s);",
		unique,
		idx.concurrently(),
		idx.Name(),
		idx.Table,
		idx.Method,
		strings.Join(idx.Columns, ", "))
}

// Drop generate sql to delete index
func (idx *Index) Drop() string {
	return fmt.Sprintf("DROP INDEX%s %s;", idx.concurrently(), idx.Name())
}

func (idx *Index) concurrently() string {
	concurrently := ""

	if idx.Concurrently {
		concurrently = fmt.Sprintf(" %s", IndexConcurrently)
	}

	return concurrently
}

package pgmigrations

// PartitionBy available types of partitioning
type PartitionBy string

// PartitionByRange create partition by range of values
const PartitionByRange PartitionBy = "RANGE"

// PartitionByHash create partition by hash of one of more values
const PartitionByHash PartitionBy = "HASH"

// PartitionByList create partition list of values
const PartitionByList PartitionBy = "LIST"

// Partition is struct for partitioning configuration
type Partition struct {
	By      PartitionBy
	Columns []string
}

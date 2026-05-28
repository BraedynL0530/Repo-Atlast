package cordinator

import (
	"time"
	"context"
)

type JobType string

const (
	JobTypeTreeSitter  JobType = "treesitter"  // sent to Python bridge
    JobTypeChurn       JobType = "churn"        // handled in Go
    JobTypeDeps        JobType = "deps"
    JobTypeSecrets     JobType = "secrets"
    JobTypeDocs        JobType = "docs"
)

type JobStatus string

const (
	JobStatusPending JobStatus = "pending"
	JobStatusDone	 JobStatus = "done"
	JobStatusFailed  JobStatus = "failed"
	JobStatusSkipped JobStatus = "skipped"
	JobStatusRunning JobStatus = "running"
)

type File struct {
	Path	   String
	Language   String
	SizeBytes  int64
	Hash	   String
	Context    []byte
}

type Job struct {
	Id		 String
	ScanId   String
	Type	 JobType
	File	 File
	Prioity  int
	EnquedAt time.time
}
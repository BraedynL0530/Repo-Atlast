package jobs

import (
	"time"
	"context"
)

type JobType string

const (
	JobTypeTreeSitter  JobType = "treesitter"  // sent to Python bridge
	JobTypeScan		   JobType = "scan"
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
	Path	   string
	Language   string
	SizeBytes  int64
	Hash	   string
	Context    []byte
}

type Job struct {
	Id		 string
	Type	 JobType
	File	 File
	Prioity  int
	EnquedAt time.time
}

type Warning struct {
    Severity string    // "high" | "medium" | "low" 
    Code     string    // machine-readable: "no-tests", "god-file", "secret-found"
    Message  string    // human-readable
    Line     int       // 0 if not line-specific
}

// Result is what comes back from a worker after processing a Job.
type Result struct {
    JobID      string
    JobType    JobType
    FilePath   string
    Status     JobStatus
    Warnings   []Warning
    Metrics    map[string]float64  // "complexity": 14.2, "line_count": 340
    Error      error               // nil on success
    Duration   time.Duration       // how long the worker took
    CacheHit   bool                // true if this came from cache, not real work
}

// ScanRequest is the top-level input — what the CLI hands to the coordinator.
type ScanRequest struct {
    ID          string
    RepoPath    string
    Incremental bool      // if true, skip files that haven't changed since last scan
    JobTypes    []JobType // which analyzers to run; nil = run all
    Ctx         context.Context
}

// ScanResult is the aggregated output of an entire scan.
type ScanResult struct {
    ScanID      string
    RepoPath    string
    StartedAt   time.Time
    FinishedAt  time.Time
    Results     []Result
    HealthScore float64    // 0–100, calculated by metrics/scoring.go
    TotalFiles  int
    SkippedFiles int       // cache hits
    FailedFiles  int
}
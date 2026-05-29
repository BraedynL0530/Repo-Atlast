package collector

import (
    "errors"
    "io/fs"
    "os"
    "path/filepath"
	"crypto/sha256"
)

// FileInfo is a simple struct — just what we need for now
type FileInfo struct {
    Path      string
    SizeBytes int64
	Hash	  string
}

func FindRepoRoot() (string, error) {
    dir, err := os.Getwd()
    if err != nil {
        return "", err
    }

    for {
        _, err := os.Stat(filepath.Join(dir, ".git"))
        if err == nil {
            return dir, nil // found it
        }

        parent := filepath.Dir(dir)
        if parent == dir {
            // hit the filesystem root, no .git found
            return "", errors.New("not a git repository")
        }
        dir = parent
    }
}

func hashFile(path string) (string, error){
	f, err := os.Open(path)
    if err != nil {
        return "", err
    }
    defer f.Close()

    h := sha256.New()
    if _, err := io.Copy(h, f); err != nil {
        return "", err
    }

    return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func Walk() ([]FileInfo, error) {
    root, err := FindRepoRoot()
    if err != nil {
        return nil, err 
    }

    var files []FileInfo

    err = filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
        if err != nil {
            return nil 
        }

        if d.IsDir() {
            return nil 
        }

        info, err := d.Info()
        if err != nil {
            return nil 
        }

		hash, err := hashFile(path)

		if err != nil {
			hash = ""
		}

		files = append(files, FileInfo{
			Path:      path,
			SizeBytes: info.Size(),
			Hash:      hash,
		})

        files = append(files, FileInfo{
            Path:      path,
            SizeBytes: info.Size(),
        })

        return nil
    })

    return files, err
}
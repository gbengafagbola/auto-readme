package main

import (
    "fmt"
    "os"
    "path/filepath"
)

func main() {
    repoPath := "https://github.com/gbengafagbola/auto-readme" 
    err := filepath.Walk(repoPath, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if info.IsDir() {
            fmt.Println("Directory:", path)
        } else {
            fmt.Println("File:", path)
        }
        return nil
    })
    if err != nil {
        fmt.Printf("Error walking the path: %v\n", err)
    }
}

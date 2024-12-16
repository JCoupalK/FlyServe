package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	var portStr, filePath, dirPath, username, password string
	var autoResolveHTML bool

	flag.StringVar(&portStr, "port", "8080", "Port to serve on")
	flag.StringVar(&portStr, "p", "8080", "Port to serve on (shorthand)")
	flag.StringVar(&filePath, "file", "", "Specific file to serve")
	flag.StringVar(&filePath, "f", "", "Specific file to serve (shorthand)")
	flag.StringVar(&dirPath, "directory", ".", "Directory to serve files from")
	flag.StringVar(&dirPath, "d", ".", "Directory to serve files from (shorthand)")
	flag.StringVar(&username, "username", "", "Username for basic authentication")
	flag.StringVar(&username, "u", "", "Username for basic authentication (shorthand)")
	flag.StringVar(&password, "password", "", "Password for basic authentication")
	flag.StringVar(&password, "pw", "", "Password for basic authentication (shorthand)")
	flag.BoolVar(&autoResolveHTML, "html", false, "Enable auto-resolution of .html files")
	flag.BoolVar(&autoResolveHTML, "h", false, "Enable auto-resolution of .html files (shorthand)")

	// Usage function
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  -p, --port    \tPort to serve on (Default is 8080)\n")
		fmt.Fprintf(os.Stderr, "  -f, --file    \tSpecific file to serve\n")
		fmt.Fprintf(
			os.Stderr,
			"  -d, --directory \tDirectory to serve files from (Default is current directory)\n",
		)
		fmt.Fprintf(os.Stderr, "  -u, --username \tUsername for basic authentication\n")
		fmt.Fprintf(os.Stderr, "  -pw, --password \tPassword for basic authentication\n")
		fmt.Fprintf(os.Stderr, "  -h, --html    \tEnable auto-resolution of .html files\n")
	}

	flag.Parse()

	port, err := strconv.Atoi(portStr)
	if err != nil {
		fmt.Println("Invalid port number")
		os.Exit(1)
	}

	// Check if the directory exists and is a directory
	dirInfo, err := os.Stat(dirPath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("Specified directory does not exist: %s\n", dirPath)
		} else {
			fmt.Printf("Error checking the specified directory: %s\n", err)
		}
		os.Exit(1)
	}
	if !dirInfo.IsDir() {
		fmt.Printf("Specified path is not a directory: %s\n", dirPath)
		os.Exit(1)
	}

	http.HandleFunc("/", basicAuth(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)

		if filePath != "" {
			http.ServeFile(w, r, filePath)
			return
		}

		fullPath := filepath.Join(dirPath, filepath.Clean(r.URL.Path))

		if autoResolveHTML {
			// Check if the file without .html exists
			if _, err := os.Stat(fullPath); os.IsNotExist(err) {
				// Try adding .html
				fullPathWithHTML := fullPath + ".html"
				if _, err := os.Stat(fullPathWithHTML); err == nil {
					fullPath = fullPathWithHTML
				}
			}
		}

		// Serve index.html if the request is for a directory that contains it
		if fileInfo, err := os.Stat(fullPath); err == nil {
			if fileInfo.IsDir() {
				indexPath := filepath.Join(fullPath, "index.html")
				if _, err := os.Stat(indexPath); err == nil {
					http.ServeFile(w, r, indexPath)
					return
				}

				// If no index.html, generate directory listing
				serveDirectoryWithCustomText(w, r.URL.Path, fullPath)
				return
			}
		}

		http.ServeFile(w, r, fullPath)
	}, username, password))

	// Start server
	address := fmt.Sprintf(":%d", port)
	log.Printf("Serving HTTP on %s ...", address)
	err = http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatal(err)
	}
}


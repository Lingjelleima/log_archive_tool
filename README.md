Go Log Archiver

This project is a utility built for learning Go (Golang) with a focus on DevOps automation. It automates the process of managing old log files to save disk space.

Functionality
- Scans a target directory for .log files.
- Filters files based on a specific age (e.g., older than 3 days).
- Compresses identified files into a single ZIP archive.
- Deletes the original log files after successful archival.

Prerequisites
- Go 1.20 or later installed on your system.

Installation
1. Clone this repository:
   git clone https://github.com

2. Navigate to the project directory:
   cd log-archiver

Usage
1. Place log files in the /logs directory.
2. Run the script using the Go toolchain:
   go run main.go

Configuration
The following variables can be modified in main.go to fit your environment:
- sourceDir: The directory to scan for logs.
- archiveDir: The directory where ZIP files will be stored.
- daysOld: The age threshold for archiving files.

Building the Binary
To create a standalone executable for deployment:
go build -o log-archiver

https://roadmap.sh/projects/log-archive-tool

License
MIT License

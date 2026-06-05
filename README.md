# Log Archive Tool

This is a CLI tool built in Go for archiving logs. It is a project designed to practice Golang for DevOps automation.

# Functionality
- Accepts a directory path as a command-line argument.
- Compresses all files in the directory into a `.tar.gz` archive.
- Names archives with a unique timestamp (`logs_archive_YYYYMMDD.tar.gz`).
- Records an audit log of every archival action in `archive_log.txt`.

# Prerequisites
- Go 1.20 or later.

# Installation & Setup
1. Clone the repository:
   git clone https://github.com
   cd log-archive-tool

2. Build the binary:
   go build -o log-archive

# Usage
Run the tool by providing the path to the logs directory:
./log-archive /path/to/logs

# License
MIT License

https://roadmap.sh/projects/log-archive-tool

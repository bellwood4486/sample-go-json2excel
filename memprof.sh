#!/usr/bin/env bash
# Usage: memprof.sh <case number>

casenum=$1
prof_file="mem${casenum}.prof"
prof_png_file="mem${casenum}.prof.png"

echo "Run case${casenum} with memory profiling"
go run ./cmd/parsetest/main.go -case "${casenum}" -memprofile "${prof_file}"
go tool pprof -png "${prof_file}" >  "${prof_png_file}"
echo "Done"

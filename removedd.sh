#!/bin/bash
declare -A arr

# shopt -s globstar enables to use **
shopt -s globstar 

# ** will match all files and directories in the current directory and subdirectories
for file in **; do 
# if file is exists continue
  [[ -f "$file" ]] || continue
  # checksum file from file
  read cksm _ < <(md5sum "$file")
  # if cksm is in arr already then remove that file
  if ((arr[$cksm]++)); then 
    rm $file
  fi
done

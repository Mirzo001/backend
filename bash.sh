#! /usr/bin/bash
while IFS="" read -r p || [ -n "$p" ]
do
  printf '%s\n' "$p"
done < output.txt


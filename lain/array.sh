#!/bin/sh
# process lines from plaintext files into a format suitable for golang arrays
# file path is relative to where the script is
mkdir -p processed
sed 's/^/   "/; s/$/",/' original/$FILE | sort > processed/$FILE.txt 

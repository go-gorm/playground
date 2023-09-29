#!/bin/sh

rm -rf testpackage*

# Loop from 1 to 100
for i in $(seq 1 10); do
  package="package$i"
  dir="testpackage/$package"
  mkdir -p "$dir"
  dirEscaped=$(echo "$dir" | sed 's/\//\\\//g')
  cat main_test.go | sed "s/main/$package/g" > "$dir/test_test.go"
done
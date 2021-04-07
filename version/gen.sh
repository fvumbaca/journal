#!/bin/bash

OUTPUT="version_gen.go"

echo "// Generated. DO NOT EDIT" > $OUTPUT
echo >> $OUTPUT
echo "package version" >> $OUTPUT
echo >> $OUTPUT
echo "const Version = \"$(git rev-parse HEAD | cut -c 1-8)\"" >> $OUTPUT
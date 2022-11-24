#! /usr/bin/env python3

"""
If you have used github.com/edsu/anon in the past, you can use this script
to convert the IP ranges into the new format for this tool.
If your ip ranges were previously inline in the general config file of anon, please move the list to a sepperate json
file first, and then call this script with.

First check that the output looks correct and that there are no errors with:
$ python convertRanges.py old.json

Then save the new content to a file:
$ python convertRanges.py old.json > new.json

"""

import json, sys

if len(sys.argv) != 2 or sys.argv[1].endswith == ".json":
    print("Please provide filename of old json file as first argument")
    sys.exit(1)

filename = sys.argv[1]

newRanges = []

with open(filename) as fh:
    data = json.loads(fh.read())

for k, v in data.items():
    newRanges.append({"name": k, "ranges": v})

print(json.dumps(newRanges, indent=4))

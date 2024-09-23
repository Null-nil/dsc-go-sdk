#!/usr/bin/python3

import json

doc = json.load(open("documentation.json","r"))
json.dump(doc, open("doc_pretty.json","w"), indent=4)

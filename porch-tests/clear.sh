#! /bin/bash

porchctl rpkg list -A | egrep '( bp | depl )'| awk '{printf("porchctl rpkg propose-delete %s -n porch-demo\n", $2)}' > ~/tmp/aaa
source ~/tmp/aaa
porchctl rpkg list -A | egrep '( bp | depl )'| awk '{printf("porchctl rpkg delete %s -n porch-demo\n", $2)}' > ~/tmp/aaa
source ~/tmp/aaa

find . -type d -name 'bp*' -exec rm -fr {} \;
find . -type d -name 'depl*' -exec rm -fr {} \;

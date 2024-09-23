#!/opt/homebrew/bin/bash

. /tmp/tko-python-env/bin/activate

$HOME/go/bin/tko-data start --name=Liam --backend=postgresql --backend-clean=false

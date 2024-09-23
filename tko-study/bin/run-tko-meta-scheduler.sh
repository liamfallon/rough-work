#!/opt/homebrew/bin/bash

. /tmp/tko-python-env/bin/activate

$HOME/go/bin/tko-meta-scheduler --verbose=10 start

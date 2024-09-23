#!/usr/bin/env python3

# Allow relative imports
import sys, pathlib
sys.path.append('/Users/liam/git/github/nephio-experimental/tko/sdk/python')

import tko

def prepare():
  target_resource = tko.get_target_resource()
  prep_log.write(str(target_resource))
  prep_log.write("prep2\n")

  #if smf is None:
  #  return True
  #if tko.GVK(resource=smf) != smf_gvk:
  #  return True

  #smf['status'] = smf.get('status', {})
  #smf['status']['test'] = 'hi'
  #tko.set_prepared(smf)
  return True

if __name__ == '__main__':
  prep_log = open('/tmp/tko.log', 'a')
  prep_log.write("prep\n")

  tko.prepare(prepare)

  prep_log.close()

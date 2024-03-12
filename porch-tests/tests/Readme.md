
# rpkg clone strategy

Defines which strategy should be used to update the package. It defaults to `resource-merge`.

* resource-merge: Perform a structural comparison of the original / updated resources, and merge the changes into the local package.
* fast-forward: Fail without updating if the local package was modified since it was fetched.
* force-delete-replace: Wipe all the local changes to the package and replace it with the remote version.

# rpkg copy replay-strategy

A boolean flag, defaults to `false`.

When `true`, all the tasks from the source PackageRevision are copied to the target PackageRevision
When `false`,  creates an `Edit` task that does a deep copy of the source PackageRevision into the target PackageRevision
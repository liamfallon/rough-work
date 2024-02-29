# Create blueprints

## Create blueprint v1 (replay strategy true)
```
porchctl rpkg init bp-rs-t --workspace v1 --repository management -n porch-demo

porchctl rpkg pull management-4d1c7c0c6d1e72852b5d70062b3b018da5ee39c8 bp-rs-t-v1 -n porch-demo

cp demo-app-1-v1-rs-t.yaml bp-rs-t-v1/demo-app-1.yaml

porchctl rpkg push management-4d1c7c0c6d1e72852b5d70062b3b018da5ee39c8 bp-rs-t-v1 -n porch-demo

porchctl rpkg propose management-4d1c7c0c6d1e72852b5d70062b3b018da5ee39c8 -n porch-demo
porchctl rpkg approve management-4d1c7c0c6d1e72852b5d70062b3b018da5ee39c8 -n porch-demo
```

## Create blueprint v1 (replay strategy false)
```
porchctl rpkg init bp-rs-f --workspace v1 --repository management -n porch-demo

porchctl rpkg pull management-82aaf44b911ab8f1be1b1097364954cb6e03811a bp-rs-f-v1 -n porch-demo

cp demo-app-1-v1-rs-f.yaml bp-rs-f-v1/demo-app-1.yaml

porchctl rpkg push management-82aaf44b911ab8f1be1b1097364954cb6e03811a bp-rs-f-v1 -n porch-demo

porchctl rpkg propose management-82aaf44b911ab8f1be1b1097364954cb6e03811a -n porch-demo
porchctl rpkg approve management-82aaf44b911ab8f1be1b1097364954cb6e03811a -n porch-demo
```
## Create blueprint v2 (from v1 with replay-strategy true)
```
porchctl rpkg copy management-4d1c7c0c6d1e72852b5d70062b3b018da5ee39c8  --workspace v2 --replay-strategy=true -n porch-demo

porchctl rpkg pull management-973980d978aa3af094ee0da738f201512dc74d5d bp-rs-t-v2 -n porch-demo

cp demo-app-1-v2-rs-t.yaml bp-rs-t-v2/demo-app-1.yaml

porchctl rpkg push management-973980d978aa3af094ee0da738f201512dc74d5d bp-rs-t-v2 -n porch-demo

porchctl rpkg propose management-973980d978aa3af094ee0da738f201512dc74d5d -n porch-demo
porchctl rpkg approve management-973980d978aa3af094ee0da738f201512dc74d5d -n porch-demo
```
## Create blueprint v2 (from v1 with replay-strategy false)
```
porchctl rpkg copy management-82aaf44b911ab8f1be1b1097364954cb6e03811a  --workspace v2 --replay-strategy=false -n porch-demo

porchctl rpkg pull management-dcb5500b679bf5d1618b5728cb27b44a5a242233 bp-rs-f-v2 -n porch-demo

cp demo-app-1-v2-rs-f.yaml bp-rs-f-v2/demo-app-1.yaml

porchctl rpkg push management-dcb5500b679bf5d1618b5728cb27b44a5a242233 bp-rs-f-v2 -n porch-demo

porchctl rpkg propose management-dcb5500b679bf5d1618b5728cb27b44a5a242233 -n porch-demo
porchctl rpkg approve management-dcb5500b679bf5d1618b5728cb27b44a5a242233 -n porch-demo
```
# Create deployments

## Create deployment v1 (replay strategy true, clone strategy resource-merge)
```
porchctl rpkg clone management-4d1c7c0c6d1e72852b5d70062b3b018da5ee39c8 depl-rs-t-s-rm --repository edge1 --workspace v1 --strategy resource-merge -n porch-demo

porchctl rpkg pull edge1-d98fe530dfe9274a28ffb35f182abf74454e439c depl-rs-t-s-rm-v1 -n porch-demo

cp demo-app-1-v1-rs-t-depl.yaml depl-rs-t-s-rm-v1/demo-app-1.yaml

porchctl rpkg push edge1-d98fe530dfe9274a28ffb35f182abf74454e439c depl-rs-t-s-rm-v1 -n porch-demo

porchctl rpkg propose edge1-d98fe530dfe9274a28ffb35f182abf74454e439c -n porch-demo
porchctl rpkg approve edge1-d98fe530dfe9274a28ffb35f182abf74454e439c -n porch-demo
```

## Create deployment v1 (replay strategy true, clone strategy fast-forward)
```
porchctl rpkg clone management-4d1c7c0c6d1e72852b5d70062b3b018da5ee39c8 depl-rs-t-s-ff --repository edge1 --workspace v1 --strategy fast-forward -n porch-demo

porchctl rpkg pull edge1-e0f12a5eeef1b3c0f0c22340f6a8ccb7472aa06c depl-rs-t-s-ff-v1 -n porch-demo

cp demo-app-1-v1-rs-t-depl.yaml depl-rs-t-s-ff-v1/demo-app-1.yaml

porchctl rpkg push edge1-e0f12a5eeef1b3c0f0c22340f6a8ccb7472aa06c depl-rs-t-s-ff-v1 -n porch-demo

porchctl rpkg propose edge1-e0f12a5eeef1b3c0f0c22340f6a8ccb7472aa06c -n porch-demo
porchctl rpkg approve edge1-e0f12a5eeef1b3c0f0c22340f6a8ccb7472aa06c -n porch-demo
```
## Create deployment v1 (replay strategy true, clone strategy force-delete-replace)
```
porchctl rpkg clone management-4d1c7c0c6d1e72852b5d70062b3b018da5ee39c8 depl-rs-t-s-fdr --repository edge1 --workspace v1 --strategy force-delete-replace -n porch-demo

porchctl rpkg pull edge1-1ed6e7a2b88e4e87604717a5b8fbe661480a5247 depl-rs-t-s-fdr-v1 -n porch-demo

cp demo-app-1-v1-rs-t-depl.yaml depl-rs-t-s-fdr-v1/demo-app-1.yaml

porchctl rpkg push edge1-1ed6e7a2b88e4e87604717a5b8fbe661480a5247 depl-rs-t-s-fdr-v1 -n porch-demo

porchctl rpkg propose edge1-1ed6e7a2b88e4e87604717a5b8fbe661480a5247 -n porch-demo
porchctl rpkg approve edge1-1ed6e7a2b88e4e87604717a5b8fbe661480a5247 -n porch-demo
```
## Create deployment v1 (replay strategy false, clone strategy resource-merge)
```
porchctl rpkg clone management-82aaf44b911ab8f1be1b1097364954cb6e03811a depl-rs-f-s-rm --repository edge1 --workspace v1 --strategy resource-merge -n porch-demo

porchctl rpkg pull edge1-62b7d3832aa2b892400a1996ada8bb92bd527e0b depl-rs-f-s-rm-v1 -n porch-demo

cp demo-app-1-v1-rs-f-depl.yaml depl-rs-f-s-rm-v1/demo-app-1.yaml

porchctl rpkg push edge1-62b7d3832aa2b892400a1996ada8bb92bd527e0b depl-rs-f-s-rm-v1 -n porch-demo

porchctl rpkg propose edge1-62b7d3832aa2b892400a1996ada8bb92bd527e0b -n porch-demo
porchctl rpkg approve edge1-62b7d3832aa2b892400a1996ada8bb92bd527e0b -n porch-demo
```
## Create deployment v1 (replay strategy false, clone strategy fast-forward)
```
porchctl rpkg clone management-82aaf44b911ab8f1be1b1097364954cb6e03811a depl-rs-f-s-ff --repository edge1 --workspace v1 --strategy fast-forward -n porch-demo

porchctl rpkg pull edge1-af025b20977051b3f94a9901e1bbd2116494ba62 depl-rs-f-s-ff-v1 -n porch-demo

cp demo-app-1-v1-rs-f-depl.yaml depl-rs-f-s-ff-v1/demo-app-1.yaml

porchctl rpkg push edge1-af025b20977051b3f94a9901e1bbd2116494ba62 depl-rs-f-s-ff-v1 -n porch-demo

porchctl rpkg propose edge1-af025b20977051b3f94a9901e1bbd2116494ba62 -n porch-demo
porchctl rpkg approve edge1-af025b20977051b3f94a9901e1bbd2116494ba62 -n porch-demo
```
## Create deployment v1 (replay strategy false, clone strategy force-delete-replace)
```
porchctl rpkg clone management-82aaf44b911ab8f1be1b1097364954cb6e03811a depl-rs-f-s-fdr --repository edge1 --workspace v1 --strategy force-delete-replace -n porch-demo

porchctl rpkg pull edge1-d6498e9255625b1e41316d3d7be9f5ac75c5b5bb depl-rs-f-s-fdr-v1 -n porch-demo

cp demo-app-1-v1-rs-f-depl.yaml depl-rs-f-s-fdr-v1/demo-app-1.yaml

porchctl rpkg push edge1-d6498e9255625b1e41316d3d7be9f5ac75c5b5bb depl-rs-f-s-fdr-v1 -n porch-demo

porchctl rpkg propose edge1-d6498e9255625b1e41316d3d7be9f5ac75c5b5bb -n porch-demo
porchctl rpkg approve edge1-d6498e9255625b1e41316d3d7be9f5ac75c5b5bb -n porch-demo
```
# Upgrade deployments

## Upgrade deployment v1 (replay strategy true, clone strategy resource-merge)
```
porchctl rpkg copy edge1-d98fe530dfe9274a28ffb35f182abf74454e439c -n porch-demo --replay-strategy=true --workspace=v2
porchctl rpkg update edge1-615e939914dc8d7c4a9f7ed981bbfe536dc78d14 --revision=v2

porchctl rpkg propose edge1-615e939914dc8d7c4a9f7ed981bbfe536dc78d14 -n porch-demo
porchctl rpkg approve edge1-615e939914dc8d7c4a9f7ed981bbfe536dc78d14 -n porch-demo
```
This worked fine.
```
porchctl rpkg copy edge1-d98fe530dfe9274a28ffb35f182abf74454e439c -n porch-demo --replay-strategy=false --workspace=v3
porchctl rpkg update edge1-b7b2ac103db316143a147ffb05d02ad2772bd0ee --revision=v2

Error: upstream source not found for package rev "depl-rs-t-s-rm"; only cloned packages can be updated 
```
## Upgrade deployment v1 (replay strategy true, clone strategy fast-forward)
```
porchctl rpkg copy edge1-e0f12a5eeef1b3c0f0c22340f6a8ccb7472aa06c -n porch-demo --replay-strategy=true --workspace=v2
porchctl rpkg update edge1-32a5d57703f9781d48674a258ac54b77f25e8482 --revision=v2

porchctl rpkg propose edge1-32a5d57703f9781d48674a258ac54b77f25e8482 -n porch-demo
porchctl rpkg approve edge1-32a5d57703f9781d48674a258ac54b77f25e8482 -n porch-demo
```
This worked fine
```
porchctl rpkg copy edge1-e0f12a5eeef1b3c0f0c22340f6a8ccb7472aa06c -n porch-demo --replay-strategy=false --workspace=v3
porchctl rpkg update edge1-0cc4d310944fc83feb857a10ecde3ebbdc2c60a7 --revision=v2

Error: upstream source not found for package rev "depl-rs-t-s-ff"; only cloned packages can be updated 
```
## Upgrade deployment v1 (replay strategy true, clone strategy force-delete-replace)
```
porchctl rpkg copy edge1-1ed6e7a2b88e4e87604717a5b8fbe661480a5247 -n porch-demo --replay-strategy=true --workspace=v2
porchctl rpkg update edge1-aba73fe5aa678b457ae0213b511ac147041a46f4 --revision=v2

porchctl rpkg propose edge1-aba73fe5aa678b457ae0213b511ac147041a46f4 -n porch-demo
porchctl rpkg approve edge1-aba73fe5aa678b457ae0213b511ac147041a46f4 -n porch-demo
```
This worked fine
```
porchctl rpkg copy edge1-1ed6e7a2b88e4e87604717a5b8fbe661480a5247 -n porch-demo --replay-strategy=false --workspace=v3
porchctl rpkg update edge1-96c6553aab93cc5b0c863a231efa4f37309c840a --revision=v3

Error: upstream source not found for package rev "depl-rs-t-s-fdr"; only cloned packages can be updated 
```
## Upgrade deployment v1 (replay strategy false, clone strategy resource-merge)
```
porchctl rpkg copy edge1-62b7d3832aa2b892400a1996ada8bb92bd527e0b -n porch-demo --replay-strategy=true --workspace=v2
porchctl rpkg update edge1-b41d6ad8584400ff5b87f7e7c0452be429efadba --revision=v2

porchctl rpkg propose edge1-b41d6ad8584400ff5b87f7e7c0452be429efadba -n porch-demo
porchctl rpkg approve edge1-b41d6ad8584400ff5b87f7e7c0452be429efadba -n porch-demo
```
This worked fine
```
porchctl rpkg copy edge1-62b7d3832aa2b892400a1996ada8bb92bd527e0b -n porch-demo --replay-strategy=false --workspace=v3
porchctl rpkg update edge1-141740c2030b17417017898de18f4025f8c04a54 --revision=v2

Error: upstream source not found for package rev "depl-rs-f-s-rm"; only cloned packages can be updated 
```
## Upgrade deployment v1 (replay strategy false, clone strategy fast-forward)
```
porchctl rpkg copy edge1-af025b20977051b3f94a9901e1bbd2116494ba62 -n porch-demo --replay-strategy=true --workspace=v2
porchctl rpkg update edge1-b9dba15120f2a130b0bf31547c08f07d2f5fe55f --revision=v2

porchctl rpkg propose edge1-b9dba15120f2a130b0bf31547c08f07d2f5fe55f -n porch-demo
porchctl rpkg approve edge1-b9dba15120f2a130b0bf31547c08f07d2f5fe55f -n porch-demo
```
This worked fine
```
porchctl rpkg copy edge1-af025b20977051b3f94a9901e1bbd2116494ba62 -n porch-demo --replay-strategy=false --workspace=v3
porchctl rpkg update edge1-39d58e9c421498296dcecaddd7af3f8644870b5b --revision=v2

Error: upstream source not found for package rev "depl-rs-f-s-ff"; only cloned packages can be updated 
```
## Update deployment v1 (replay strategy false, clone strategy force-delete-replace)
```
porchctl rpkg copy edge1-d6498e9255625b1e41316d3d7be9f5ac75c5b5bb -n porch-demo --replay-strategy=true --workspace=v2
porchctl rpkg update edge1-fb4177a78032bcba5b183d55c5f3bf70e3d62a1d --revision=v2

porchctl rpkg propose edge1-fb4177a78032bcba5b183d55c5f3bf70e3d62a1d -n porch-demo
porchctl rpkg approve edge1-fb4177a78032bcba5b183d55c5f3bf70e3d62a1d -n porch-demo
```
This worked fine
```
porchctl rpkg copy edge1-d6498e9255625b1e41316d3d7be9f5ac75c5b5bb -n porch-demo --replay-strategy=false --workspace=v3
porchctl rpkg update edge1-1b32709a7f4fcd1e14d689489a1c2269c87b6591 --revision=v3

Error: upstream source not found for package rev "depl-rs-f-s-fdr"; only cloned packages can be updated 
```

# Reproduce bug

## Create blueprint v1
```
porchctl rpkg init bp --workspace v1 --repository management -n porch-demo

porchctl rpkg pull management-08f3186a5068f90c7873cf2a9b6c56fdc64dcac9 bp-v1 -n porch-demo

cp demo-app-1-bp-v1.yaml bp-v1/demo-app-1.yaml

porchctl rpkg push management-08f3186a5068f90c7873cf2a9b6c56fdc64dcac9 bp-v1 -n porch-demo

porchctl rpkg propose management-08f3186a5068f90c7873cf2a9b6c56fdc64dcac9 -n porch-demo
porchctl rpkg approve management-08f3186a5068f90c7873cf2a9b6c56fdc64dcac9 -n porch-demo
```
## Create blueprint v2
```
porchctl rpkg copy management-08f3186a5068f90c7873cf2a9b6c56fdc64dcac9  --workspace v2 -n porch-demo

porchctl rpkg pull management-8f9d314ce0ed3718910880812c193fd041771e1a bp-v2 -n porch-demo

cp demo-app-1-bp-v2-value-fail.yaml bp-v2/demo-app-1.yaml

porchctl rpkg push management-8f9d314ce0ed3718910880812c193fd041771e1a bp-v2 -n porch-demo

porchctl rpkg propose management-8f9d314ce0ed3718910880812c193fd041771e1a -n porch-demo
porchctl rpkg approve management-8f9d314ce0ed3718910880812c193fd041771e1a -n porch-demo
```
## Create deployment v1
```
porchctl rpkg clone management-08f3186a5068f90c7873cf2a9b6c56fdc64dcac9 depl --repository edge1 --workspace v1 -n porch-demo

porchctl rpkg pull edge1-721f9963e8b7461f5a53aa614e03baded3de791a depl-v1 -n porch-demo

cp demo-app-1-depl.yaml depl-v1/demo-app-1.yaml

porchctl rpkg push edge1-721f9963e8b7461f5a53aa614e03baded3de791a depl-v1 -n porch-demo

porchctl rpkg propose edge1-721f9963e8b7461f5a53aa614e03baded3de791a -n porch-demo
porchctl rpkg approve edge1-721f9963e8b7461f5a53aa614e03baded3de791a -n porch-demo
```
## Upgrade deployment v1
```
porchctl rpkg copy edge1-721f9963e8b7461f5a53aa614e03baded3de791a -n porch-demo --replay-strategy=true --workspace=v2
porchctl rpkg update edge1-1ebfa3e4ef95d04744d5069743cd350d910d743e --revision=v2

Error: Internal error occurred: error applying patch: conflict: fragment line does not match src line 
```

```
porchctl rpkg copy edge1-721f9963e8b7461f5a53aa614e03baded3de791a -n porch-demo --replay-strategy=false --workspace=v2
porchctl rpkg update edge1-1ebfa3e4ef95d04744d5069743cd350d910d743e --revision=v2

Error: upstream source not found for package rev "depl"; only cloned packages can be updated 
```

# Bug scenario with no blueprint change 

## Create blueprint v1
```
porchctl rpkg init bp --workspace v1 --repository management -n porch-demo

porchctl rpkg pull management-08f3186a5068f90c7873cf2a9b6c56fdc64dcac9 bp-v1 -n porch-demo

cp demo-app-1-bp-v1.yaml bp-v1/demo-app-1.yaml

porchctl rpkg push management-08f3186a5068f90c7873cf2a9b6c56fdc64dcac9 bp-v1 -n porch-demo

porchctl rpkg propose management-08f3186a5068f90c7873cf2a9b6c56fdc64dcac9 -n porch-demo
porchctl rpkg approve management-08f3186a5068f90c7873cf2a9b6c56fdc64dcac9 -n porch-demo
```
## Create blueprint v2
```
porchctl rpkg copy management-08f3186a5068f90c7873cf2a9b6c56fdc64dcac9  --workspace v2 -n porch-demo

porchctl rpkg pull management-8f9d314ce0ed3718910880812c193fd041771e1a bp-v2 -n porch-demo
porchctl rpkg propose management-8f9d314ce0ed3718910880812c193fd041771e1a -n porch-demo
porchctl rpkg approve management-8f9d314ce0ed3718910880812c193fd041771e1a -n porch-demo
```
## Create deployment v1
```
porchctl rpkg clone management-08f3186a5068f90c7873cf2a9b6c56fdc64dcac9 depl --repository edge1 --workspace v1 -n porch-demo

porchctl rpkg pull edge1-721f9963e8b7461f5a53aa614e03baded3de791a depl-v1 -n porch-demo

cp demo-app-1-depl.yaml depl-v1/demo-app-1.yaml

porchctl rpkg push edge1-721f9963e8b7461f5a53aa614e03baded3de791a depl-v1 -n porch-demo

porchctl rpkg propose edge1-721f9963e8b7461f5a53aa614e03baded3de791a -n porch-demo
porchctl rpkg approve edge1-721f9963e8b7461f5a53aa614e03baded3de791a -n porch-demo
```
## Upgrade deployment v1
```
porchctl rpkg copy edge1-721f9963e8b7461f5a53aa614e03baded3de791a -n porch-demo --replay-strategy=true --workspace=v2
porchctl rpkg update edge1-1ebfa3e4ef95d04744d5069743cd350d910d743e --revision=v2
```
This works fine.

# Bug scenario with blueprint parameter change (fails)

## Create blueprint v1
```
porchctl rpkg init bp --workspace v1 --repository management -n porch-demo

porchctl rpkg pull management-08f3186a5068f90c7873cf2a9b6c56fdc64dcac9 bp-v1 -n porch-demo

cp demo-app-1-bp-v1.yaml bp-v1/demo-app-1.yaml

porchctl rpkg push management-08f3186a5068f90c7873cf2a9b6c56fdc64dcac9 bp-v1 -n porch-demo

porchctl rpkg propose management-08f3186a5068f90c7873cf2a9b6c56fdc64dcac9 -n porch-demo
porchctl rpkg approve management-08f3186a5068f90c7873cf2a9b6c56fdc64dcac9 -n porch-demo
```
## Create blueprint v2
```
porchctl rpkg copy management-08f3186a5068f90c7873cf2a9b6c56fdc64dcac9  --workspace v2 -n porch-demo

porchctl rpkg pull management-8f9d314ce0ed3718910880812c193fd041771e1a bp-v2 -n porch-demo

cp demo-app-1-bp-v2-new-par-fail.yaml bp-v2/demo-app-1.yaml

porchctl rpkg push management-8f9d314ce0ed3718910880812c193fd041771e1a bp-v2 -n porch-demo

porchctl rpkg propose management-8f9d314ce0ed3718910880812c193fd041771e1a -n porch-demo
porchctl rpkg approve management-8f9d314ce0ed3718910880812c193fd041771e1a -n porch-demo
```
## Create deployment v1
```
porchctl rpkg clone management-08f3186a5068f90c7873cf2a9b6c56fdc64dcac9 depl --repository edge1 --workspace v1 -n porch-demo

porchctl rpkg pull edge1-721f9963e8b7461f5a53aa614e03baded3de791a depl-v1 -n porch-demo

cp demo-app-1-depl.yaml depl-v1/demo-app-1.yaml

porchctl rpkg push edge1-721f9963e8b7461f5a53aa614e03baded3de791a depl-v1 -n porch-demo

porchctl rpkg propose edge1-721f9963e8b7461f5a53aa614e03baded3de791a -n porch-demo
porchctl rpkg approve edge1-721f9963e8b7461f5a53aa614e03baded3de791a -n porch-demo
```
## Upgrade deployment v1
```
porchctl rpkg copy edge1-721f9963e8b7461f5a53aa614e03baded3de791a -n porch-demo --replay-strategy=true --workspace=v2
porchctl rpkg update edge1-1ebfa3e4ef95d04744d5069743cd350d910d743e --revision=v2
Error: Internal error occurred: error applying patch: conflict: fragment line does not match src line 
```

# Bug scenario with blueprint parameter change (works)

## Create blueprint v1
```
porchctl rpkg init bp --workspace v1 --repository management -n porch-demo

porchctl rpkg pull management-08f3186a5068f90c7873cf2a9b6c56fdc64dcac9 bp-v1 -n porch-demo

cp demo-app-1-bp-v1.yaml bp-v1/demo-app-1.yaml

porchctl rpkg push management-08f3186a5068f90c7873cf2a9b6c56fdc64dcac9 bp-v1 -n porch-demo

porchctl rpkg propose management-08f3186a5068f90c7873cf2a9b6c56fdc64dcac9 -n porch-demo
porchctl rpkg approve management-08f3186a5068f90c7873cf2a9b6c56fdc64dcac9 -n porch-demo
```
## Create blueprint v2
```
porchctl rpkg copy management-08f3186a5068f90c7873cf2a9b6c56fdc64dcac9 --replay-strategy=true --workspace v2 -n porch-demo

porchctl rpkg pull management-8f9d314ce0ed3718910880812c193fd041771e1a bp-v2 -n porch-demo

cp demo-app-1-bp-v2-new-par-ok.yaml bp-v2/demo-app-1.yaml

porchctl rpkg push management-8f9d314ce0ed3718910880812c193fd041771e1a bp-v2 -n porch-demo

porchctl rpkg propose management-8f9d314ce0ed3718910880812c193fd041771e1a -n porch-demo
porchctl rpkg approve management-8f9d314ce0ed3718910880812c193fd041771e1a -n porch-demo
```
## Create deployment v1
```
porchctl rpkg clone management-08f3186a5068f90c7873cf2a9b6c56fdc64dcac9 depl --repository edge1 --workspace v1 -n porch-demo

porchctl rpkg pull edge1-721f9963e8b7461f5a53aa614e03baded3de791a depl-v1 -n porch-demo

cp demo-app-1-depl.yaml depl-v1/demo-app-1.yaml

porchctl rpkg push edge1-721f9963e8b7461f5a53aa614e03baded3de791a depl-v1 -n porch-demo

porchctl rpkg propose edge1-721f9963e8b7461f5a53aa614e03baded3de791a -n porch-demo
porchctl rpkg approve edge1-721f9963e8b7461f5a53aa614e03baded3de791a -n porch-demo
```
## Upgrade deployment v1
```
porchctl rpkg copy edge1-721f9963e8b7461f5a53aa614e03baded3de791a -n porch-demo --replay-strategy=true --workspace=v2
porchctl rpkg update edge1-1ebfa3e4ef95d04744d5069743cd350d910d743e --revision=v2
```
This works

# Bug scenario with blueprint parameter change (works)

## Create blueprint v1
```
porchctl rpkg init bp --workspace v1 --repository management -n porch-demo

porchctl rpkg pull management-08f3186a5068f90c7873cf2a9b6c56fdc64dcac9 bp-v1 -n porch-demo

cp demo-app-1-bp-v1.yaml bp-v1/demo-app-1.yaml

porchctl rpkg push management-08f3186a5068f90c7873cf2a9b6c56fdc64dcac9 bp-v1 -n porch-demo

porchctl rpkg propose management-08f3186a5068f90c7873cf2a9b6c56fdc64dcac9 -n porch-demo
porchctl rpkg approve management-08f3186a5068f90c7873cf2a9b6c56fdc64dcac9 -n porch-demo
```
## Create blueprint v2
```
porchctl rpkg copy management-08f3186a5068f90c7873cf2a9b6c56fdc64dcac9  --workspace v2 -n porch-demo --replay-strategy=true

porchctl rpkg pull management-8f9d314ce0ed3718910880812c193fd041771e1a bp-v2 -n porch-demo

cp demo-app-1-bp-v2-new-par-test.yaml bp-v2/demo-app-1.yaml

porchctl rpkg push management-8f9d314ce0ed3718910880812c193fd041771e1a bp-v2 -n porch-demo

porchctl rpkg propose management-8f9d314ce0ed3718910880812c193fd041771e1a -n porch-demo
porchctl rpkg approve management-8f9d314ce0ed3718910880812c193fd041771e1a -n porch-demo
```
## Create deployment v1
```
porchctl rpkg clone management-08f3186a5068f90c7873cf2a9b6c56fdc64dcac9 depl --repository edge1 --workspace v1 -n porch-demo

porchctl rpkg pull edge1-721f9963e8b7461f5a53aa614e03baded3de791a depl-v1 -n porch-demo

cp demo-app-1-depl.yaml depl-v1/demo-app-1.yaml

porchctl rpkg push edge1-721f9963e8b7461f5a53aa614e03baded3de791a depl-v1 -n porch-demo

porchctl rpkg propose edge1-721f9963e8b7461f5a53aa614e03baded3de791a -n porch-demo
porchctl rpkg approve edge1-721f9963e8b7461f5a53aa614e03baded3de791a -n porch-demo
```
## Upgrade deployment v1
```
porchctl rpkg copy edge1-721f9963e8b7461f5a53aa614e03baded3de791a -n porch-demo --replay-strategy=true --workspace=v2
porchctl rpkg update edge1-1ebfa3e4ef95d04744d5069743cd350d910d743e --revision=v2
```
This works
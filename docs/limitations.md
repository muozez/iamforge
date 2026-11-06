# Known Limitations

This document tracks known limitations and non-goals of `iamforge`.

## Static Analysis Scope
1. **Dynamic SDK Invocations:**
   Calls constructed dynamically (such as `getattr(client, action_name)(...)` or metaprogramming constructs) cannot be reliably traced statically. `iamforge` emits a warning and marks these calls as unresolved rather than guessing permissions.
2. **Inter-procedural Data Flow (v1):**
   In Phase 1, variable tracking is scoped locally within functions and files. Multi-package call graphs and cross-module object propagation will be expanded in later iterations.
3. **Runtime Permissions:**
   Actions and conditions that depend purely on runtime evaluation (e.g. dynamic STS assume-role session tags) cannot be inferred solely from source ASTs.

## AWS Account Interaction
- `iamforge` does not communicate with live AWS APIs, CloudTrail, or IAM Access Analyzer. All mapping and resolution is performed offline via embedded authorization schemas.

# iamforge

> Zero-backend Go CLI to scan code and Terraform configurations and generate least-privilege AWS IAM policies and minimal Terraform blocks.

[![CI](https://github.com/muozez/iamforge/actions/workflows/ci.yml/badge.svg)](https://github.com/muozez/iamforge/actions/workflows/ci.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Overview

Writing AWS IAM policies by hand often leads to security vulnerabilities (`Resource: "*"` wildcard over-permissioning) or frustrating runtime `AccessDenied` errors in AWS Lambda.

`iamforge` statically scans your application source code (starting with Python/`boto3`) and Terraform files locally:
- Extracts AWS SDK calls and parameters.
- Maps SDK calls to exact IAM actions without relying on guessing.
- Resolves ARN resources directly from code or linked Terraform configurations.
- Generates least-privilege `aws_iam_policy_document` Terraform HCL or IAM policy JSON.

### Core Principles
- **Zero Backend / Zero Telemetry:** Pure static analysis. Never makes network calls or accesses your cloud account.
- **Pure Go (CGO-free):** Single static binary compiled with `CGO_ENABLED=0` across Linux, macOS, and Windows.
- **Deterministic Output:** Identical input produces identical output bytes, perfect for CI pipelines.
- **Fail-Safe Warnings:** Never silently falls back to wildcard `*`. Always emits clear warnings with `TODO` markers when resources cannot be resolved.

## Quick Start

### Build from source
```bash
go build -o bin/iamforge ./cmd/iamforge
```

### Version check
```bash
./bin/iamforge version
# or JSON format
./bin/iamforge version --json
```

## Roadmap

- [x] **Phase 0:** Project foundation, Cobra CLI skeleton, CI & goreleaser config, pure-Go Python parser spike & ADR-0001
- [ ] **Phase 1:** Python (`boto3`) MVP extractor, core `model` and `policy` packages, JSON renderer
- [ ] **Phase 2:** Complete IAM action mapping dataset generator (`genmap`), `explain` command
- [ ] **Phase 3:** Terraform HCL parser & resource ARN resolver, `data.aws_iam_policy_document` renderer
- [ ] **Phase 4:** Node.js and Go SDK extractors
- [ ] **Phase 5:** Hardening, `--strict` mode, policy size limit checks, fuzz testing
- [ ] **Phase 6:** v0.1.0 release (Homebrew, Scoop, multi-platform binaries)

## License

This project is licensed under the [MIT License](LICENSE).

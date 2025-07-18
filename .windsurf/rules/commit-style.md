---
trigger: always_on
---

# Commit Style Guide

## Introduction

This document defines the commit message conventions for the Fury Credits MIA. Consistent commit messages help with:

- Automated changelog generation
- Understanding project history
- Facilitating code reviews
- Enabling better context in git history

## Commit Structure

Each commit message consists of:

```
<type>(<scope>): <subject>

<body>

<footer>
```

### Elements

- **Type**: Identifies the kind of change
- **Scope** (optional): Area of the codebase affected
- **Subject**: Brief description of the change
- **Body** (optional): Detailed explanation
- **Footer** (optional): References to issues, breaking changes

## Types

- **feat**: New feature
- **fix**: Bug fix
- **docs**: Documentation changes
- **style**: Formatting, missing semi-colons, etc; no code change
- **refactor**: Code refactoring without changing functionality
- **test**: Adding or modifying tests
- **chore**: Maintenance tasks, dependency updates, etc
- **perf**: Performance improvements

## Rules

1. Commit in English languaje.
2. Use imperative, present tense: "add" not "added" or "adds"
3. First line should be 50 characters or less
4. Reference issues and pull requests in the footer
5. Breaking changes must be noted in the footer prefixed with "BREAKING CHANGE:"
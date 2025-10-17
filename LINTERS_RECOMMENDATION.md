# Recommended golangci-lint Linters

This document explains the recommended linters that have been added to the `.golangci.yml` configuration file for the gopsutil project.

## Summary of Changes

The following linters have been added to enhance code quality, maintainability, and catch potential bugs:

### Previously Disabled but Now Enabled

1. **errcheck** - Critical for production code
   - Checks for unchecked errors in Go code
   - Configured to check type assertions and exclude common false positives like `io.Closer.Close()`
   - Essential for reliability in system monitoring code

2. **unused** - Code quality and maintenance
   - Detects unused constants, variables, functions, and types
   - Helps keep the codebase clean and maintainable
   - Important for a multi-platform project with many build tags

### New Linters for Modern Go (1.20+)

3. **gochecksumtype** - Type safety
   - Performs exhaustiveness checks on Go "sum types"
   - Helps catch missing cases in type switches
   - Particularly useful for the various platform-specific implementations

4. **intrange** - Modern Go idioms (Go 1.22+)
   - Suggests using integer range loops where appropriate
   - Makes code more idiomatic and readable
   - Auto-fixable

5. **copyloopvar** - Go 1.22+ optimization
   - Detects unnecessary loop variable copies
   - In Go 1.22+, loop variables are automatically scoped to iterations
   - Helps clean up code that worked around the old behavior

### Bug Prevention

6. **bodyclose** - Resource leak prevention
   - Ensures HTTP response bodies are closed properly
   - Critical for preventing resource leaks in long-running processes
   - Important for system monitoring tools that may make HTTP calls

7. **errname** - Error handling conventions
   - Ensures sentinel errors are prefixed with `Err`
   - Ensures error types are suffixed with `Error`
   - Promotes consistent error naming across the codebase

8. **nilerr** - Logic error detection
   - Finds code that returns nil even after checking error is not nil
   - Catches common copy-paste errors
   - Helps maintain correct error handling flow

9. **makezero** - Slice initialization bugs
   - Finds slice declarations with non-zero initial length that may cause bugs
   - Prevents subtle bugs where appending may not work as expected

### Code Quality and Style

10. **mirror** - Correct API usage
    - Detects wrong patterns of bytes/strings usage
    - Helps use the correct package for string/byte operations
    - Auto-fixable

11. **usestdlibvars** - Standard library constants
    - Detects places where stdlib variables/constants could be used
    - Promotes use of standard library constants over magic values
    - Auto-fixable

12. **godot** - Documentation quality
    - Ensures comments end with proper punctuation
    - Configured to check only declarations
    - Improves documentation consistency

13. **stylecheck** - Style consistency
    - Modern replacement for golint
    - Provides style checks beyond what revive covers
    - Configured to skip ST1003 (handled by revive)
    - Auto-fixable

14. **goimports** - Import management
    - Ensures imports are properly formatted and grouped
    - Configured to group local imports (github.com/shirou/gopsutil) separately
    - Auto-fixable
    - Complements the existing gci formatter

15. **unconvert** - Code cleanliness
    - Removes unnecessary type conversions
    - Makes code cleaner and easier to read
    - Auto-fixable

## Configuration Considerations

### errcheck Configuration
```yaml
errcheck:
  check-type-assertions: true
  check-blank: false
  exclude-functions:
    - (io.Closer).Close
```
- Checks type assertions for potential panics
- Excludes common patterns like defer Close() where error handling is often omitted

### godot Configuration
```yaml
godot:
  scope: declarations
  capital: false
```
- Only checks declaration comments (not all comments)
- Doesn't enforce capitalization to reduce noise

### goimports Configuration
```yaml
goimports:
  local-prefixes: github.com/shirou/gopsutil
```
- Groups gopsutil internal imports separately
- Maintains consistency with existing gci formatter

### stylecheck Configuration
```yaml
stylecheck:
  checks:
    - all
    - -ST1003  # Covered by revive
```
- Disables ST1003 since it's handled by revive's var-naming rule

## Benefits for gopsutil

These linters are particularly beneficial for gopsutil because:

1. **Multi-platform Support**: The project has many platform-specific implementations, and linters like `unused` and `gochecksumtype` help ensure all code paths are correct

2. **System Monitoring**: As a system monitoring library, resource management is critical. Linters like `bodyclose` and `errcheck` help prevent resource leaks

3. **Public API**: The library is widely used. Linters like `errname` and `godot` help maintain professional documentation and error handling conventions

4. **Go Version**: The project uses Go 1.24, so modern linters like `intrange` and `copyloopvar` can take advantage of newer language features

5. **Testing**: With `testifylint` already enabled, adding more error checking linters ensures test code is also high quality

## Auto-fixable Linters

Many of the new linters support automatic fixes:
- copyloopvar
- errorlint (already enabled)
- fatcontext (already enabled)
- gocritic (already enabled)
- goimports
- intrange
- mirror
- perfsprint (already enabled)
- unconvert
- usestdlibvars

This means many issues can be automatically corrected by running:
```bash
golangci-lint run --fix
```

## Further Recommendations

If the project wants to be even more strict in the future, consider:

1. **exhaustive** - For enum-like constants exhaustiveness checking
2. **goconst** - For detecting repeated strings that could be constants
3. **prealloc** - For performance optimization of slice pre-allocation
4. **revive** additional rules - More style rules can be enabled
5. **wrapcheck** - Ensure errors from external packages are wrapped
6. **nilnil** - Prevent returning (nil, nil)

These are not included yet as they may require more code changes or could be too noisy initially.

# golangci-lint Linters Enhancement - Summary

## Quick Overview

This PR adds **13 recommended linters** to the golangci-lint configuration, selected specifically for the gopsutil project based on expert analysis. It also moves `goimports` to the formatters section where it belongs in golangci-lint v2.

## What Changed

**Before:** 23 enabled linters (2 were disabled: errcheck, unused)  
**After:** 36 enabled linters (+13 new)

Note: This also re-enables 2 previously disabled linters (errcheck and unused) which are counted among the 13 new additions.

### New Linters Added

| Category | Linters | Auto-fix |
|----------|---------|----------|
| **Error Handling** | errcheck, errname, nilerr | Partial |
| **Code Quality** | unused, unconvert, makezero | Yes |
| **Modern Go (1.22+)** | intrange, copyloopvar, gochecksumtype | Yes |
| **Resource Management** | bodyclose | No |
| **Style & Docs** | godot | Yes |
| **Best Practices** | mirror, usestdlibvars | Yes |

### Configuration Fixes for golangci-lint v2

- ✅ **goimports** - Moved from linters to formatters section (formatters and linters are separate in v2)
- ✅ **stylecheck** - Removed (functionality merged into staticcheck in v2, ST checks already available)

## Key Highlights

### 1. Critical for Production (High Priority)
- ✅ **errcheck** - Catches unchecked errors (was disabled, now enabled)
- ✅ **bodyclose** - Prevents HTTP response body leaks
- ✅ **unused** - Removes dead code (was disabled, now enabled)

### 2. Leverages Modern Go
- ✅ **intrange** - Uses Go 1.22+ integer range loops
- ✅ **copyloopvar** - Optimizes loop variables (Go 1.22+)
- ✅ **gochecksumtype** - Type exhaustiveness checks

### 3. Auto-fixable (Easy Wins)
Many linters support `golangci-lint run --fix`:
- mirror, intrange, copyloopvar
- unconvert, usestdlibvars
- godot, makezero

And formatters with `golangci-lint run --fix`:
- goimports, gci, gofumpt

## Why These Linters?

### For gopsutil Specifically

1. **Multi-platform codebase** 
   - `unused` helps manage platform-specific code
   - `unconvert` in safe mode preserves architecture-specific conversions

2. **System monitoring library**
   - `bodyclose` prevents resource leaks
   - `errcheck` ensures reliable error handling

3. **Public API / Library**
   - `errname` enforces error naming conventions
   - `godot` improves documentation quality
   - ST checks in `staticcheck` maintain professional code style

4. **Modern Go support (1.24)**
   - Project can benefit from Go 1.22+ features
   - New linters leverage modern language improvements

## Configuration Decisions

All new linters are carefully configured:

- **errcheck**: Excludes common patterns like `defer file.Close()` where errors are typically ignored
- **unconvert**: Uses `safe: true` mode to avoid breaking platform-specific type conversions
- **godot**: Only checks declarations (not all comments) to reduce noise
- **goimports**: Configured as a formatter (not a linter) with grouping for gopsutil internal imports
- **staticcheck**: ST1003 check disabled (covered by existing revive rules)

## Next Steps

### Immediate
1. **Review this PR** - Check if the linter selections make sense
2. **Merge** - Configuration is ready to use
3. **CI will validate** - GitHub Actions will run with golangci-lint v2

### After Merge (Optional)
1. **Run auto-fix**: `golangci-lint run --fix` to automatically fix many issues
2. **Address remaining issues**: Review and fix linter warnings
3. **Consider stricter rules**: See `LINTERS_RECOMMENDATION.md` for future options

## Impact on CI

- ✅ Configuration is compatible with golangci-lint v2
- ✅ YAML syntax validated
- ⚠️ CI may report new issues (expected, this is the goal)
- 💡 Many issues can be auto-fixed

## Resources

- **Detailed Documentation**: See `LINTERS_RECOMMENDATION.md`
- **Configuration**: See `.golangci.yml`
- **golangci-lint docs**: https://golangci-lint.run/

## Questions?

If you have questions about:
- **Why a specific linter?** → See `LINTERS_RECOMMENDATION.md`
- **How to configure?** → Check `.golangci.yml` comments
- **CI failures?** → Many issues are auto-fixable with `--fix`

---

**TL;DR**: This PR adds 15 carefully selected linters to catch bugs, improve code quality, and leverage modern Go features. Many are auto-fixable. Configuration is production-ready and specifically tailored for gopsutil's multi-platform nature.

# Testing Strategy

## Test Packages

**Core:**

- `testing` (stdlib) - Base testing framework
- `github.com/stretchr/testify` - Assertions, mocks, test suites
- `modernc.org/sqlite` - In-memory database for integration tests
- `github.com/jackc/pgx/v5` - PostgreSQL testing

## Test Structure

```
internal/
├── features/
│   └── [feature]/
│       ├── handlers_test.go        # Unit tests
│       └── services/
│           └── *_service_test.go   # Unit tests
├── store/
│   ├── queries_test.go             # Integration tests
│   └── testutil/
│       ├── fixtures.go             # Test data
│       └── db.go                   # DB helpers
└── integration/
    └── *_flow_test.go              # E2E tests
```

## Test Categories

### Unit Tests (Fast, ~ms)

- Services with mocked store.Queries
- Handlers with mocked services
- Pure functions (formatTime, sqlNullString, encryption)
- **Pattern:** Mock dependencies, test business logic

### Integration Tests (Medium, ~100ms)

- sqlc queries against in-memory SQLite
- Services with real database
- NATS pub/sub patterns
- **Pattern:** Real DB, isolated features

### E2E Tests (Slow, ~seconds)

- Full HTTP request/response cycles
- SSE connections and real-time updates
- Multi-service flows
- **Pattern:** Real dependencies, full stack

## Task Commands

```bash
task test              # Run all tests
task test:unit         # Unit tests only (-short)
task test:integration  # Integration tests only
task test:coverage     # Generate coverage report
```

## Testing Patterns

### Services

```go
// Mock store.Queries interface
type MockQueries struct { mock.Mock }

func TestService_Method(t *testing.T) {
    mockQueries := new(MockQueries)
    service := NewService(mockQueries, "encryption-key-32-bytes-long!")

    mockQueries.On("QueryMethod", mock.Anything, mock.Anything).Return(expected, nil)

    result, err := service.Method(ctx, input)

    assert.NoError(t, err)
    assert.Equal(t, expected, result)
    mockQueries.AssertExpectations(t)
}
```

### Handlers

```go
func TestHandler_Endpoint(t *testing.T) {
    mockService := new(MockService)
    nc, _ := nats.Connect(nats.DefaultURL)
    handlers := NewHandlers(mockService, nc)

    mockService.On("Method", mock.Anything, input).Return(output, nil)

    req := httptest.NewRequest("POST", "/path", body)
    w := httptest.NewRecorder()

    handlers.Endpoint(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
}
```

### Store Queries

```go
type QueriesTestSuite struct {
    suite.Suite
    db      *sql.DB
    queries *store.Queries
}

func (s *QueriesTestSuite) SetupTest() {
    s.db, s.queries = testutil.NewTestDB(s.T())
}

func (s *QueriesTestSuite) TestQuery() {
    result, err := s.queries.QueryMethod(ctx, params)
    s.NoError(err)
    s.Equal(expected, result)
}

func TestQueriesTestSuite(t *testing.T) {
    suite.Run(t, new(QueriesTestSuite))
}
```

## Best Practices

1. **Use table-driven tests** for multiple scenarios
2. **Test error paths** explicitly
3. **Mock external dependencies** (NATS, PostgreSQL connections)
4. **Use testify suites** for setup/teardown
5. **Add t.Parallel()** for independent tests
6. **Test context cancellation** for long-running operations
7. **Use build tags** for slow tests: `// +build integration`
8. **Never commit test databases** - use in-memory or cleanup

## Coverage Goals

- **Services:** 80%+ coverage (business logic critical)
- **Handlers:** 70%+ coverage (orchestration layer)
- **Store queries:** 90%+ coverage (data integrity critical)
- **Overall:** 75%+ coverage

## CI/CD

Tests run automatically on push/PR via GitHub Actions. Coverage reports upload to Codecov.

## Key Files

- `internal/store/testutil/db.go` - Test database setup
- `internal/store/testutil/fixtures.go` - Test data factories
- Test files colocated with implementation: `*_test.go`

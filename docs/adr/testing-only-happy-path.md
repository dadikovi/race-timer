# Until further notice, for new code changes only the "happy path" must be tested

## Status

Accepted

## Context

Happy path tests are when we test the application behavior under normal circumstances:

- Application is properly integrated with the database
- Valid parameters are passed
- The state of the application / database is as expected

For a more robust application we should test the behavior eg. for invalid parameters.

## Decision

At this point we prefer development pace over robustness. So we will test only the "happy path".

## Consequences

Development will be faster, but the behavior of the application will be less predictable in error scenarios.
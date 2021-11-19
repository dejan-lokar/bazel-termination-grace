Run a shell and golang test that do not complete within 3s timeout and 3s grace termination period.
Expected result is for both tests to be killed after ~6s. That is not what happens for go test that
is killed after ~3s.

sh_test target is killed after 6s as expected (3s timeout + 3s grace).
```
 bazel test --test_output=streamed --cache_test_results=no --test_timeout=3 --local_termination_grace_seconds=3 //:sh_test
WARNING: Streamed test output requested. All tests will be run locally, without sharding, one at a time
INFO: Build options --cache_test_results, --test_sharding_strategy, and --test_timeout have changed, discarding analysis cache.
INFO: Analyzed target //:sh_test (0 packages loaded, 362 targets configured).
INFO: Found 1 test target...
Waiting for a signal... PID: 5177
-- Test timed out at 2021-11-19 09:50:50 UTC --
Terminated: 15
Trapped: INT or TERM

TIMEOUT: //:sh_test (Summary)
      /private/var/tmp/_bazel_cookie/2659f7d85f913a8b6fc1ee12315cf346/execroot/__main__/bazel-out/darwin-fastbuild/testlogs/sh_test/test.log
Target //:sh_test up-to-date:
  bazel-bin/sh_test
INFO: Elapsed time: 6.554s, Critical Path: 6.10s
INFO: 2 processes: 2 darwin-sandbox.
INFO: Build completed, 1 test FAILED, 2 total actions
//:sh_test                                                              TIMEOUT in 6.0s
  /private/var/tmp/_bazel_cookie/2659f7d85f913a8b6fc1ee12315cf346/execroot/__main__/bazel-out/darwin-fastbuild/testlogs/sh_test/test.log

INFO: Build completed, 1 test FAILED, 2 total actions
```

go_test target is killed after 3s. Termination grace appears to be ignored.
```
bazel test --test_output=streamed --cache_test_results=no --test_timeout=3 --local_termination_grace_seconds=3 //:go_test
WARNING: Streamed test output requested. All tests will be run locally, without sharding, one at a time
INFO: Analyzed target //:go_test (0 packages loaded, 7176 targets configured).
INFO: Found 1 test target...
Waiting for a signal... PID: 5279-- Test timed out at 2021-11-19 09:52:32 UTC --

TIMEOUT: //:go_test (Summary)
      /private/var/tmp/_bazel_cookie/2659f7d85f913a8b6fc1ee12315cf346/execroot/__main__/bazel-out/darwin-fastbuild/testlogs/go_test/test.log
Target //:go_test up-to-date:
  bazel-bin/go_test_/go_test
INFO: Elapsed time: 3.499s, Critical Path: 3.11s
INFO: 2 processes: 2 darwin-sandbox.
INFO: Build completed, 1 test FAILED, 2 total actions
//:go_test                                                              TIMEOUT in 3.0s
  /private/var/tmp/_bazel_cookie/2659f7d85f913a8b6fc1ee12315cf346/execroot/__main__/bazel-out/darwin-fastbuild/testlogs/go_test/test.log

INFO: Build completed, 1 test FAILED, 2 total actions
```

load("@io_bazel_rules_go//go:def.bzl",  "go_test")

package(default_visibility = ["//visibility:public"])

sh_test(
    name = "sh_test",
    srcs = ["trap_test.sh"],
)

go_test(
   name = "go_test",
   srcs = ["trap_test.go"],
)


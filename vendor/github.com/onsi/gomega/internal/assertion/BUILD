load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["assertion.go"],
    importmap = "github.com/kubernetes-sigs/poseidon/vendor/github.com/onsi/gomega/internal/assertion",
    importpath = "github.com/onsi/gomega/internal/assertion",
    visibility = ["//vendor/github.com/onsi/gomega:__subpackages__"],
    deps = ["//vendor/github.com/onsi/gomega/types:go_default_library"],
)

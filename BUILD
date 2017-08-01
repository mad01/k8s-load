load("@io_bazel_rules_go//go:def.bzl", "gazelle", "go_binary", "go_library", "go_prefix")

go_prefix("github.com/mad01/k8s-load")

gazelle(name = "gazelle")

go_library(
    name = "go_default_library",
    srcs = [
        "handler.go",
        "main.go",
    ],
    visibility = ["//visibility:private"],
)

go_binary(
    name = "k8s-load",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
)

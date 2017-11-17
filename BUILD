load("@io_bazel_rules_go//go:def.bzl", "gazelle", "go_library", "go_test")

gazelle(
    name = "gazelle",
    external = "vendored",
    prefix = "github.com/mattes/migrate",
)

go_library(
    name = "go_default_library",
    srcs = [
        "log.go",
        "migrate.go",
        "migration.go",
        "util.go",
    ],
    importpath = "github.com/mattes/migrate",
    visibility = ["//visibility:public"],
    deps = [
        "//database:go_default_library",
        "//source:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "migrate_test.go",
        "migration_test.go",
        "util_test.go",
    ],
    importpath = "github.com/mattes/migrate",
    library = ":go_default_library",
    deps = [
        "//database/stub:go_default_library",
        "//source:go_default_library",
        "//source/stub:go_default_library",
    ],
)

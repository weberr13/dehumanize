# dehumanize
Convert "humanized" sizes to approximate int64 sizes

Specifically written to convert elasticsearch's tendency to convert sizes to "human readable" 1.2mb style sizes.

Currently supports formats like:

1.1b
1kb
1.304mb
1.3gb
1.2tb

default for an "unconvertable" format is 0

Tests in goconvey

# timefmt

Utility to convert many date time formats into a readable format as well as Unix timestamps in different precisions.

## Usage
```
$ timefmt 2017-03-23 22:00:00 UTC
2017-03-23 18:00:00 -0400 EDT
2017-03-23 22:00:00 +0000 UTC
1490306400              seconds
1490306400000           milliseconds
1490306400000000        microseconds
1490306400000000000     nanoseconds

# nanoseconds
$ timefmt 1507955104960258397
2017-10-14 00:25:04.960258397 -0400 EDT
2017-10-14 04:25:04.960258397 +0000 UTC
1507955104              seconds
1507955104960           milliseconds
1507955104960258        microseconds
1507955104960258397     nanoseconds

# microseconds
$ timefmt 1507955104960258
2017-10-14 00:25:04.960258 -0400 EDT
2017-10-14 04:25:04.960258 +0000 UTC
1507955104              seconds
1507955104960           milliseconds
1507955104960258        microseconds
1507955104960258000     nanoseconds

# milliseconds
$ timefmt 1507955104960
2017-10-14 00:25:04.96 -0400 EDT
2017-10-14 04:25:04.96 +0000 UTC
1507955104              seconds
1507955104960           milliseconds
1507955104960000        microseconds
1507955104960000000     nanoseconds

# seconds
$ timefmt 1507955104
2017-10-14 00:25:04 -0400 EDT
2017-10-14 04:25:04 +0000 UTC
1507955104              seconds
1507955104000           milliseconds
1507955104000000        microseconds
1507955104000000000     nanoseconds

$ timefmt now
2017-11-01 16:08:27.908112455 -0400 EDT
2017-11-01 20:08:27.908112455 +0000 UTC
1509566907              seconds
1509566907908           milliseconds
1509566907908112        microseconds
1509566907908112455     nanoseconds

$ timefmt -unix-ns now
1509566907908112455

$ timefmt -unix-us now
1509566907908112

$ timefmt -unix-ms now
1509566907908

$ timefmt -unix-s now
1509566907
```
# OpenGauss Server Exporter

面向OpenGuass服务器的 Prometheus 监控采集器

Prometheus exporter for OpenGauss server metrics.

支持版本 Supported versions:
* OpenGauss >= 2.0.0.

注意：并非所有的方法支持OpenGauss 2.0.0以下的版本

NOTE: Not all collection methods are supported on OpenGauss < 2.0.0

# 1. 编译与运行Building and running

## 1.1 必备权限 Required Grants

```sql
CREATE USER 'exporter'@'localhost' IDENTIFIED BY 'XXXXXXXX' WITH MAX_USER_CONNECTIONS 3;
GRANT PROCESS, REPLICATION CLIENT, SELECT ON *.* TO 'exporter'@'localhost';
```

NOTE: It is recommended to set a max connection limit for the user to avoid overloading the server with monitoring scrapes under heavy load. This is not supported on all MySQL/MariaDB versions; for example, MariaDB 10.1 (provided with Ubuntu 18.04) [does _not_ support this feature](https://mariadb.com/kb/en/library/create-user/#resource-limit-options).

## 1.2 编译 Build

``` sh
git clone https://github.com/opengauss_plugin/opengauss_exporter.git
cd opengauss_exporter
go build
```

## 1.3 运行 Running

通过系统环境变量运行:
Running using an environment variable:

``` sh
export DATA_SOURCE_NAME="postgresql://exporter:XXXXXXXX@localhost:5432/postgres?sslmode=disable"

./opengauss_exporter
```

通过配置文件运行：
Running using ~/.opengauss.cnf:
``` sh
./opengauss_exporter
```


# 3.TODOLIST
- [x] (1) 应用级指标(重要) 预计时间：20210831
- [ ] (2) 系统级指标 预计时间：20210907
- [ ] (3) 数据库锁指标分析 预计时间：20210914
- [ ] (4) 数据库活跃链接指标分析 预计时间：20210921
- [ ] (5) 采集器设置参数分析 预计时间：20210928

# 4.效果预览（Grafana仪表盘示例）
## 4.1 添加DataSource
![1](https://raw.githubusercontent.com/opengauss_plugin/opengauss_exporter/master/showcase/1og.png)

## 4.2 查看DataSource
![2](https://raw.githubusercontent.com/opengauss_plugin/opengauss_exporter/master/showcase/2og.png)

## 4.3 创建DashBoard
![3](https://raw.githubusercontent.com/opengauss_plugin/opengauss_exporter/master/showcase/3og.png)

## 4.4 展示DashBoard
![4](https://raw.githubusercontent.com/opengauss_plugin/opengauss_exporter/master/showcase/4og.png)

# 5.Contributor

- [ryanemax](https://github.com/ryanemax)(刘雨飏), developer.
- [lyh2002](https://github.com/lyh2002)(刘宇航), developer.
- ***(万时超), developer.
- ***(陈若飞), developer.



# 附件：《采集器指标清单 Metrics Params》
## (1) 应用级指标(重要)
- 应用在运行时，语句性能的分析，用于作为优化SQL语句的参考数值

``` toml
# 慢查询分析
og_slow_select_count 2

# 查询频数分析
og_frequency_count 3
```

## (2) 系统级指标
- CPU、内存、垃圾回收情况等系统运行指标
``` toml
# HELP go_gc_duration_seconds A summary of the GC invocation durations.
# TYPE go_gc_duration_seconds summary
go_gc_duration_seconds{quantile="0"} 1.9007e-05
go_gc_duration_seconds{quantile="0.25"} 3.4058e-05
go_gc_duration_seconds{quantile="0.5"} 4.5702e-05
go_gc_duration_seconds{quantile="0.75"} 7.459e-05
go_gc_duration_seconds{quantile="1"} 0.001367991
go_gc_duration_seconds_sum 15.697670137
go_gc_duration_seconds_count 171695
# HELP go_goroutines Number of goroutines that currently exist.
# TYPE go_goroutines gauge
go_goroutines 10
# HELP go_info Information about the Go environment.
# TYPE go_info gauge
go_info{version="go1.11"} 1
# HELP go_memstats_alloc_bytes Number of bytes allocated and still in use.
# TYPE go_memstats_alloc_bytes gauge
go_memstats_alloc_bytes 3.53452e+06
# HELP go_memstats_alloc_bytes_total Total number of bytes allocated, even if freed.
# TYPE go_memstats_alloc_bytes_total counter
go_memstats_alloc_bytes_total 4.4309618728e+11
# HELP go_memstats_buck_hash_sys_bytes Number of bytes used by the profiling bucket hash table.
# TYPE go_memstats_buck_hash_sys_bytes gauge
go_memstats_buck_hash_sys_bytes 1.615564e+06
# HELP go_memstats_frees_total Total number of frees.
# TYPE go_memstats_frees_total counter
go_memstats_frees_total 5.321817792e+09
# HELP go_memstats_gc_cpu_fraction The fraction of this program's available CPU time used by the GC since the program started.
# TYPE go_memstats_gc_cpu_fraction gauge
go_memstats_gc_cpu_fraction 3.650956773482041e-06
# HELP go_memstats_gc_sys_bytes Number of bytes used for garbage collection system metadata.
# TYPE go_memstats_gc_sys_bytes gauge
go_memstats_gc_sys_bytes 2.37568e+06
# HELP go_memstats_heap_alloc_bytes Number of heap bytes allocated and still in use.
# TYPE go_memstats_heap_alloc_bytes gauge
go_memstats_heap_alloc_bytes 3.53452e+06
# HELP go_memstats_heap_idle_bytes Number of heap bytes waiting to be used.
# TYPE go_memstats_heap_idle_bytes gauge
go_memstats_heap_idle_bytes 6.1923328e+07
# HELP go_memstats_heap_inuse_bytes Number of heap bytes that are in use.
# TYPE go_memstats_heap_inuse_bytes gauge
go_memstats_heap_inuse_bytes 4.694016e+06
# HELP go_memstats_heap_objects Number of allocated objects.
# TYPE go_memstats_heap_objects gauge
go_memstats_heap_objects 25995
# HELP go_memstats_heap_released_bytes Number of heap bytes released to OS.
# TYPE go_memstats_heap_released_bytes gauge
go_memstats_heap_released_bytes 0
# HELP go_memstats_heap_sys_bytes Number of heap bytes obtained from system.
# TYPE go_memstats_heap_sys_bytes gauge
go_memstats_heap_sys_bytes 6.6617344e+07
# HELP go_memstats_last_gc_time_seconds Number of seconds since 1970 of last garbage collection.
# TYPE go_memstats_last_gc_time_seconds gauge
go_memstats_last_gc_time_seconds 1.6293815931829379e+09
# HELP go_memstats_lookups_total Total number of pointer lookups.
# TYPE go_memstats_lookups_total counter
go_memstats_lookups_total 0
# HELP go_memstats_mallocs_total Total number of mallocs.
# TYPE go_memstats_mallocs_total counter
go_memstats_mallocs_total 5.321843787e+09
# HELP go_memstats_mcache_inuse_bytes Number of bytes in use by mcache structures.
# TYPE go_memstats_mcache_inuse_bytes gauge
go_memstats_mcache_inuse_bytes 3456
# HELP go_memstats_mcache_sys_bytes Number of bytes used for mcache structures obtained from system.
# TYPE go_memstats_mcache_sys_bytes gauge
go_memstats_mcache_sys_bytes 16384
# HELP go_memstats_mspan_inuse_bytes Number of bytes in use by mspan structures.
# TYPE go_memstats_mspan_inuse_bytes gauge
go_memstats_mspan_inuse_bytes 48184
# HELP go_memstats_mspan_sys_bytes Number of bytes used for mspan structures obtained from system.
# TYPE go_memstats_mspan_sys_bytes gauge
go_memstats_mspan_sys_bytes 98304
# HELP go_memstats_next_gc_bytes Number of heap bytes when next garbage collection will take place.
# TYPE go_memstats_next_gc_bytes gauge
go_memstats_next_gc_bytes 4.194304e+06
# HELP go_memstats_other_sys_bytes Number of bytes used for other system allocations.
# TYPE go_memstats_other_sys_bytes gauge
go_memstats_other_sys_bytes 545324
# HELP go_memstats_stack_inuse_bytes Number of bytes in use by the stack allocator.
# TYPE go_memstats_stack_inuse_bytes gauge
go_memstats_stack_inuse_bytes 491520
# HELP go_memstats_stack_sys_bytes Number of bytes obtained from system for stack allocator.
# TYPE go_memstats_stack_sys_bytes gauge
go_memstats_stack_sys_bytes 491520
# HELP go_memstats_sys_bytes Number of bytes obtained from system.
# TYPE go_memstats_sys_bytes gauge
go_memstats_sys_bytes 7.176012e+07
# HELP go_threads Number of OS threads created.
# TYPE go_threads gauge
go_threads 7

```
## (3) 数据库锁指标分析

``` toml
# HELP og_exporter_last_scrape_duration_seconds Duration of the last scrape of metrics from OpenGauss.
# TYPE og_exporter_last_scrape_duration_seconds gauge
og_exporter_last_scrape_duration_seconds 0.018552361
# HELP og_exporter_last_scrape_error Whether the last scrape of metrics from OpenGauss resulted in an error (1 for error, 0 for success).
# TYPE og_exporter_last_scrape_error gauge
og_exporter_last_scrape_error 0
# HELP og_exporter_scrapes_total Total number of times OpenGauss was scraped for metrics.
# TYPE og_exporter_scrapes_total counter
og_exporter_scrapes_total 427015
# HELP og_locks_count Number of locks
# TYPE og_locks_count gauge
og_locks_count{datname="nova",mode="accessexclusivelock",server="localhost:5432"} 0
og_locks_count{datname="nova",mode="accesssharelock",server="localhost:5432"} 0
og_locks_count{datname="nova",mode="exclusivelock",server="localhost:5432"} 0
og_locks_count{datname="nova",mode="rowexclusivelock",server="localhost:5432"} 0
og_locks_count{datname="nova",mode="rowsharelock",server="localhost:5432"} 0
og_locks_count{datname="nova",mode="sharelock",server="localhost:5432"} 0
og_locks_count{datname="nova",mode="sharerowexclusivelock",server="localhost:5432"} 0
og_locks_count{datname="nova",mode="shareupdateexclusivelock",server="localhost:5432"} 0
```

## (4) 数据库活跃链接指标分析
``` toml

# HELP og_stat_activity_count number of connections in this state
# TYPE og_stat_activity_count gauge
og_stat_activity_count{datname="nova",server="localhost:5432",state="active"} 0
og_stat_activity_count{datname="nova",server="localhost:5432",state="disabled"} 0
og_stat_activity_count{datname="nova",server="localhost:5432",state="fastpath function call"} 0
og_stat_activity_count{datname="nova",server="localhost:5432",state="idle"} 7
og_stat_activity_count{datname="nova",server="localhost:5432",state="idle in transaction"} 0
og_stat_activity_count{datname="nova",server="localhost:5432",state="idle in transaction (aborted)"} 0

# HELP og_stat_activity_max_tx_duration max duration in seconds any active transaction has been running
# TYPE og_stat_activity_max_tx_duration gauge
og_stat_activity_max_tx_duration{datname="nova",server="localhost:5432",state="active"} 0
og_stat_activity_max_tx_duration{datname="nova",server="localhost:5432",state="disabled"} 0
og_stat_activity_max_tx_duration{datname="nova",server="localhost:5432",state="fastpath function call"} 0
og_stat_activity_max_tx_duration{datname="nova",server="localhost:5432",state="idle"} 0
og_stat_activity_max_tx_duration{datname="nova",server="localhost:5432",state="idle in transaction"} 0
og_stat_activity_max_tx_duration{datname="nova",server="localhost:5432",state="idle in transaction (aborted)"} 0
```

## (5)采集器设置参数分析
``` toml
# HELP og_static Version string as reported by postgres
# TYPE og_static untyped
og_static{server="localhost:5432",short_version="11.7.0",version="OpenGauss 11.7 (Debian 11.7-0+deb10u1) on x86_64-pc-linux-gnu, compiled by gcc (Debian 8.3.0-6) 8.3.0, 64-bit"} 1
# HELP og_up Whether the last scrape of metrics from OpenGauss was able to connect to the server (1 for yes, 0 for no).
# TYPE og_up gauge
og_up 1
# HELP opengauss_exporter_build_info A metric with a constant '1' value labeled by version, revision, branch, and goversion from which opengauss_exporter was built.
# TYPE opengauss_exporter_build_info gauge
opengauss_exporter_build_info{branch="",goversion="go1.11",revision="",version="0.0.1"} 1
# HELP process_cpu_seconds_total Total user and system CPU time spent in seconds.
# TYPE process_cpu_seconds_total counter
process_cpu_seconds_total 3386.7
# HELP process_max_fds Maximum number of open file descriptors.
# TYPE process_max_fds gauge
process_max_fds 65535
# HELP process_open_fds Number of open file descriptors.
# TYPE process_open_fds gauge
process_open_fds 10
# HELP process_resident_memory_bytes Resident memory size in bytes.
# TYPE process_resident_memory_bytes gauge
process_resident_memory_bytes 1.1956224e+07
# HELP process_start_time_seconds Start time of the process since unix epoch in seconds.
# TYPE process_start_time_seconds gauge
process_start_time_seconds 1.61409946935e+09
# HELP process_virtual_memory_bytes Virtual memory size in bytes.
# TYPE process_virtual_memory_bytes gauge
process_virtual_memory_bytes 1.15081216e+08
# HELP process_virtual_memory_max_bytes Maximum amount of virtual memory available in bytes.
# TYPE process_virtual_memory_max_bytes gauge
process_virtual_memory_max_bytes -1
# HELP promhttp_metric_handler_requests_in_flight Current number of scrapes being served.
# TYPE promhttp_metric_handler_requests_in_flight gauge
promhttp_metric_handler_requests_in_flight 1
# HELP promhttp_metric_handler_requests_total Total number of scrapes by HTTP status code.
# TYPE promhttp_metric_handler_requests_total counter
promhttp_metric_handler_requests_total{code="200"} 427012
promhttp_metric_handler_requests_total{code="500"} 1
promhttp_metric_handler_requests_total{code="503"} 0

```

# Copyright

http://www.futurestack.cn

Copyright © 2022 RyaneMax. All Rights Reserved.
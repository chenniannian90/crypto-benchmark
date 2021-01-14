# crypto-benchmark
用于测试各加密库签名/验签性能

### 依赖 
北大国密库需要必要的头文件
安装GmSSL 参考 http://gmssl.org/docs/install.html

```shell script
go test -bench . -run=none 

# 在本人 mac (2.7 GHz Intel Core i5, 8 GB 1867 MHz DDR3)上测试, 
# 推荐用户下载测试程序在专门机器测试
BenchmarkSign_CCS-4                28863             48782 ns/op 中国网安国密签名
BenchmarkSign_FABRIC-4             22818             59280 ns/op fabric默认配置签名
BenchmarkSign_PKU-4                20856             54522 ns/op 北大国密签名
BenchmarkSign_TJ-4                  2148            693351 ns/op 同济国密签名

BenchmarkVerify_CCS-4              10000            101032 ns/op 中国网安国密验签
BenchmarkVerify_FABRIC-4            8752            159480 ns/op fabric默认配置验签
BenchmarkVerify_PKU-4               8112            147563 ns/op 北大国密验签
BenchmarkVerify_TJ-4                 421           2921647 ns/op 同济国密验签

BenchmarkHash_CCS-4               598477              2086 ns/op 中国网安国密hash
BenchmarkHash_Fabric-4           1437912               790 ns/op fabric默认配置hash
BenchmarkHash_PKU-4               341826              3683 ns/op 北大国密hash
BenchmarkHash_TJ-4                596324              1938 ns/op 同济国密hash

```
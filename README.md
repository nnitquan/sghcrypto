# sghcrypto

## Introduce
* A command-line applications to encrypt or decrypt your important data
* Very easy to use, support interactive prompt, and as simple as possible
* Currently support aes crypto

## Install
Currently only support install cli from source.
"dev ensure" require a proxy if in china.
This is the linux example:
### Set your proxy
```
$ export http_proxy=http://localhost:58787
$ export https_proxy=http://localhost:58787
```
### Download source codes
```
$ cd $GOPATH/src && git clone https://github.com/sanguohot/sghcrypto
```
### Downlad packages
```
$ cd sghcrypto && dep ensure
```
### Install
```
$ go install
```
### Enjoin
```
$ sghcrypto
A command-line applications to encrypt or decrypt your important data
Very easy to use, support interactive prompt, and as simple as possible
```

## Examples
Help command:
```
$ sghcrypto h
NAME:
   sghcrypto - crypto for important data!

USAGE:
   sghcrypto [global options] command [command options] [arguments...]

VERSION:
   1.0.1

AUTHOR:
   Sanguohot <hw535431@163.com>

COMMANDS:
     encrypt, e, en  encrypt a message
     decrypt, d, de  decrypt a message
     help, h         Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```
AES encrypt use interactive prompt:
```
$ sghcrypto e hello
✔ Crypto Key(16 bytes): ****************
encrypt hello ===> WGM7Dteil6I1WDByRV3a2-yeXvD01YmXgIQUHYmOIhc
```
AES decrypt use interactive prompt:
```
$ sghcrypto d WGM7Dteil6I1WDByRV3a2-yeXvD01YmXgIQUHYmOIhc
✔ Crypto Key(16 bytes): ****************
decrypt WGM7Dteil6I1WDByRV3a2-yeXvD01YmXgIQUHYmOIhc ===> hello
```
AES encrypt use key flag:
```
$ sghcrypto e hello -k 1234567890123456
encrypt hello ===> wnOFhpMJhbEQ0UoLu_ugtvPfxZ54ozON8njBnHP8qLc
```
AES decrypt use key flag:
```
$ sghcrypto d wnOFhpMJhbEQ0UoLu_ugtvPfxZ54ozON8njBnHP8qLc -k 1234567890123456
decrypt wnOFhpMJhbEQ0UoLu_ugtvPfxZ54ozON8njBnHP8qLc ===> hello
```
Use CRYPTO_KEY env also support:
### Windows
```
$ set CRYPTO_KEY=1234567890123456
$ sghcrypto e hello
encrypt hello ===> Rs_jvGfUZFefVFiiUhkMCnkRLQuYuRkcDVYPtL9RZAI
```
### Linux
```
$ export CRYPTO_KEY=1234567890123456
$ sghcrypto d Rs_jvGfUZFefVFiiUhkMCnkRLQuYuRkcDVYPtL9RZAI
decrypt Rs_jvGfUZFefVFiiUhkMCnkRLQuYuRkcDVYPtL9RZAI ===> hello
```
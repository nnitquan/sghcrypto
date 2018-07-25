# sghcrypto

## Introduce
* A command-line applications to encrypt or decrypt your important data
* Very easy to use, support interactive prompt, and as simple as possible
* Currently support aes and ecies crypto

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
### Build it
```
$ go install
```
### Test it
```
$ sghcrypto e hello
Use the arrow keys to navigate: ↓ ↑ → ←
? Select Algorithm:
    aes
  ▸ ecies
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
ECIES encrypt use interactive prompt:
```
$ sghcrypto e hello
✔ ecies
✔ Public Key(64 bytes): ********************************************************************************************************************************
ecies encrypt hello ===> 0x04f30a03ae0ac8fcc6c8bba087ec27c42acef5d3ed870586a4a681e53fd5084a0a465c6557691acc3f732e144c858321ddee9e7e4c6c1312e1c0cd348611bf7925c80ef4dacf2ec2e166e5701fee77c4814287bd9a251895a252671fee8f8165ec54350f6bf1f00af34b70262a98c17ce3d1f9a428e9
```
ECIES decrypt use interactive prompt:
```
$ sghcrypto d 0x04f30a03ae0ac8fcc6c8bba087ec27c42acef5d3ed870586a4a681e53fd5084a0a465c6557691acc3f732e144c858321ddee9e7e4c6c1312e1c0cd348611bf7925c80ef4dacf2ec2e166e5701fee77c4814287bd9a251895a252671fee8f8165ec54350f6bf1f00af34b70262a98c17ce3d1f9a428e9
✔ ecies
✔ Private Key(32 bytes): ****************************************************************
ecies decrypt 0x04f30a03ae0ac8fcc6c8bba087ec27c42acef5d3ed870586a4a681e53fd5084a0a465c6557691acc3f732e144c858321ddee9e7e4c6c1312e1c0cd348611bf7925c80ef4dacf2ec2e166e5701fee77c4814287bd9a251895a252671fee8f8165ec54350f6bf1f00af34b70262a98c17ce3d1f9a428e9 ===> hello
```
AES encrypt use interactive prompt:
```
$ sghcrypto e hello
✔ aes
✔ AES Key(16 bytes): ****************
aes encrypt hello ===> AMCHTDF-q0jKdiXuNvoKcXYWvrx9HnDDLRriC6QVuFk
```
AES decrypt use interactive prompt:
```
$ sghcrypto d WGM7Dteil6I1WDByRV3a2-yeXvD01YmXgIQUHYmOIhc
✔ aes
✔ AES Key(16 bytes): ****************
aes decrypt WGM7Dteil6I1WDByRV3a2-yeXvD01YmXgIQUHYmOIhc ===> hello
```
AES encrypt use aeskey and alg flag:
```
$ sghcrypto e hello -aeskey 1234567890123456 --alg aes
aes encrypt hello ===> u1S3RUI5MTi2Yz_HBgkCIXhE25jAFUON1TeJuHAmUWc
```
AES decrypt use key flag:
```
$ sghcrypto d u1S3RUI5MTi2Yz_HBgkCIXhE25jAFUON1TeJuHAmUWc -aeskey 1234567890123456 --alg aes
aes decrypt u1S3RUI5MTi2Yz_HBgkCIXhE25jAFUON1TeJuHAmUWc ===> hello
```
### Use CRYPTO_ALG,AES_KEY,PRIVATE_KEY,PUBLIC_KEY env also support:
#### Windows
```
$ set CRYPTO_ALG=aes
$ set AES_KEY=1234567890123456
$ sghcrypto e hello
aes encrypt hello ===> Rs_jvGfUZFefVFiiUhkMCnkRLQuYuRkcDVYPtL9RZAI
$ sghcrypto d Rs_jvGfUZFefVFiiUhkMCnkRLQuYuRkcDVYPtL9RZAI
aes decrypt Rs_jvGfUZFefVFiiUhkMCnkRLQuYuRkcDVYPtL9RZAI ===> hello
```
#### Linux
```
$ export CRYPTO_ALG=ecies
$ export PRIVATE_KEY=fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19
$ export PUBLIC_KEY=9a7df67f79246283fdc93af76d4f8cdd62c4886e8cd870944e817dd0b97934fdd7719d0810951e03418205868a5c1b40b192451367f28e0088dd75e15de40c05
$ sghcrypto e hello
ecies encrypt hello ===> 0x0427a11ff1d6911e639d49caf155e8cea65ab59b36b9b5efc5290b78ead5e728752580fa0627baffd52dd928431d3847d2990892a72e4e60e26f43587c95f3239afde03044c55e9f2f249977a040e8fa1094f54d9ec078748a2cd1efc7a254ebaacd35378654a4ae8a7e1fc76be0476c28e4db4a6286
$ sghcrypto d 0x0427a11ff1d6911e639d49caf155e8cea65ab59b36b9b5efc5290b78ead5e728752580fa0627baffd52dd928431d3847d2990892a72e4e60e26f43587c95f3239afde03044c55e9f2f249977a040e8fa1094f54d9ec078748a2cd1efc7a254ebaacd35378654a4ae8a7e1fc76be0476c28e4db4a6286
ecies decrypt 0x0427a11ff1d6911e639d49caf155e8cea65ab59b36b9b5efc5290b78ead5e728752580fa0627baffd52dd928431d3847d2990892a72e4e60e26f43587c95f3239afde03044c55e9f2f249977a040e8fa1094f54d9ec078748a2cd1efc7a254ebaacd35378654a4ae8a7e1fc76be0476c28e4db4a6286 ===> hello
```
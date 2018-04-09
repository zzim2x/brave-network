# CLI

## glide 설치
```
curl https://glide.sh/get | sh
OR
brew install glide
```

## run

```
$ bin/brave
brave

Usage:
  brave [command]

Available Commands:
  account     account
  help        Help about any command
  keypair     keypair
  transaction transaction

Flags:
  -h, --help   help for brave
```

## 계정 생성과 함께 초기 잔고 전송

```
SEED=`bin/brave keypair generate | awk '{print $2}'`
ADDRESS=`bin/brave keypair parse --seed $SEED | awk '{print $2}'`

bin/brave transaction fund \
    --seed $SOURCE_ACCOUNT_SEED \
    --address $ADDRESS \
    --amount 2400
```


## 펀딩 받은 계정 또는 루트계정 잔고 확인

```
$ bin/brave account balance --alias root
My account address: GBADF6QQKKMM7A5N26VXJQGBYTXPWMURJVUXZ4VEHIROJJS2Q22ONK3E
type: native balance: 99999987137.9989300

$ bin/brave account balance --alias uncle
My account address: GAQMCPSHBREWQEUGGLSRTMGENQ6E6AIV4B7HB6KGBSUJCKYDQXIG3LAD
type: native balance: 410.0000000

$ bin/brave account balance --alias bob
My account address: GBWGODG5WLEH2BKRIQCLGMDOARCAVFI6KEKM7BKZL7EEXNBQMZOAP7ML
type: native balance: 2551.9999800

$ bin/brave account balance --alias jin
My account address: GCDLMIBCNNAR7S42XFSKAWWIUYM4RWKEJYQBQ6MO6WSP5QNQLSRO2ASL
type: native balance: 9899.9999800
```

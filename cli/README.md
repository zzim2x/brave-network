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

bin/brave transaction create \
    --seed $SOURCE_ACCOUNT_SEED \
    --address $ADDRESS \
    --amount 2400
```


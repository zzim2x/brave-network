# brave-network

**Independent stellar network; Brave Network**

local 에선 빌드 오래걸리고 공간을 차지하기에 stellar-core, horizon image hub.docker.com 에서 관리합니다.

* [docker image](https://github.com/zzim2x/stellar-docker)
* [automated build](https://hub.docker.com/r/zzim2x)

## 설치

```
# glide 설치 https://github.com/Masterminds/glide#install
$ brew install glide

$ go get github.com/zzim2x/brave-network/cli
$ cd `go env GOPATH`/src/github.com/zzim2x/brave-network/cli
$ make
```

## seed & node id 생성

현재 4대의 validator에서 사용할 node seed와 validator 설정값에 넣을 pubkey 생성해야합니다.

```
$ bin/brave keypair generate
Seed: SDFFK4WSRSUYGTWTXZTEEBWCKYN43LDU634ZYJTDD2BKG635PLXFJOPU

$ bin/brave keypair parse --seed=SDFFK4WSRSUYGTWTXZTEEBWCKYN43LDU634ZYJTDD2BKG635PLXFJOPU
Address: GD77GK4MNNODTAH6FTDM3I7W3UGILJTG4XFIMHLPSIHD76VBDWYLVOJ6
```

위 명령어를 4번해서 얻은 seed, id를 core.env 설정에 넣습니다. 현재는 미리 세팅되어져있습니다.

## 구동

```
$ cd `go env GOPATH`/src/github.com/zzim2x/brave-network/brave/docker-compose
$ docker-compose up
```

## 동작 확인

```
$ curl 'http://127.0.0.1:1168{1,2,3,4}/info'

[1/4]: http://127.0.0.1:11681/info --> <stdout>
--_curl_--http://127.0.0.1:11681/info
{
   "info" : {
      "build" : "v9.2.0",
      "ledger" : {
         "age" : 2,
         "baseFee" : 100,
         "baseReserve" : 100000000,
         "closeTime" : 1523597379,
         "hash" : "cfe5103543ff78d43a81ec5fee423c3b192da3698952c898bdb814ddd0aa45ba",
         "num" : 34116,
         "version" : 0
      },
      "network" : "Public Brave Stellar Network ; April 2018",
      "peers" : {
         "authenticated_count" : 3,
         "pending_count" : 0
      },
      "protocol_version" : 9,
      "quorum" : {
         "34115" : {
            "agree" : 4,
            "disagree" : 0,
            "fail_at" : 2,
            "hash" : "05d904",
            "missing" : 0,
            "phase" : "EXTERNALIZE"
         }
      },
      "startedOn" : "2018-04-13T05:23:30Z",
      "state" : "Synced!"
   }
}

[2/4]: http://127.0.0.1:11682/info --> <stdout>
--_curl_--http://127.0.0.1:11682/info
{
   ...
}

[3/4]: http://127.0.0.1:11683/info --> <stdout>
--_curl_--http://127.0.0.1:11683/info
{
   ...
}

[4/4]: http://127.0.0.1:11684/info --> <stdout>
--_curl_--http://127.0.0.1:11684/info
{
   ...
}
```

## 계정 생성

```
$ cd `go env GOPATH`/src/github.com/zzim2x/brave-network/cli

ROOT_SEED=SCTJ4RIYLZLA42675VYOE4QMERVDMQGTUGL44FLPXP3MN6JKD76CLG4M
ACCOUNT_SEED=`bin/brave keypair generate | awk '{print $2}'`
ACCOUNT_ADDRESS=`bin/brave keypair parse --seed=$NEW_ACCOUNT | awk '{print $2}'`

bin/brave transaction fund --seed $ROOT_SEED --address $ACCOUNT_ADDRESS --amount 10240
> transaction posted in ledger: 34221

bin/brave account balance --address $ACCOUNT_ADDRESS
> My account address: GAAL6DUVRMJTHGDW3PZSAMDLDPBR5GM6ENDTD2KGR2P6PPXHFXHHJQVG
> type: native balance: 10240.0000000
```

## 잔고 확인

```
cat config.yaml

network:
  # default network passphrase
  passphrase: "Public Brave Stellar Network ; April 2018"
  # horizon server url
  horizon: "http://127.0.0.1:8000"

account:
  root: GBADF6QQKKMM7A5N26VXJQGBYTXPWMURJVUXZ4VEHIROJJS2Q22ONK3E
  uncle: GAQMCPSHBREWQEUGGLSRTMGENQ6E6AIV4B7HB6KGBSUJCKYDQXIG3LAD
  # SA2HJJJFXGT3W2YFD5D3ARCAP5XV7RFBHS7FSSOCMSAG6UICLWNEEFBW
  bob: GBWGODG5WLEH2BKRIQCLGMDOARCAVFI6KEKM7BKZL7EEXNBQMZOAP7ML
  # SBCJYEB4OOAIAEX6ZQOV3VJ6YPVDF2WAAGXVHUD6L64GUZIUDE2SZYBE
  jin: GCDLMIBCNNAR7S42XFSKAWWIUYM4RWKEJYQBQ6MO6WSP5QNQLSRO2ASL
  # SBAZXBPX46HHHJFCZNAL5PS3FTFOKIVMMBDPBPOOFURCVSFRIBGQ2TNE
```

## 전송

```
bin/brave account balance --alias root
> My account address: GBADF6QQKKMM7A5N26VXJQGBYTXPWMURJVUXZ4VEHIROJJS2Q22ONK3E
> type: native balance: 99999956417.9988900

bin/brave transaction payment --seed $ROOT_SEED --address $ACCOUNT_ADDRESS --amount 10240
> transaction posted in ledger: 34335
> My account address: GAAL6DUVRMJTHGDW3PZSAMDLDPBR5GM6ENDTD2KGR2P6PPXHFXHHJQVG
> type: native balance: 20480.0000000

bin/brave account balance --alias root
> My account address: GBADF6QQKKMM7A5N26VXJQGBYTXPWMURJVUXZ4VEHIROJJS2Q22ONK3E
> type: native balance: 99999946177.9988800
```


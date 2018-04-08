# brave-network

**Independent stellar network; Brave Network**

* [docker image](https://github.com/zzim2x/stellar-docker)
* [automated build](https://hub.docker.com/r/zzim2x)

## init

* validator 노드로 쓸 만큼 `stellar-core --genseed` 생성해서 설정에 추가.
* `stellar-core --convertid` 로 PubKey 추출.

## docker-compose

docker-compose 기반으로 작성. 생각보다 깔끔하게 나오지 않는다. 시간되면 kubernetes 기반으로 만들어볼까 싶기도.

# brave-e2e

brave network 동작 하는가 테스트 코드

* transaction

# testnet

잘 동작하나 체크용

[상세 문서](https://github.com/zzim2x/brave-network/wiki/Brave-Network)

# shy-network

[fba implementation](https://github.com/zzim2x/shy)

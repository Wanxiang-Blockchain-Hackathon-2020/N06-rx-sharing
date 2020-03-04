#### # N06-ping.pub
<div align="center">

<h1>方块链 - 处方共享区块链</h2>

[![version](https://img.shields.io/github/tag/Wanxiang-Blockchain-Hackathon-2020/N06-rx-sharing.svg)](https://github.com/Wanxiang-Blockchain-Hackathon-2020/N06-rx-sharing/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/Wanxiang-Blockchain-Hackathon-2020/N06-rx-sharing)](https://goreportcard.com/report/github.com/Wanxiang-Blockchain-Hackathon-2020/N06-rx-sharing)
[![API Reference](https://godoc.org/github.com/Wanxiang-Blockchain-Hackathon-2020/N06-rx-sharing?status.svg)](https://godoc.org/github.com/Wanxiang-Blockchain-Hackathon-2020/N06-rx-sharing)
[![GitHub](https://img.shields.io/github/license/Wanxiang-Blockchain-Hackathon-2020/N06-rx-sharing.svg)](https://github.com/Wanxiang-Blockchain-Hackathon-2020/N06-rx-sharing/blob/master/LICENSE.md)

</div>

### 背景

2020新冠肺炎来袭。给中国的医疗体系带来了严俊的考验，特别是医疗信息化的数据透明度，数据孤岛的互联互通，数据的真假难辨是最为突出的问题，这些问题传统的IT技术很难完美解决。我们团队积极响应万向区块链公益黑客松的号召，尝试使用区块链技术来解决医疗行业的一些问题，为促进区块链在医疗行业的应用落地提供示范。

### 解决问题
针对国家医药分离的改革举措，提出处方流转共享区块链解决方案

### 应用场景
以城市或者地区为单位成立跨区域医疗信息协作共享平台。1.医生在医院为患者开处方，2.患者接受电子处方(在链上)，3.患者在就近药店买药(包括处方药)，4.药店核验处方真伪，药店依据处方售药。5.留存备案/政府监管

### 方案特点

- 去中心化，多方节点共同参与
- 数据不可篡改，不可抵赖，可追溯
- 数据加密上链，保护数据隐私。
- 数据公开透明。

### 使用帮助

#### 1.生成秘钥

```sh
iMac:~ liangping$ rxcli tx admin keygen --from ping
6e348eaeabff9556f06309d407e1d36726efc42736081976c34d765baba1663d
iMac:~ liangping$ rxcli tx admin keygen --from doctor
db2c1d03b7f364603458dec2cc4ecde060731e61002be6d2077f059aaa78a878
iMac:~ liangping$ rxcli tx admin keygen --from store
8c1b20e6f11e20853d48a297203cb57e459ed745632772076e7cd513cfebd759
```

#### 2 注册/登记/绑定

注意：医生，药店的注册，只能有特殊权限的用户才能操作，操作步骤类似。
```
iMac:~ liangping$ rxcli tx register patient --birthday "2000-10-10" --from ping --name "丽丽" --gender "女" --pubkey="6e348eaeabff9556f06309d407e1d36726efc42736081976c34d765baba1663d" -y --chain-id test
height: 3
txhash: 4AAB16E57906A73DE8DCB596802D7D8FA2E55BB07212E322ECCD179562659AE5
codespace: ""
code: 0
data: ""
rawlog: '[]'
logs: []
info: ""
gaswanted: 0
gasused: 0
tx: null
timestamp: ""

```
#### 3 查询绑定信息
```sh
iMac:~ liangping$ rxcli query admin patient --chain-id test  6e348eaeabff9556f06309d407e1d36726efc42736081976c34d765baba1663d
pubkey: 6e348eaeabff9556f06309d407e1d36726efc42736081976c34d765baba1663d
name: 丽丽
gender: 女
birthday: 2000-10-10T00:00:00Z
encrypted: ""
envelope: ""

```
#### 4 医生：开处方

处方数据`--rx "处方明文数据"`加密上链，数据采用点对点加密，秘钥随机生成，任何人均不知道，任何第三方无法解开

```sh
iMac:~ liangping$ rxcli tx doctor prescribe --from doctor --patient 6e348eaeabff9556f06309d407e1d36726efc42736081976c34d765baba1663d --rx '{name:"病人",rx:"other field"}' --memo "memo" -y --chain-id test --fees 1stake --gas-adjustment 2
height: 0
txhash: 6D366C042DF250FA5F022E6D3E05A51B167A98A914FF962ED82F6B336AA18194
codespace: ""
code: 0
data: ""
rawlog: '[]'
logs: []
info: ""
gaswanted: 0
gasused: 0
tx: null
timestamp: ""
```

#### 5 患者：查看处方
查看密文处方
```
iMac:~ liangping$ rxcli query patient rxs --keyname ping --chain-id test
- id: 6e31583233299
  doctor: db2c1d03b7f364603458dec2cc4ecde060731e61002be6d2077f059aaa78a878
  patient: 6e348eaeabff9556f06309d407e1d36726efc42736081976c34d765baba1663d
  status: "1"
  time: 2020-03-03T11:01:39.770366Z
  encrypted: T0QIFEHwHzGt3Edo48bcg82G4ReurkjMPZcCioDHDFA7qkXGuFpGw31RZszTWxUJ#v+nAvsKhVbKLopxI
  memo: memo
  salestore: ""
  saletime: 0001-01-01T00:00:00Z
```
查看明文处方
```
iMac:~ liangping$ rxcli query patient rx --decrypt --keyname ping --chain-id test --rx-id 6e31583233299
id: 6e31583233299
doctor: db2c1d03b7f364603458dec2cc4ecde060731e61002be6d2077f059aaa78a878
patient: 6e348eaeabff9556f06309d407e1d36726efc42736081976c34d765baba1663d
status: "1"
time: 2020-03-03T11:01:39.770366Z
encrypted: '{name:"病人",rx:"other field"}'
memo: memo
salestore: ""
saletime: 0001-01-01T00:00:00Z
```

#### 6 患者：授权药店或者其他医生查看处方
```
iMac:~ liangping$ rxcli tx patient authorize --from ping --chain-id test --rx-id 6e31583233299 --recipient 8c1b20e6f11e20853d48a297203cb57e459ed745632772076e7cd513cfebd759
{"chain_id":"test","account_number":"3","sequence":"3","fee":{"amount":[],"gas":"200000"},"msgs":[{"type":"patient/authorize","value":{"from":"cosmos14pkakt8apdm0e49tzp6gy3lwe8u04ajched5qm","patient":"6e348eaeabff9556f06309d407e1d36726efc42736081976c34d765baba1663d","drugstore":"8c1b20e6f11e20853d48a297203cb57e459ed745632772076e7cd513cfebd759","id":"yJ50OGoSVpfSJB7vAXPvPCPRkCVRPAvJBrxZ3zZ5I4M=#0SNsDhJ4cDOZhvVE","envelope":""}}],"memo":""}

confirm transaction before signing and broadcasting [y/N]: y
height: 0
txhash: 5D137F82DCD2C0E432790BD065C50E22635BC5427ABFE9E85F8E7FEE0DC909A3
codespace: ""
code: 0
data: ""
rawlog: '[]'
logs: []
info: ""
gaswanted: 0
gasused: 0
tx: null
timestamp: ""
```

#### 7 患者：查看授权记录
```
iMac:~ liangping$ rxcli query patient permits --keyname ping --chain-id test --rx-id 6e31583233299
- visitor: 6e348eaeabff9556f06309d407e1d36726efc42736081976c34d765baba1663d
  envelope: QpAzkFhF/Q/PhCRa0Hm5JKeGPDHWusAyDqvtYuHo7n4=#jddYYeiSh8KzZsBY
- visitor: 8c1b20e6f11e20853d48a297203cb57e459ed745632772076e7cd513cfebd759
  envelope: hQci40sTwmCrvlIOPO5lcdXwZwfNXH/gL6jCqxXHFP4=#iTCGsZDZADsUKE0e
```

#### 8 药店：查看处方
默认显示处方密文数据
```
iMac:~ liangping$ rxcli query drugstore view --keyname ping --chain-id test --rx-id 6e31583233299 --patient 6e348eaeabff9556f06309d407e1d36726efc42736081976c34d765baba1663d
id: 6e31583233299
doctor: db2c1d03b7f364603458dec2cc4ecde060731e61002be6d2077f059aaa78a878
patient: 6e348eaeabff9556f06309d407e1d36726efc42736081976c34d765baba1663d
status: "1"
time: 2020-03-03T11:01:39.770366Z
encrypted: T0QIFEHwHzGt3Edo48bcg82G4ReurkjMPZcCioDHDFA7qkXGuFpGw31RZszTWxUJ#v+nAvsKhVbKLopxI
memo: memo
salestore: ""
saletime: 0001-01-01T00:00:00Z
```
添加`--decrypt`显示处方明文数据
```
iMac:~ liangping$ rxcli query drugstore view --keyname store --chain-id test --patient 6e348eaeabff9556f06309d407e1d36726efc42736081976c34d765baba1663d --rx-id 6e31583250214 --decrypt
private hex: e1b0f79b20aa783b6aed89e1d82a5db06b7b560bbce3051090e2bb12f644f09820d00a25c9
 data key of decrypt: 27780605d14b56d01d6c80613b2a99e1 
id: 6e31583250214
doctor: db2c1d03b7f364603458dec2cc4ecde060731e61002be6d2077f059aaa78a878
patient: 6e348eaeabff9556f06309d407e1d36726efc42736081976c34d765baba1663d
status: "1"
time: 2020-03-03T15:43:34.222759Z
encrypted: '{name:"病人",rx:"other field"}'
memo: test drug decrypt
salestore: ""
saletime: 0001-01-01T00:00:00Z
```
#### 9 药店：售药
```
iMac:~ liangping$ rxcli tx drugstore sale --from store --chain-id test --rx-id 6e31583233299 --patient 6e348eaeabff9556f06309d407e1d36726efc42736081976c34d765baba1663d -y
height: 0
txhash: E64A8360A07711F8CC791F7B30318990293B31B75042E2F7A39B6CDA820DB64C
codespace: ""
code: 0
data: ""
rawlog: '[]'
logs: []
info: ""
gaswanted: 0
gasused: 0
tx: null
timestamp: ""
```
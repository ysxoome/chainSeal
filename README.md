# chainSeal

本项目根据高山科技git上公开的合约代码学习改写，在原有基础上加入了IPFS分布式存储服务以及一些细微的改变。原项目git地址:https://github.com/gaoshan-tech/sealcontract

# 项目组成
+ FISCO BCOS 2.9.0: FISCO BCOS是由国内企业主导研发、对外开源、安全可控的企业级金融联盟链底层平台，由金链盟开源工作组协作打造，并于2017年正式对外开源。
+ WeBASE: WeBASE(WeBank Blockchain Application Software Extension)是在区块链应用和FISCO-BCOS节点之间搭建的一套通用组件。
+ IPFS: 星际文件系统(InterPlanetary File System)IPFS是一个分布式的存储系统，在区块链领域经常作为外部组件被使用，用于存储区块链上的大型数据或文件。通过将数据存储在IPFS上，可以减轻区块链存储的负担，提高整体的可扩展性和性能。
+ goZero: go-zero 是一个集成了各种工程实践的 web 和 rpc 框架。包含极简的 API 定义和生成工具 goctl，可以根据定义的 api 文件一键生成 Go, iOS, Android, Kotlin, Dart, TypeScript, JavaScript 代码，并可直接运行。

系统是多机部署的区块链网络，同一群组两个节点部署在两台服务器上，节点使用rocksDB存储结构，WeBASE直接配置使用已有的多机区块链网络，整体概览如下图所示：
FISCO WeBASE:
<img width="1920" alt="截屏2023-08-25 09 35 35" src="https://github.com/ysxoome/chainSeal/assets/56349544/578f24e0-86f8-42d2-81d4-e8b1b5eb5a0a">
IPFS:
<img width="1920" alt="IPFS" src="https://github.com/ysxoome/chainSeal/assets/56349544/72f38f04-2f3d-4b99-93f1-c4208d4c4ac2">


# 文件结构
```
chainseal               服务目录
├── blockchain          链交互
│   ├── _ci             用户私钥
│   ├── contracts       合约代码
│   ├── sdk             链交互连接文件
│   └── tool            链上工具类，包含client初始化和合约session生成等
└── server              数据签名上链校验模块
    ├── common          整体大工具类
    │   ├── constant    常量文件
    │   ├── engine      合约调用方法集合
    │   ├── models      结构体文件
    │   ├── service     业务主逻辑文件
    │   ├── sh          私钥生成脚本文件
    │   ├── utils       工具类
    │   └── xerr        错误代码定义文件
    ├── config          用户私钥信息、用户did文件
    ├── etc             goZero静态配置文件
    └── internal        goZero单个服务内部文件
    │   ├── config      静态配置文件对应结构体声明
    │   ├── handler     路由管理
    │   ├── logic       业务目录 调用common包
    │   ├── svc         依赖注入目录
    │   └── types       结构体存放目录
    └── chain.go        程序启动入口
```
# personVC
主要参数项：
+ KeyStore(由个人私钥生成，包括id、地址、私钥)
+ 姓名
+ 身份证号

主要功能：
+ 生成凭证：admin使用个人的私钥文件生成一个可信凭证，取哈希即为链上identity。其中的issuer为admin，proof项为admin对凭证签名得到。
+ 验证凭证有效性：验证其是否为admin颁布签名及链上状态。

示例：
```json
{
  "claim": {
    "did": "did:gid:bee881dd425a1c5659061cc51e3eb7fda267ce11",
    "name": "测试",
    "type": "IDCard"
  },
  "context": "测试",
  "expirationDate": "2024-08-10 13:48:43.928892",
  "id": "ca12d0ac-4033-4e74-94c3-151b069bf7f2",
  "issuanceDate": "2023-08-11 13:48:43.910192",
  "issuer": "did:gid:bee881dd425a1c5659061cc51e3eb7fda267ce11",
  "proof": {
    "created": "2023-08-11 13:48:43.910192",
    "creator": "did:gid:bee881dd425a1c5659061cc51e3eb7fda267ce11",
    "signature": "c0V0Jv0fZu56QPdKwbcQzDa+anu2FIrcqqeIk+AULkB2/YR9EgSdm/uFZ0kSWpoo/u+gJvotQSX67j6mkK+5kQA=",
    "type": "Secp256k1"
    }
}
```
凭证哈希上链
<img width="1018" alt="VCHash上链" src="https://github.com/ysxoome/chainSeal/assets/56349544/5fa69a1c-85c5-48f7-97a2-394d1f0ed159">


# companyVC
主要参数项：
+ 法人did
+ 企业名称
+ 企业编号
+ 法人姓名
+ 企业类型

主要功能：
+ 生成凭证：调用链上方法，根据企业的法人did和企业编号生成其公司的公章地址，其公章地址也为companyDid。对凭证取哈希即为链上identity。其中的issuer为admin，proof项为admin对凭证进行签名得到。
+ 验证凭证有效性：验证其是否为admin颁布签名及链上状态。
+ 授权：使用公司的私钥文件对公章地址进行授权，私钥文件必须和公司的公章地址对应，targetDid传入授权用户信息，由authDid授权给其使用
+ 取消授权：所有步骤和授权相同，只是取消其使用印章权限

示例：
```json
{
  "claim": {
    "code": "91320000732251710X1691734407220",
    "did": "did:gid:7CFD7D437c048A8fe8E1b6df7D739C09C44A89A7",
    "legalDid": "did:gid:bee881dd425a1c5659061cc51e3eb7fda267ce11",
    "legalName": "张三1691734407220",
    "name": "测试",
    "type": "1"
  },
  "context": "测试",
  "expirationDate": "2033-08-08 14:13:27.259242",
  "id": "eceeed8e-4356-4e2b-8ca7-14d9d4e42f2b",
  "issuanceDate": "2023-08-11 14:13:27.259242",
  "issuer": "did:gid:bee881dd425a1c5659061cc51e3eb7fda267ce11",
  "proof": {
    "created": "2023-08-11 14:13:27.259242",
    "creator": "did:gid:bee881dd425a1c5659061cc51e3eb7fda267ce11",
    "signature": "n6ezr9LKM5byTDQvn3rJRVdVh8e0LW/U+nBNA67WGZsmN1f/yNrvbG+KqCqjhPZeJ/eSMHKfn0EKd/RioWk0rQA=",
    "type": "Secp256k1"
  }
}
```
凭证哈希上链 同personVC

构造企业签章地址
<img width="1021" alt="构造企业签章地址" src="https://github.com/ysxoome/chainSeal/assets/56349544/2a8971a0-cb6b-4715-98d4-8bb92e492d82">
企业签章授权
<img width="1022" alt="approval" src="https://github.com/ysxoome/chainSeal/assets/56349544/6eabdf21-d80f-40d9-b1d3-8bdf3d1978aa">
企业签章取消授权
<img width="1021" alt="revoke" src="https://github.com/ysxoome/chainSeal/assets/56349544/ab20e88f-bcb4-4a21-b7b6-14b58e4e032f">
# contracts
主要参数项：
+ 企业可信凭证
+ IPFS上传文件返回的cid码
+ IPFS上传文件的哈希值
+ 嵌套签名
+ 公章信息（json字符串格式）

# 主要功能
+ 生成合约凭证：通过企业凭证、印章信息、合同信息生成合约凭证，并由admin进行签名。
+ 嵌套生成合约凭证：将上一次的合约凭证作为参数结合新的合约信息生成新的合约凭证。
+ 上链前校验：验证一些重要信息如合约格式、链上issuer是否仍有效、验证其中的签名信息和该印章地址是否有授权。
+ 合约凭证上链：通过合约凭证的cid码、合约文件哈希、签名哈希、印章地址、印章类型和合约的签名人（admin）信息完成上链。
+ 合约凭证查询：通过上述信息查询该次上链的在该文件哈希中的序号，通过文件哈希和序号查询详细信息
+ 溯源：根据合约凭证的嵌套签名追溯到每次的盖章情况。

示例：
```json
{
    "context": "测试科技研究院",
    "id": "0db99332-a5c7-41b6-94ce-a220b3b3e354",
    "contractHash": "0xce779b350231c5eb8b1bba07167f8760f1312c5ef2c2226dd4fd488a2eb7c7f9",
    "nestSignature": {
        "context": "测试科技研究院",
        "id": "98032b59-bc3f-46b4-9292-85419b4d30e4",
        "contractHash": "0xce779b350231c5eb8b1bba07167f8760f1312c5ef2c2226dd4fd488a2eb7c7f9",
        "nestSignature": {
            "context": "测试科技研究院",
            "id": "0f97e527-c2ee-4897-8f01-f312a0b4ac0d",
            "contractHash": "0xce779b350231c5eb8b1bba07167f8760f1312c5ef2c2226dd4fd488a2eb7c7f9",
            "nestSignature": null,
            "sealType": 1,
            "signerVC": {
                "context": "测试科技研究院",
                "id": "853e94b6-fc23-4abd-8409-bdf0b013eda5",
                "issuer": "did:gid:83309d045a19c44dc3722d15a6abd472f95866ac",
                "issuanceDate": "2023-08-28 08:53:14",
                "expirationDate": "2033-08-25 08:53:14",
                "claim": {
                    "creditCode": "9132000073225110X",
                    "did": "did:gid:07A17B46fCe6588A6ED1abE9cbf837b598B1eCB8",
                    "legalDid": "did:gid:bee881dd425a1c5659061cc51e3eb7fda267ce11",
                    "legalName": "张三",
                    "name": "测试信息技术研究院有限公司",
                    "type": "1"
                },
                "proof": {
                    "created": "2023-08-28 08:53:14",
                    "creator": "did:gid:83309d045a19c44dc3722d15a6abd472f95866ac",
                    "signature": "u2M/IHbgf7Sae3w3K89Cp7FWvY1S9IQtwfGTF57XWH51D3n3dn6WHCkouqiiCbpwwBhjuuWr1PewXM2s18in5AE=",
                    "type": "Secp256k1"
                }
            },
            "signDate": "2023-08-28 08:53:14",
            "sealsClaim": {
                "name": "测试科技研究院",
                "phone": "12345678912",
                "position": "0",
                "sign": "https://www.zbiti.com",
                "signIndex": "0",
                "signType": "0",
                "x": "204.575",
                "y": "198.73"
            },
            "proof": {
                "created": "2023-08-28 08:53:14",
                "creator": "did:gid:83309d045a19c44Dc3722D15A6AbD472f95866aC",
                "signature": "opKiywsmJZH+AB9Yai9TWrzPXktC875xTRjOExiRnroBMDAGu+x2hutV/lDBkXdAU0Rz01tIVsl2fHvogcKmagE=",
                "type": "Secp256k1"
            }
        },
        "sealType": 1,
        "signerVC": {
            "context": "测试科技研究院",
            "id": "e77097b1-95a0-45bb-9fde-8e3898c571c5",
            "issuer": "did:gid:83309d045a19c44dc3722d15a6abd472f95866ac",
            "issuanceDate": "2023-08-28 08:53:14",
            "expirationDate": "2033-08-25 08:53:14",
            "claim": {
                "creditCode": "9132000073225110X",
                "did": "did:gid:07A17B46fCe6588A6ED1abE9cbf837b598B1eCB8",
                "legalDid": "did:gid:bee881dd425a1c5659061cc51e3eb7fda267ce11",
                "legalName": "张三",
                "name": "测试信息技术研究院有限公司",
                "type": "1"
            },
            "proof": {
                "created": "2023-08-28 08:53:14",
                "creator": "did:gid:83309d045a19c44dc3722d15a6abd472f95866ac",
                "signature": "u2M/IHbgf7Sae3w3K89Cp7FWvY1S9IQtwfGTF57XWH51D3n3dn6WHCkouqiiCbpwwBhjuuWr1PewXM2s18in5AE=",
                "type": "Secp256k1"
            }
        },
        "signDate": "2023-08-28 08:53:15",
        "sealsClaim": {
            "name": "测试科技研究院",
            "phone": "12345678912",
            "position": "0",
            "sign": "https://www.zbiti.com",
            "signIndex": "0",
            "signType": "0",
            "x": "204.575",
            "y": "198.73"
        },
        "proof": {
            "created": "2023-08-28 08:53:15",
            "creator": "did:gid:83309d045a19c44Dc3722D15A6AbD472f95866aC",
            "signature": "eahNynIswW0J5dP4Qtco/4pAUU5GZW30yuS6p/Qg9dUeOWqK56RosqCcfnMtWIIDgO7D3wry/gdJH5bCTQ0BtAE=",
            "type": "Secp256k1"
        }
    },
    "sealType": 1,
    "signerVC": {
        "context": "测试科技研究院",
        "id": "05bd8e42-2a09-4f38-ae3d-ec3cb4b18efd",
        "issuer": "did:gid:83309d045a19c44dc3722d15a6abd472f95866ac",
        "issuanceDate": "2023-08-28 08:53:15",
        "expirationDate": "2033-08-25 08:53:15",
        "claim": {
            "creditCode": "9132000073225110X",
            "did": "did:gid:07A17B46fCe6588A6ED1abE9cbf837b598B1eCB8",
            "legalDid": "did:gid:bee881dd425a1c5659061cc51e3eb7fda267ce11",
            "legalName": "张三",
            "name": "测试信息技术研究院有限公司",
            "type": "1"
        },
        "proof": {
            "created": "2023-08-28 08:53:15",
            "creator": "did:gid:83309d045a19c44dc3722d15a6abd472f95866ac",
            "signature": "u2M/IHbgf7Sae3w3K89Cp7FWvY1S9IQtwfGTF57XWH51D3n3dn6WHCkouqiiCbpwwBhjuuWr1PewXM2s18in5AE=",
            "type": "Secp256k1"
        }
    },
    "signDate": "2023-08-28 08:53:15",
    "sealsClaim": {
        "name": "测试科技研究院",
        "phone": "12345678912",
        "position": "0",
        "sign": "https://www.zbiti.com",
        "signIndex": "0",
        "signType": "0",
        "x": "204.575",
        "y": "198.73"
    },
    "proof": {
        "created": "2023-08-28 08:53:15",
        "creator": "did:gid:83309d045a19c44Dc3722D15A6AbD472f95866aC",
        "signature": "lOe8YoEKxuRLTcyA5tq60JRcl8pIqu8ZtJB9AYQov8RTFBW7j+1rIvYoYlZRynxCN52CSFW5Hkn69ujJ/a8kRQA=",
        "type": "Secp256k1"
    }
}
```
合同信息上链
<img width="1022" alt="addSign" src="https://github.com/ysxoome/chainSeal/assets/56349544/233f44d9-6d60-4cf8-b1cf-82515325fd0d">

# 接口截图
个人凭证生成
<img width="1920" alt="genPVC" src="https://github.com/ysxoome/chainSeal/assets/56349544/61a3bfda-ae88-4f68-939e-4fee03822b08">
企业凭证生成
<img width="1920" alt="genCVC" src="https://github.com/ysxoome/chainSeal/assets/56349544/3544e06e-7828-4a26-ba9e-3b97c4f2f91f">
企业签章授权
<img width="1920" alt="approval" src="https://github.com/ysxoome/chainSeal/assets/56349544/72117d52-8591-4f88-9751-74a00f76594d">
企业签章取消授权
<img width="1920" alt="revork" src="https://github.com/ysxoome/chainSeal/assets/56349544/0b5506c9-931d-4b59-8612-52125733053a">
合同文件联合IPFS上链
<img width="1920" alt="addSign" src="https://github.com/ysxoome/chainSeal/assets/56349544/0de5fdc6-2f73-4a39-b7e2-ca6b176e547b">
链上合同查询
<img width="1920" alt="querySign" src="https://github.com/ysxoome/chainSeal/assets/56349544/f97d23ad-8b19-4e9f-8485-2bd6f066d636">
CID文件下载
<img width="1920" alt="downloadCID" src="https://github.com/ysxoome/chainSeal/assets/56349544/d7a989b3-d724-4d57-b7fe-7f44ceaf5041">





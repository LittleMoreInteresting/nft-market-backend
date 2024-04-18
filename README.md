# NFT-Market Backend by go-Kratos

## 部署
 
- 复制配置文件 `config.yaml.example` 为 `configs/config.yaml` 修改 **
- `make build` 编译项目 生成可执行文件 `bin/nft-market-backend`
- 执行项目 `./bin/nft-market-backend -conf <configs 配置文件目录>`


TODO 
- 后端批量查询nft元数据并返回
- https://docs.moralis.io/web3-data-api/evm/reference/get-multiple-nfts?tokens=[{%22token_address%22:%220xa4991609c508b6d4fb7156426db0bd49fe298bd8%22,%22token_id%22:%2212%22},{%22token_address%22:%220x3c64dc415ebb4690d1df2b6216148c8de6dd29f7%22,%22token_id%22:%221%22},{%22token_address%22:%220x3c64dc415ebb4690d1df2b6216148c8de6dd29f7%22,%22token_id%22:%22200%22}]&normalizeMetadata=false&media_items=true&chain=eth
- 
# GetTorrentListOkResponse

**Properties**

| Name    | Type                                    | Required | Description |
| :------ | :-------------------------------------- | :------- | :---------- |
| Data    | []torrents.GetTorrentListOkResponseData | ❌       |             |
| Detail  | string                                  | ❌       |             |
| Success | bool                                    | ❌       |             |

# GetTorrentListOkResponseData

**Properties**

| Name             | Type                  | Required | Description |
| :--------------- | :-------------------- | :------- | :---------- |
| Active           | bool                  | ❌       |             |
| AuthId           | string                | ❌       |             |
| Availability     | float64               | ❌       |             |
| CreatedAt        | string                | ❌       |             |
| DownloadFinished | bool                  | ❌       |             |
| DownloadPresent  | bool                  | ❌       |             |
| DownloadSpeed    | float64               | ❌       |             |
| DownloadState    | string                | ❌       |             |
| Eta              | float64               | ❌       |             |
| ExpiresAt        | string                | ❌       |             |
| Files            | []torrents.DataFiles1 | ❌       |             |
| Hash             | string                | ❌       |             |
| Id               | float64               | ❌       |             |
| InactiveCheck    | float64               | ❌       |             |
| Magnet           | string                | ❌       |             |
| Name             | string                | ❌       |             |
| Peers            | float64               | ❌       |             |
| Progress         | float64               | ❌       |             |
| Ratio            | float64               | ❌       |             |
| Seeds            | float64               | ❌       |             |
| Server           | float64               | ❌       |             |
| Size             | float64               | ❌       |             |
| TorrentFile      | bool                  | ❌       |             |
| UpdatedAt        | string                | ❌       |             |
| UploadSpeed      | float64               | ❌       |             |

# DataFiles1

**Properties**

| Name      | Type    | Required | Description |
| :-------- | :------ | :------- | :---------- |
| Id        | float64 | ❌       |             |
| Md5       | string  | ❌       |             |
| Mimetype  | string  | ❌       |             |
| Name      | string  | ❌       |             |
| S3Path    | string  | ❌       |             |
| ShortName | string  | ❌       |             |
| Size      | float64 | ❌       |             |
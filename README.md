# YouTube Livestreaming Notificator

YouTube上でのライブストリーミングの枠を検知してSlackへ通知します。
以下の機能があります。
- 配信枠が設定された際に通知
- 配信予定日時または配信タイトルが変更された際に通知
- 配信開始時刻にリマインド通知

# How to use

[sample config](./config.yml.sample)を元に `config.yml` を用意してください。

```
Youtube:
  ApiKey: <GOOGLE_API_KEY>
Slack: <SLACK_WEBHOOK_URL>
TargetAccounts: 
  - <TARGET_ACCOUNT_EXTERNALID>
```

`TARGET_ACCOUNT_EXTERNALID` はYoutubeのチャンネルIDです。  
カスタムURLが設定されているチャンネルではブラウザ画面上から取得することができません。  
ブラウザの開発者ツールを利用して `ExternalId` などで検索するとjsonフォーマットで記載がある文字列を使用してください。

また、現在は5分間隔で処理を行っていますが、YouTube Data API v3のデフォルトの1日あたりのquotaは10,000のため、検索対象のチャンネルが多い場合はリミットを超過する可能性があります。  
リミットを超過する場合は、quotaの割当量増加を申請するか、更新レートを変更してください。

# Milestones
- [ ] Discordへの通知サポート
- [x] 枠設定後に配信日時が変更された場合に配信開始通知のリスケジューリング
- [ ] docker-composeのサポート
- [ ] 更新レートの設定サポート
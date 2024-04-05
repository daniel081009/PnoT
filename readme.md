# PnoT

## Description

노션은 너무 무겁고, 옵시디언은 동기화가 불편하다. 안전하게 노트를 저장하고 싶은데, 빠르게 동기화하고 싶어서 만드는 중이다.

전용서버도 운용할 각오를 가진 사람을 위한 노트 앱이다.

PnoT은 노트앱을 쉽게 저장,공유,동기화, 커스터마이징이 가능한 노트 앱을 목표로 만들고 있다.

## Features

- PWA (Progressive Web App)
- P2P (Peer to Peer)
- 1초마다 자동 동기화 (Auto Sync)
- 노트 히스토리 (Note History)
- 마크다운 에디터 (Markdown Editor)
- 손쉬운 노트 공유 (Note Sharing)
- 노트 암호화 (Note Encryption)

## Tech Stack

### Frontend

- PWA (Progressive Web App)
- Lit
- socket.io

### Backend

- Golang
- socket.io

## issue

노트를 연동할때 노트를 수정하면 암호화하고 서버로 보내고 서버에선 저장하고 그대로 다른 클라이언트에게 보내준다. 속도가 괜찮을까?
특히 저사양의 기기에서는 파일이 크면 더더욱 느릴것이다.
그래서 합리화를 하면 서버에서 암호화를 해서 저장하는건데 이러면 서버가 암호화된 파일을 볼수있는데 목적에 맞지 않다.

## TODO BackEnd

- [x] Auth
  - [x] user rsa keygen
  - [x] user file auth
- [x] File Upload,Download
  - [x] file upload
  - [x] file download
- [x] File History
- [ ] File Sync
- [ ] File Share

## TODO FrontEnd

- [] Auth
  - [] login
  - [] signup
- [] File CRUD
  - [] File Create
  - [] File Read
  - [] File Update
  - [] File Delete
  - [] File History
- [] File Sync
- [] File Share

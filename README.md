# SSH 접속을 통해서 VM 시스템 데이터를 수집하는 프로그램

이 프로그램은 SSH를 통해 가상 머신(VM)에 접속하여 시스템 데이터를 수집하는 기능을 제공합니다. 수집된 데이터는 성능 모니터링, 문제 해결, 리소스 관리 등에 활용될 수 있습니다.

## 목차
1. [개요](#개요)
2. [특징](#특징)
3. [구성 파일](#구성-파일)

## 개요
이 프로그램은 여러 VM에 대해 SSH를 통해 연결을 설정하고, CPU 사용률, 메모리 사용량, 디스크 사용량 등 다양한 시스템 데이터를 수집합니다. 수집된 데이터는 로컬 파일 시스템에 저장되거나 원격 데이터베이스로 전송될 수 있습니다.

## 특징
- **멀티 스레드 지원**: 동시에 여러 VM에 접속하여 데이터를 수집합니다.
- **확장성**: 다양한 시스템 데이터를 수집할 수 있도록 쉽게 확장 가능합니다.
- **안정성**: SSH 연결이 실패할 경우 재시도 로직이 포함되어 있습니다.
- **유연성**: 사용자 정의 명령어를 통해 추가 데이터를 수집할 수 있습니다.

## 구성 파일
```
.
├── README.md
├── collector
│   └── collector.go
├── ft.txt
├── go.mod
├── go.sum
├── internal
│   ├── dblinker
│   │   ├── dbconnection.go
│   │   ├── dbmd
│   │   │   ├── BtAgent.go
│   │   │   ├── BtCollectorAttr.go
│   │   │   ├── BtModuleAttr.go
│   │   │   ├── BtResource.go
│   │   │   ├── BtResourceAttr.go
│   │   │   ├── BtResourceStat.go
│   │   │   ├── VwResourceDetail.go
│   │   │   ├── bt_collector.gen.go
│   │   │   ├── bt_collector_provider.gen.go
│   │   │   ├── bt_collector_type.gen.go
│   │   │   ├── bt_module.gen.go
│   │   │   ├── bt_provider.gen.go
│   │   │   ├── bt_provider_attr.gen.go
│   │   │   └── bt_provider_type.gen.go
│   │   └── load_db.go
│   ├── lslinker
│   │   ├── lsapi
│   │   │   └── lsapi.go
│   │   └── lsmd
│   │       └── lsmd.go
│   ├── ssh
│   │   ├── client.go
│   │   ├── cmdline
│   │   │   └── cmdline.go
│   │   └── disk
│   │       └── disk.go
│   ├── sysdef
│   │   └── init.go
│   └── sysenv
│       └── sysenv.go
├── main.go
├── main_LoadEnv.go
├── main_LoadYml.go
├── print_sdk_version.sh
├── push_sdk_newversion.sh
├── setting.yml
└── utils
    └── router
        └── router.go
```

# Kisara
Kisara 是一个用于CTF/AWD等竞赛的Docker集群管理工具，现在尚处于研发阶段

## 为什么叫Kisara
来自于我很喜欢的动画《契约之吻》的女主Kisara，她是一个恶魔！而我对项目的命名基本都是美少女，所以就这样了（雾

## Kisara的由来
因为Irina（伊蕾娜）项目中需要对Docker管理进行解耦，同时需要接入集群，使用Irina原本的单机Docker是无法实现的，因此有了Kisara，
同时Kisara使用了Takina（泷奈）来实现灵活的内网穿透从而暴露靶机服务，请确保配置了 `/conf/takina_client.yaml` 文件，并且确保Takina服务端启动

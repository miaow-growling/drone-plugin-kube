# drone-k8s插件

```.env
  - name: publish
    image: plugins/docker
    settings:
      dockerfile: ./build/Dockerfile
      registry: registry.cn-shenzhen.aliyuncs.com
      repo: registry.cn-shenzhen.aliyuncs.com/yuanshuai/drone-ys-kube
      username:
        from_secret: registry_username
      password:
        from_secret: registry_password
      tags:
        - ${DRONE_BUILD_NUMBER}
```

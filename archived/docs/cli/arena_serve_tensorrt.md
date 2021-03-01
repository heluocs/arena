## arena serve tensorrt

Submit tensorRT inference serving job to deploy and serve machine learning models.

### Synopsis

Submit tensorRT inference serving job to deploy and serve machine learning models.

```
arena serve tensorrt [flags]
```

### Options

```
      --allowMetrics             Open Metric
      --command string           the command will inject to container's command.
      --cpu string               the request cpu of each replica to run the serve.
  -d, --data stringArray         specify the trained models datasource to mount for serving, like <name_of_datasource>:<mount_point_on_job>
      --enableIstio              enable Istio for serving or not (disable Istio by default)
  -e, --envs stringArray         the environment variables
      --exposeService            expose service using Istio gateway for external access or not (not expose by default)
      --gpumemory int            the limit GPU memory of each replica to run the serve.
      --gpus int                 the limit GPU count of each replica to run the serve.
      --grpcPort int             the port of grpc serving server (default 8001)
  -h, --help                     help for tensorrt
      --httpPort int             the port of http serving server (default 8000)
      --image string             the docker image name of serve job, and the default image is registry.cn-beijing.aliyuncs.com/xiaozhou/tensorrt-serving:18.12-py3 (default "registry.cn-beijing.aliyuncs.com/xiaozhou/tensorrt-serving:18.12-py3")
      --imagePullPolicy string   the policy to pull the image, and the default policy is IfNotPresent (default "IfNotPresent")
      --memory string            the request memory of each replica to run the serve.
      --metricPort int           the port of metrics server (default 8002)
      --modelName string         the model name for serving
      --modelPath string         the model path for serving in the container
      --modelStore string        the path of tensorRT model path
      --replicas int             the replicas number of the serve job. (default 1)
      --servingName string       the serving name
      --servingVersion string    the serving version
```

### Options inherited from parent commands

```
      --arena-namespace string   The namespace of arena system service, like tf-operator (default "arena-system")
      --config string            Path to a kube config. Only required if out-of-cluster
      --loglevel string          Set the logging level. One of: debug|info|warn|error (default "info")
  -n, --namespace string         the namespace of the job (default "default")
      --pprof                    enable cpu profile
      --trace                    enable trace
```

### SEE ALSO

* [arena serve](arena_serve.md)	 - Serve a job.

###### Auto generated by spf13/cobra on 24-Apr-2019
## arena submit

Submit a training job.

### Synopsis

Submit a job.

Available Commands:
  tfjob,tf             Submit a TFJob.
  pytorchjob,pytorch   Submit a PyTorchJob.
  mpijob,mpi           Submit a MPIJob.
  etjob,et             Submit a ETJob.
  horovod,hj           Submit a Horovod Job.
  volcanojob,vj        Submit a VolcanoJob.
    

```
arena submit [flags]
```

### Options

```
  -h, --help   help for submit
```

### Options inherited from parent commands

```
      --arena-namespace string   The namespace of arena system service, like tf-operator (default "arena-system")
      --config string            Path to a kube config. Only required if out-of-cluster
      --loglevel string          Set the logging level. One of: debug|info|warn|error (default "info")
  -n, --namespace string         the namespace of the job
      --pprof                    enable cpu profile
      --trace                    enable trace
```

### SEE ALSO

* [arena](arena.md)	 - arena is the command line interface to Arena
* [arena submit etjob](arena_submit_etjob.md)	 - Submit ETJob as training job.
* [arena submit horovodjob](arena_submit_horovodjob.md)	 - Submit horovodjob as training job.
* [arena submit mpijob](arena_submit_mpijob.md)	 - Submit MPIjob as training job.
* [arena submit pytorchjob](arena_submit_pytorchjob.md)	 - Submit PyTorchJob as training job.
* [arena submit tfjob](arena_submit_tfjob.md)	 - Submit a TFJob as training job.
* [arena submit volcanojob](arena_submit_volcanojob.md)	 - Submit a Volcano job.

###### Auto generated by spf13/cobra on 5-Mar-2021

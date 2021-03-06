## arena

arena is the command line interface to Arena

### Synopsis

arena is the command line interface to Arena

```
arena [flags]
```

### Options

```
      --arena-namespace string   The namespace of arena system service, like tf-operator (default "arena-system")
      --config string            Path to a kube config. Only required if out-of-cluster
  -h, --help                     help for arena
      --loglevel string          Set the logging level. One of: debug|info|warn|error (default "info")
  -n, --namespace string         the namespace of the job
      --pprof                    enable cpu profile
      --trace                    enable trace
```

### SEE ALSO

* [arena attach](arena_attach.md)	 - Attach a training job and execute some commands
* [arena completion](arena_completion.md)	 - output shell completion code for the specified shell (bash or zsh)
* [arena data](arena_data.md)	 - manage data.
* [arena delete](arena_delete.md)	 - Delete a training job and its associated instances
* [arena get](arena_get.md)	 - Display a training job details
* [arena list](arena_list.md)	 - List all the training jobs
* [arena logs](arena_logs.md)	 - Print the logs of a training job
* [arena logviewer](arena_logviewer.md)	 - Display Log Viewer URL of a training job
* [arena prune](arena_prune.md)	 - Prune the history jobs
* [arena scalein](arena_scalein.md)	 - Scale in a training job.
* [arena scaleout](arena_scaleout.md)	 - Scale out a  elastic training job.
* [arena serve](arena_serve.md)	 - Serve a job.
* [arena submit](arena_submit.md)	 - Submit a training job.
* [arena top](arena_top.md)	 - Display Resource (GPU) usage.
* [arena version](arena_version.md)	 - Print version information

###### Auto generated by spf13/cobra on 5-Mar-2021

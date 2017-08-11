# Snap processor plugin - change detector
Snap plugin intended to detect changes in metrics

It's used in the [Snap framework](http://github.com:intelsdi-x/snap).

1. [Getting Started](#getting-started)
  * [System Requirements](#system-requirements)
  * [Installation](#installation)
  * [Configuration and Usage](#configuration-and-usage)
2. [Documentation](#documentation)
  * [Examples](#examples)
  * [Roadmap](#roadmap)
3. [Community Support](#community-support)
4. [Contributing](#contributing)
5. [License](#license-and-authors)
6. [Acknowledgements](#acknowledgements)

## Getting Started
### System Requirements 
* [golang 1.7+](https://golang.org/dl/) (needed only for building)

### Operating systems
All OSs currently supported by snap:
* Linux/amd64
* Darwin/amd64

### Installation
#### Download processor plugin binary:
You can get the pre-built binaries for your OS and architecture under the plugin's [release](https://github.com/intelsdi-x/snap-plugin-processor-change-detector/releases) page.  For Snap, check [here](https://github.com/intelsdi-x/snap/releases).


#### To build the plugin binary:
Fork https://github.com/intelsdi-x/snap-plugin-processor-change-detector

Clone repo into `$GOPATH/src/github.com/intelsdi-x/`:

```
$ git clone https://github.com/<yourGithubID>/snap-plugin-processor-change-detector.git
```

Build the plugin by running make within the cloned repo:
```
$ make
```
This builds the plugin in `./build/`

### Configuration and Usage
* Set up the [Snap framework](https://github.com/intelsdi-x/snap/blob/master/README.md#getting-started)

## Documentation

The intention of this plugin is to detect changes between current and previous values of metrics.
Current value of metric is set as a metric's data, previous value of metric is saved in tag for metric.

The plugin can be configured by following parameters:
- `rules` - regular expressions which are used to select metrics in which changes need to be detected, regular expressions are separeted by `|`.

Notice: Special characters in regular expressions needs to be escaped.

### Examples
In this example we run CPU collector, change detector processor and file publisher to write data into file.

Documentation for Snap collector cpu plugin can be found [here](https://github.com/intelsdi-x/snap-plugin-collector-cpu).

In one terminal window, open the snap daemon with log level 1 (`-l 1`) and disabled plugin signing check (`-t 0`):
```
$ snapteld -t 0 -l 1
```

In another terminal window:

Download and load collector, processor and publisher plugins
```
$ wget http://snap.ci.snap-telemetry.io/plugins/snap-plugin-collector-cpu/latest/linux/x86_64/snap-plugin-collector-cpu
$ wget http://snap.ci.snap-telemetry.io/plugins/snap-plugin-processor-change-detector/latest/linux/x86_64/snap-plugin-processor-change-detector
$ wget http://snap.ci.snap-telemetry.io/plugins/snap-plugin-publisher-file/latest/linux/x86_64/snap-plugin-publisher-file
$ chmod 755 snap-plugin-*
$ snaptel plugin load snap-plugin-collector-cpu
$ snaptel plugin load snap-plugin-processor-change-detector
$ snaptel plugin load snap-plugin-publisher-file
```

See available metrics for your system
```
$ snaptel metric list
```

Create a task file - see examplary task manifests in [examples/tasks](examples/tasks/):

```yaml

---
  version: 1
  schedule:
    type: "simple"
    interval: "1s"
  max-failures: 10
  workflow:
    collect:
      metrics:
        /intel/procfs/cpu/*/utilization_percentage: {}
      process:
        - plugin_name: "change-detector"
          config:
            rules : "/intel/procfs/cpu/.*/utilization_percentage"
          publish:
            - plugin_name: "file"
              config:
                file: "/tmp/cpu_change_detector.log"
```

Start task:
```
$ snaptel task create -t task.yaml
```

This data is published to a file `/tmp/cpu_change_detector.log`

To stop task:
```
$ snaptel task stop <task id>
```

### Roadmap
There isn't a current roadmap for this plugin, but it is in active development. As we launch this plugin, we do not have any outstanding requirements for the next release. If you have a feature request, please add it as an [issue](https://github.com/intelsdi-x/snap-plugin-processor-change-detector/issues/new) and/or submit a [pull request](https://github.com/intelsdi-x/snap-plugin-processor-change-detector/pulls).

## Community Support
This repository is one of **many** plugins in **Snap**, a powerful telemetry framework. See the full project at http://github.com/intelsdi-x/snap To reach out to other users, head to the [main framework](https://github.com/intelsdi-x/snap#community-support)

## Contributing
We love contributions!

There's more than one way to give back, from examples to blogs to code updates. See our recommended process in [CONTRIBUTING.md](CONTRIBUTING.md).

## License
[Snap](http://github.com:intelsdi-x/snap), along with this plugin, is an Open Source software released under the Apache 2.0 [License](LICENSE).

## Acknowledgements
* Author: [Katarzyna Kujawa](https://github.com/katarzyna-z/)

And **thank you!** Your contribution, through code and participation, is incredibly important to us.
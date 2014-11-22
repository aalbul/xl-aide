XL Aide
=======

XL Aide is the tool to import and export snapshots of XLD data to accelerate the process of issue reproduction.

Build
-----

You need to download the latest version of [Go SDK](https://golang.org/dl/).
Then you need to set your GOROOT, how to do it you can have a look [here](https://golang.org/doc/install).

For Linux/Mac

```
$ build.sh```

For Windows

```
$ build.cmd
```

Installation
-------------

You need to copy xl-aide binary file to the root of the XLD installation. In addition to that you need to create a **xla-config.yml** file in
your user home directory, you can take a default configuration from **xla-config-sample.yml**, there you need to configure the access to your
 Jira and credentials of the admin user of XLD.


## Usage

To get the full list of all parameters:

```
    ./xl-aide -h
```

You will see such result:

```
    Usage of ./xl-aide:
      -export=true: By default you are exporting
      -force=false: Export XLA package and replace the previous uploaded package
      -import=false: Imports the data for specified issue
      -issue="": Specify your Jira issue, i.e. -issue=DEPL-6501
      -restart=false: Restart the server after importing the XLA
```

If you want to export current snapshot of XLD you need to provide Jira issue number to which this snapshot will be attached.

```
    ./xl-aide -issue=DEPL-6000
```

In case of success upload you will see next output:

```
XLA attachment [/private/tmp/xl-deploy/xla-snapshot.zip] has been successfully uploaded.
```

If the snapshot was already uploaded before:

```
Jira issue DEPL-6000 already has XLA attachment
```

But if you are sure that you are really want to re-upload with the current snapshot:

```
./xl-aide -issue=DEPL-6000 -force
```

If you want import the snapshot:

```
./xl-aide -issue=DEPL-6000 -import
```

If you want import the snapshot and restart the XLD server:

```
./xl-aide -issue=DEPL-6000 -import -restart
```

If you will provide not existent Jira issue you see this:

```
./xl-aide -issue=BLA-2
Issue [BLA-2] has not been found
```

If you will try to import from Jira issue where xla snapshot is not attached you see this:

```
Nothing to import. XLA attachment for issue [DEPL-6888] has not been found.
```


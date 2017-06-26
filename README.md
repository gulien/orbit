<p align="center">
    <img src="https://user-images.githubusercontent.com/8983173/26898223-7187b060-4bcb-11e7-831b-7174ce586fc5.png" alt="orbit's logo" width="200" height="200" />
</p>
<h3 align="center">Orbit</h3>
<p align="center">A simple tool for running commands and generating files from templates</p>
<p align="center">
    <a href="https://travis-ci.org/gulien/orbit"><img src="https://img.shields.io/travis/gulien/orbit.svg?label=linux+build" alt="Travis CI"></a>
    <a href="https://ci.appveyor.com/project/gulien/orbit"><img src="https://img.shields.io/appveyor/ci/gulien/orbit.svg?label=windows+build" alt="AppVeyor"></a>
    <a href="https://godoc.org/github.com/gulien/orbit"><img src="https://godoc.org/github.com/gulien/orbit?status.svg" alt="GoDoc"></a>
    <a href="https://goreportcard.com/report/gulien/orbit"><img src="https://goreportcard.com/badge/github.com/gulien/orbit" alt="Go Report Card"></a>
</p>

---

Orbit started with the need to find a cross-platform alternative of `make`
and `sed -i` commands. As it does not aim to be as powerful as these two
commands, Orbit offers an elegant solution for running commands and generating
files from templates, whatever the platform you're using.

# Menu

* [Install](#install)
* [Generating a file from a template](#generating-a-file-from-a-template)
* [Defining and running commands](#defining-and-running-commands)

## Install

Download the latest release of Orbit from the [releases page](../../../releases).
You can get Orbit for a large range of OS and architecture.

The file you downloaded is a compressed archive. You'll need to extract the
Orbit binary and move it into a folder where you can execute it easily.

**Linux/MacOS:**

```
tar -xzf orbit*.tar.gz orbit
sudo mv ./orbit /usr/local/bin && chmod +x /usr/local/bin/orbit
```

**Windows:**

Right click on the file and choose *Extract All*.

Move the binary to a folder like `C:\Orbit`.
Then, add it in your Path system environment variables. Click
*System, Advanced system settings, Environment Variables*... and
open *Path* under *System variables*. Edit the *Variable value* by adding
the folder with the Orbit binary.

Alright, you're almost done :metal:! Let's check your installation by running:

```
orbit version
```

## Generating a file from a template

Orbit uses the *Go* package `html/template` under the hood to generate a
file from a template.

You'll find interesting information in the official documentation:

* https://golang.org/pkg/text/template/
* https://golang.org/pkg/html/template/

### Simple case

The command for generating a file is quite simple:

```
orbit generate -t=the_path_of_your_template -o=the_path_of_the_resulting_file
```

If no output is specified, Orbit will print the result to *Stdout*.

### Providing data

Of course, you're able to tell Orbit where to find data which will be applied
to the template.

**YAML files:**

The flag `-v` allows you to specify one or many *YAML* files:

```
orbit generate [...] -v=the_path.yml
orbit generate [...] -v=key_1,file_1.yml
orbit generate [...] -v=key_1,file_1.yml;key_2,file_2.yml
```

As you can see, you're able to provide a basic mapping for your files:

* with mapping, your data will be accessible in your template through `{{ .Values.my_key.my_data }}`.
* otherwise through `{{ .Values.default.my_data }}`.

**.env files:**

The flag `-e` allows you to specify one or many *.env* files:

```
orbit generate [...] -e=.env
orbit generate [...] -e=key_1,.env_1
orbit generate [...] -e=key_1,.env_1;key_2,.env_2
```

As you can see, it works the same way as the `-v` flag:

* with mapping, your data will be accessible in your template through `{{ .EnvFiles.my_key.my_data }}`.
* otherwise through `{{ .EnvFiles.default.my_data }}`.

**Good to know:** you'll find interesting examples in the [assets folder](.assets).

## Defining and running commands

Like the `make` command with its `Makefile`, Orbit requires a
configuration file (by default, `orbit.yml`) where you define
your Orbit commands:

```
commands:
  - use: "my_first_command"
    run:
      - command [args]
      - command [args]
      - ...
  - use: "my_second_command"
    run:
      - command [args]
      - command [args]
      - ...
```

* the `use` attribute is the name of your Orbit command.
* the `run` attribute is the stack of external commands to run.
* an external command is a binary which is available in your `$PATH`.

Once you've defined your `orbit.yml` file, you're able
to run your Orbit command with:

```
orbit run my_first_command
orbit run my_second_command
orbit run my_first_command my_second_command
```

Notice that you may run nested Orbit commands :metal:!

### A configuration file as a template

A cool feature of Orbit is its ability to read its configuration through
a template.

For example, if you need to run a platform specific script, you may write:

```
commands:
  - use: "script"
    run:
    {{- if ne .Os "windows" }}
      - sh my_script.sh
    {{- else }}
      - cmd.exe /c .my_script.bat
    {{- end }}
```

There are two important things to notice:

1. Orbit provides the OS name at runtime with `{{ .Os }}` (you may find
all available names in the [official documentation](https://golang.org/doc/install/source#environment) - `$GOOS` column).
2. Adding a dash (e.g `{{-`) will not add break lines / spaces, otherwise
Orbit might fail to read your configuration file.

You may also use `-v` and `-e` flags for providing custom data.

If you need to specify a custom configuration file path (e.g different from `orbit.yml`),
you can provide it thanks to the `-c` flag.

---

Would you like to update this documentation ? Feel free to open an [issue](../../../issues).

<p align="center">
    <img src="https://user-images.githubusercontent.com/8983173/26898223-7187b060-4bcb-11e7-831b-7174ce586fc5.png" alt="orbit's logo" width="200" height="200" />
</p>
<h3 align="center">Orbit</h3>
<p align="center">A powerful task runner for executing commands and generating files from templates</p>
<p align="center">
    <a href="https://travis-ci.org/gulien/orbit">
        <img src="https://img.shields.io/travis/gulien/orbit.svg?label=linux+build" alt="Travis CI">
    </a>
    <a href="https://ci.appveyor.com/project/gulien/orbit">
        <img src="https://img.shields.io/appveyor/ci/gulien/orbit.svg?label=windows+build" alt="AppVeyor">
    </a>
    <a href="https://godoc.org/github.com/gulien/orbit">
        <img src="https://godoc.org/github.com/gulien/orbit?status.svg" alt="GoDoc">
    </a>
    <a href="https://goreportcard.com/report/gulien/orbit">
        <img src="https://goreportcard.com/badge/github.com/gulien/orbit" alt="Go Report Card">
    </a>
    <a href="https://codecov.io/gh/gulien/orbit">
        <img src="https://codecov.io/gh/gulien/orbit/branch/master/graph/badge.svg" alt="Codecov">
    </a>
</p>

---

Orbit started with the need to find a cross-platform alternative of `make`
and `sed -i` commands. As it does not aim to be as powerful as these two
commands, Orbit offers an elegant solution for running tasks and generating
files from templates, whatever the platform you're using.

# Menu

* [Install](#install)
* [Generating a file from a template](#generating-a-file-from-a-template)
* [Defining and running tasks](#defining-and-running-tasks)

## Install

Download the latest release of Orbit from the [releases page](../../releases).
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

Orbit uses the *Go* package `text/template` under the hood as a template
engine. It provides a interesting amount of logic for your templates.

The [Go documentation](https://golang.org/pkg/text/template/) and the
[Hugo documentation](http://gohugo.io/templates/go-templates/) cover
a lot of features that aren't mentioned here. Don't hesitate to take a look
at these links to understand the *Go* template engine! :smiley:

Also, Orbit provides [Sprig](http://masterminds.github.io/sprig/) library
and two custom functions:

* `os` which returns the current OS name at runtime (you may find all available names in the
[official documentation](https://golang.org/doc/install/source#environment)).
* `debug` which returns `true` if the `-d --debug` flag has been past to Orbit.

### Command description

#### Base

```
orbit generate [flags]
```

#### Flags

##### `-f --file`

Specify the path of the template. This flag is **required**.

##### `-o --output`

Specify the output file which will be generated from the template.

**Good to know:** if no output is specified, Orbit will print the result to *Stdout*.

##### `-p --payload`

The flag `-p` allows you to specify many data sources which will be applied to your template:

```
orbit generate [...] -p key_1,file_1.yml
orbit generate [...] -p key_1,file_1.yml;key_2,file_2.toml;key_3,file_3.json;key_4,.env;key_5,some raw data
```

As you can see, Orbit handles 5 types of data sources:

* *YAML* files (`*.yaml`, `*.yml`)
* *TOML* files (`*.toml`)
* *JSON* files (`*.json`)
* *.env* files
* raw data

The data will be accessible in your template through `{{ .Orbit.my_key.my_data }}`.

If you don't want to specify the payload each time your running `orbit generate`,
you may also create a file named `orbit-payload.yml` in the folder where your running your command:

```yaml
payload:

    - key: my_key
      value: my_file.yml
      
    - key: my_other_key
      value: Some raw data
```

By doing so, running `orbit generate [...]` will be equivalent to 
running `orbit generate [...] -p my_key,my_file.yml;my_other_key,Some raw data`.

**Note:** you are able to override a data source from the file `orbit-payload.yml` if
you set the same key in the `-p` flag.

##### `-d --debug`

Displays a detailed output.

### Basic example

Let's create our simple template `template.yml`:

```yaml
companies:
{{- range $company := .Orbit.Values.companies }}
  - name: {{ $company.name }}
    launchers:
  {{- range $launcher := $company.launchers }}
    - {{ $launcher }}
  {{- end }}
{{- end }}
```

And the data provided a *YAML* file named `data-source.yml`:

```yaml
companies:
  - name: SpaceX
    launchers:
      - Falcon 9
      - Falcon Heavy
  - name: Blue Origin
    launchers:
      - New Shepard
      - New Glenn

agencies:
  - name: ESA
    launchers:
      - Ariane 5
      - Vega
```

The command for generating a file from this template is quite simple:

```
orbit generate -f template.yml -p Values,data-source.yml -o companies.yml
```

This command will create the `companies.yml` file with this content:

```yaml
companies:
  - name: SpaceX
    launchers:
      - Falcon 9
      - Falcon Heavy
  - name: Blue Origin
    launchers:
      - New Shepard
      - New Glenn
```

## Defining and running tasks

### Command description

#### Base

```
orbit run [tasks] [flags]
```

#### Flags

##### `-f --file`

Like the `make` command with its `Makefile`, Orbit requires a
configuration file (*YAML*, by default `orbit.yml`) where you define
your tasks:

```yaml
tasks:
  - use: my_first_task
    short: My first task short description
    run:
      - command [args]
      - command [args]
      - ...
  - use: my_second_task
    short: My second task short description
    private: true
    run:
      - command [args]
      - command [args]
      - ...
```

* the `use` attribute is the name of your task.
* the `short` attribute is optional and is displayed when running `orbit run`
* the `private` attribute is optional and hides the considered task when running `orbit run`
* the `run` attribute is the stack of commands to run.
* a command is a binary which is available in your `$PATH`.

Once you've created your `orbit.yml` file, you're able
to run your tasks with:

```
orbit run my_first_task
orbit run my_second_task
orbit run my_first_task my_second_task
```

Notice that you may run nested tasks :metal:!

Also a cool feature of Orbit is its ability to read its configuration through
a template.

For example, if you need to execute a platform specific script, you may write:

```yaml
tasks:
  - use: script
    run:
    {{ if ne "windows" os }}
      - my_script.sh
    {{ else }}
      - .\my_script.bat
    {{ end }}
```

**Note:** Orbit will automatically detect the shell you're using. 
Running the task `script` from the previous example will in fact executes `cmd.exe /c \.my_script.bat` on
Windows or `/bin/sh -c my_script.sh` (or `/bin/zsh -c my_script.sh` etc.) on others OS.

##### `-p --payload`

The flag `-p` allows you to specify many data sources which will be applied to your configuration file.

It works the same as the `-p` flag from the `generate` command.

Of course, you may also create a file named `orbit-payload.yml` in the same folder as your configuration file.

##### `-d --debug`

Displays a detailed output.

### Basic example

Let's create our simple configuration file `orbit.yml`:

```yaml
tasks:
  - use: prepare
    run:
     - orbit generate -f configuration.template.yml -o configuration.yml -p Data,config.json
     - echo "configuration.yml has been succesfully created!"
```

You are now able to run the task `prepare` with:

```
orbit run prepare
```

This task will:

* create a file named `configuration.yml`
* print `configuration.yml has been succesfully created!`


Voilà! :smiley:

---

Would you like to update this documentation ? Feel free to open an [issue](../../issues).

# i3-layout-per-workspace

This tool will allow you to force a layout on a workspace.

## Requirements

You must have go 1.17+ installed.

also, this tool uses the official golang i3 IPC interface, which comes with two requirements:

> * The i3(1) binary must be in $PATH so that the IPC socket path can be retrieved.
> * For transparent version checks to work, the running i3 version must be ≥ 4.3 (released 2012-09-19).

https://github.com/i3/go-i3#assumptions 

## Setup

Feel free to change the path of the clone to wherever your want :)

- clone the repostiory => `git clone https://github.com/JulienTant/i3-layout-per-workspace.git ~/.config/i3/i3-layout-per-workspace`
- compile the code (`make` or using `go build`) => `cd ~/.config/i3/i3-layout-per-workspace; make`
  - it will create the file binary `~/.config/i3/layout-per-ws/layout-per-ws` 
- add a `exec` or `exec_always` command in your i3 config file (`~/.config/i3/config`) like so:
```shell
exec_always --no-startup-id ~/.config/i3/layout-per-ws/layout-per-ws "your workspace name" tabbed "second workspace" stacked
```
it also works with variables
```shell
set $ws_browser "Browser"
set $ws_code "IDE"
exec_always --no-startup-id ~/.config/i3/layout-per-ws/layout-per-ws $ws_code stacked $ws_messaging tabbed
```

## Usage

You will have to choose between the 4 different layouts offer by i3: `splith`, `splitv`, `stacked` or `tabbed`

For each workspace that you want to force on a certain layout, you will have to add the pair `"workspace name" "layout"` as an argument to the program.

So if I have a workspace `Messaging` that I always want `stacked`, and a workspace `Code Zone` that I always want `tabbed`. I will run: `layout-per-ws "Messaging" "stacked" "Code Zone" "tabbed"`

While it's not mandatory to have each argument quoted if they don't contain a space or a unicode character, it will make your life easier if something does not work.

The software also log to STDOUT a few stuff, so feel free to try it out of i3 before adding it in the config!

```shell
$ ./layout-per-ws " Messaging" "tabbed" " IDE" "stacked"

2021/12/19 20:54:48 building pairs
2021/12/19 20:54:48  Messaging => tabbed
2021/12/19 20:54:48  IDE => stacked
```
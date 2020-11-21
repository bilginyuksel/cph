# cordova-plugin-helper

This tool's aim is to make cordova plugin development smooth, fast and stable. It will help you from the beginning to the end of the plugin development. It will automatically create plugin files. Then you can start developing your applications. Whenever you want to sync your plugin.xml you can sync the plugin.xml with the files you have created.

### Commands

1. Global `--help` command.
```bash
Usage: cph.exe <command>

Flags:
  -h, --help     Show context-sensitive help.
      --debug    Enable debug mode.

Commands:
  plugin --project-name=STRING --domain=STRING --home-page=STRING
    Use this command to create a cordova plugin from scratch.

  plugin-xml
    You can use the functions in that command to manipulate plugin.xml file
    under cordova plugin root project directory.

  add-license
    Add license to files.

Run "cph.exe <command> --help" for more information on a command.
```

2. Detailed help commands

You can type `--help` for every sub command. Help commands will help you to use the application. Example commands:

> `cph add-license --help`

> `cph plugin-xml --help`

> `cph plugin --help`


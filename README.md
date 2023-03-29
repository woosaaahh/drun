## drun - detached run

`drun` is a little tool to run any commands in a 'detached' way, _i.e._ without waiting for them to finish.

**USAGE :** `drun program [arg...]`

```sh
drun firefox
drun libreoffice <path/to/document>
drun vlc <path/to/video>
```

**IMPORTANT NOTES :**

- This program does only one thing, running a 'detached' command ;
- You cannot use the stdin/stdout/stderr of the ran command, even if there is an error ;
- You cannot get the exit code of the ran command.

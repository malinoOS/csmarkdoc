# csmarkdoc
Generate markdown documentation for C# code. This tool parses your code and looks for XML comments.

It currently does not support features in XML comments like `<see>`. But that will be added down the road.

[See what it can do](https://github.com/malinoOS/malino/wiki/libmalino-(C%23))

## Limitations
- You must only have ONE namespace defined in the entire file list.
- It doesn't quite understand the concept of enums yet.

## Usage

`csmarkdoc` outputs markdown through stdout. You can pass this to a file by adding ` > file.md` to the end of a command, for example, `csmarkdoc file.cs > file.md`.

```
csmarkdoc <files>
```

If you want to parse just one file, then just do this:

```
csmarkdoc file.cs
```

If you want to parse all C# files, just do this:

```
csmarkdoc *.cs
```
# csmarkdoc
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/malinoOS/csmarkdoc?style=for-the-badge&logo=files)


A very tiny markdown documentation generator for C# code. This tool parses your code and looks for XML comments.

It currently does not support features in XML comments like `<see>`. But that will be added down the road.

[See what it can do](https://github.com/malinoOS/malino/wiki/libmalino-(C%23))

## Limitations
- You must only have ONE namespace defined in the entire file list.
- It doesn't quite understand the concept of enums yet.

## Usage

To use it properly (You can use `csmarkdoc` without adding comments, but it's not recommended), you have to put XML comments in your code, like the `///` ones, find information about them [here.](https://learn.microsoft.com/en-us/dotnet/csharp/language-reference/xmldoc/)

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

# goldmark-superscript

[GoldMark](https://github.com/yuin/goldmark/) superscript extension.

This implements the [`superscript`](https://pandoc.org/MANUAL.html#extension-superscript-subscript) of pandoc.

```markdown
H~2~O is a liquid.  2^10^ is 1024.
```

```html
<p>H~2~O is a liquid.  2<sup>10</sup> is 1024.</p>
```

```go
var md = goldmark.New(superscript.Enable)
var source = []byte("H~2~O is a liquid.  2^10^ is 1024.")
err := md.Convert(source, os.Stdout)
if err != nil {
    log.Fatal(err)
}
```

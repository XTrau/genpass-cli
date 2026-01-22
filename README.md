# genpass

`genpass` — простой CLI-генератор паролей, написанный на Go.

### Флаги для генерации

- `--size` установка длины генерируемого пароля (по умолчанию 8)
- строчных букв английского алфавита (`-a`)
- заглавных букв английского алфавита (`-A`)
- цифры (`-n`)

## Примеры использования

```
$ genpass
> onwrjucq
```

```
$ genpass --size=15 
> oxfyrawyhgeerru
```

```
$ genpass --size=20 -n -a
> xjifnjh3u4wl2qkmfpp6
```

```
$ genpass --size=50 -n -a -A
> AKrLGmxisFvApzgF6ZMpLe208ZC49e8bJBK57aOdrP7F8SiMM1
```
# Hamlet64 

##### The world's fastest *(lol, only)* Shakespearean binary-to-text encoder.

**Practical?** No. <br>
**Useful?** Probably not. <br>
**Dramatic?** O yes, my lord!

### Usage

```bash
# run it directly
go run . <filename>

# or build it
go build -o hamlet64 .
./hamlet64 <filename>
```

### Install globally

```bash
go build -o hamlet64 .
sudo mv hamlet64 /usr/local/bin/

# now use it from anywhere
hamlet64 <filename>
```

### Quick test

```bash
echo -n "Hi" > test.txt
go run . test.txt
```

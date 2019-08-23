# uml2image
uml2image is discord bot to convert uml to image.

# Install
Install JRE(or JDK) and Graphviz in advance.

clone uml2image.
```bash
git clone github.com/minami14/uml2image
```

build uml2image.
```bash
cd uml2image
go build -o uml2image cmd/main.go
```

Save plantuml.jar in the same directory as the built binary.
```bash
wget -O plantuml.jar http://sourceforge.net/projects/plantuml/files/plantuml.jar/download
```

# Usage
```bash
./uml2image [Your token]
```

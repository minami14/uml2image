# uml2image
uml2image is discord bot to convert uml to image.

# Install
### Docker
```bash
docker run -d -e Token[Your token] minami14/uml2image
```

### Binary
Install JRE(or JDK) and Graphviz.
```bash
sudo apt install default-jre graphviz
```

clone this repository.
```bash
git clone https://github.com/minami14/uml2image
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

Write uml with discord as follows
````    
```uml
@startuml
Class01 <|-- Class02
Class03 *-- Class04
Class05 o-- Class06
Class07 .. Class08
Class09 -- Class10
@enduml
```
````

![example](https://raw.githubusercontent.com/minami14/uml2image/master/example.png)

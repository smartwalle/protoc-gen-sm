package grpc

import (
	"fmt"
	"github.com/golang/protobuf/protoc-gen-go/generator"
	"strings"
)

const generatedCodeVersion = 4

func init() {
	generator.RegisterPlugin(new(oneof))
}

type oneof struct {
	gen *generator.Generator
}

func (g *oneof) Name() string {
	return "oneof"
}

var (
	contextPkg string
	grpcPkg    string
)

func (g *oneof) Init(gen *generator.Generator) {
	g.gen = gen
}

func (g *oneof) objectNamed(name string) generator.Object {
	g.gen.RecordTypeUse(name)
	return g.gen.ObjectNamed(name)
}

func (g *oneof) typeName(str string) string {
	return g.gen.TypeName(g.objectNamed(str))
}

func (g *oneof) P(args ...interface{}) { g.gen.P(args...) }

func (g *oneof) Generate(file *generator.FileDescriptor) {

	var packageName = *file.Package

	for _, mt := range file.FileDescriptorProto.MessageType {
		for _, od := range mt.Field {
			if od.OneofIndex != nil {

				var methodName = fmt.Sprintf("New%s%s() ", mt.GetName(), od.GetName())
				var fieldType = fmt.Sprintf("%s_%s", mt.GetName(), od.GetName())

				var typeName = strings.TrimPrefix(od.GetTypeName(), ".")

				if strings.HasPrefix(typeName, packageName+".") {
					typeName = strings.TrimPrefix(typeName, packageName+".")
				}

				g.P("func ", methodName, "*", fieldType ," {")
				g.P("var m = ", "&", fieldType, "{}")
				g.P("m.", od.GetName(), " = &", typeName+"{}")
				g.P("return m")
				g.P("}")
				g.P()
			}
		}

	}
}

func (g *oneof) GenerateImports(file *generator.FileDescriptor) {
}

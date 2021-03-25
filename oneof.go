/*
 *
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package main

import (
	"fmt"
	"google.golang.org/protobuf/compiler/protogen"
	"path/filepath"
	"strings"
)

// generateOneof generates a _grpc.pb.go file containing gRPC service definitions.
func generateOneof(gen *protogen.Plugin, file *protogen.File) *protogen.GeneratedFile {
	filename := file.GeneratedFilenamePrefix + "_oneof.pb.go"
	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	g.P("// Code generated by protoc-gen-sm. DO NOT EDIT.")
	g.P()
	g.P("package ", file.GoPackageName)
	g.P()

	for _, message := range file.Messages {
		for _, oneof := range message.Oneofs {

			ifName := message.GoIdent.GoName + oneof.GoName
			g.P("// 需要手动删除 ", file.GeneratedFilenamePrefix+".pb.go 中的 ", ifName, " 接口声明")
			g.P("type ", ifName, " interface {")
			g.P(oneofInterfaceName(oneof))
			g.P()
			g.P("Number() int32")
			g.P()
			g.P("String() string")
			g.P()
			g.P("Filename() string")
			g.P("}")
			g.P()

			g.P("func (x *", message.GoIdent.GoName, ") Get", ifName, "()", ifName, "{")
			g.P("return x.", oneof.GoName, ".(", ifName, ")")
			g.P("}")
			g.P()

			for _, field := range oneof.Fields {
				var methodName = fmt.Sprintf("New%s%s() ", message.GoIdent.GoName, field.Message.GoIdent.GoName)
				g.P("func ", methodName, "*", field.GoIdent, "{")
				g.P("var m = &", field.GoIdent, "{}")
				g.P("m.", field.GoName, "= &", g.QualifiedGoIdent(field.Message.GoIdent), "{}")
				g.P("return m")
				g.P("}")
				g.P()

				g.P("func (*", field.GoIdent.GoName, ") Number() int32 {")
				g.P("return ", field.Desc.Number())
				g.P("}")
				g.P()

				g.P("func (*", field.GoIdent.GoName, ") String() string {")
				g.P("return \"", field.Message.GoIdent.GoName, "\"")
				g.P("}")
				g.P()

				g.P("func (*", field.GoIdent.GoName, ") Filename() string {")
				g.P("return \"", strings.Split(filepath.Base(field.Message.Location.SourceFile), ".")[0], "\"")
				g.P("}")
				g.P()
			}
		}
		//for _, od := range mt.Oneofs {
		//	if od.OneofIndex != nil {
		//
		//		var methodName = fmt.Sprintf("New%s%s() ", mt.GetName(), od.GetName())
		//		var fieldType = fmt.Sprintf("%s_%s", mt.GetName(), od.GetName())
		//
		//		var typeName = strings.TrimPrefix(od.GetTypeName(), ".")
		//
		//		if strings.HasPrefix(typeName, packageName+".") {
		//			typeName = strings.TrimPrefix(typeName, packageName+".")
		//		}
		//
		//		g.P("func ", methodName, "*", fieldType ," {")
		//		g.P("var m = ", "&", fieldType, "{}")
		//		g.P("m.", od.GetName(), " = &", typeName+"{}")
		//		g.P("return m")
		//		g.P("}")
		//		g.P()
		//	}
		//}
	}
	return g
}

func oneofInterfaceName(oneof *protogen.Oneof) string {
	return "is" + oneof.GoIdent.GoName
}

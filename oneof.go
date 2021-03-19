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
)

// generateOneOf generates a _grpc.pb.go file containing gRPC service definitions.
func generateOneOf(gen *protogen.Plugin, file *protogen.File) *protogen.GeneratedFile {
	if len(file.Services) == 0 {
		return nil
	}
	filename := file.GeneratedFilenamePrefix + "_oneof.pb.go"
	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	g.P("// Code generated by protoc-gen-sm. DO NOT EDIT.")
	g.P()
	g.P("package ", file.GoPackageName)
	g.P()

	for _, message := range file.Messages {
		for _, oneOf := range message.Oneofs {
			for _, field := range oneOf.Fields {
				var methodName = fmt.Sprintf("New%s%s() ", message.GoIdent.GoName, field.Message.GoIdent.GoName)
				var fieldType = fmt.Sprintf("%s_%s", message.GoIdent.GoName, field.Message.GoIdent.GoName)

				g.P("func ", methodName, "*", fieldType, "{")
				g.P("var m = &", fieldType, "{}")
				g.P("m.", field.Message.GoIdent.GoName, "= &", field.Message.GoIdent.GoName, "{}")
				g.P("return m")
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

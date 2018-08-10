//  Copyright (c) 2018 Minoru Osuka
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package registry

import (
	"fmt"
	"reflect"
)

type TypeRegistry map[string]reflect.Type

var Types = make(TypeRegistry, 0)

func RegisterType(name string, t reflect.Type) {
	if _, exists := Types[name]; exists {
		panic(fmt.Errorf("attempted to register duplicate index: %s", name))
	}
	Types[name] = t
}

func TypeByName(name string) reflect.Type {
	return Types[name]
}

func TypeNameByInstance(instance interface{}) string {
	return reflect.TypeOf(instance).Elem().String()
}

func TypeInstanceByName(name string) interface{} {
	return reflect.New(TypeByName(name)).Interface()
}

/*
Copyright © 2021 Andy Zhou

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	_ "github.com/z-t-y/flogo/cli/auth"
	_ "github.com/z-t-y/flogo/cli/column"
	_ "github.com/z-t-y/flogo/cli/comment"
	_ "github.com/z-t-y/flogo/cli/config"
	_ "github.com/z-t-y/flogo/cli/notification"
	_ "github.com/z-t-y/flogo/cli/post"
	_ "github.com/z-t-y/flogo/cli/self"
	"github.com/z-t-y/flogo/cmd"
)

func main() {
	cmd.Execute()
}

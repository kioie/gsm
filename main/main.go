/*
 * // Licensed to the Apache Software Foundation (ASF) under one
 * // or more contributor license agreements.  See the NOTICE file
 * // distributed with this work for additional information
 * // regarding copyright ownership.  The ASF licenses this file
 * // to you under the Apache License, Version 2.0 (the
 * // "License"); you may not use this file except in compliance
 * // with the License.  You may obtain a copy of the License at
 * //
 * //   http://www.apache.org/licenses/LICENSE-2.0
 * //
 * // Unless required by applicable law or agreed to in writing,
 * // software distributed under the License is distributed on an
 * // "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * // KIND, either express or implied.  See the License for the
 * // specific language governing permissions and limitations
 * // under the License.
 *
 *
 *
 *
 * author: Eddy Kioi
 * project: gcp-secret-manager
 * date: 15/06/2020, 14:12
 */

package main

import (
	"fmt"
	gcp_secret_manager "github.com/kioie/gcp-secret-manager"
)


func main() {
	gcp_secret_manager.ProjectID = "secret-manager-test"
	//projectID := "secret-manager-test"
	//connect :=&gcp_secret_manager.NewClient{}
	//	if !gcp_secret_manager.SecretExists("projects/1092054168008/secrets/my-secret") {
	//		fmt.Println("Secret does not Exists")
	//	}

	//	fmt.Print(gcp_secret_manager.CreateSecret(projectID, []byte("my secret data")))
	//result := gcp_secret_manager.GetSecret("my-secret", "", projectID)
	//fmt.Println(result)
	//	newResult := gcp_secret_manager.AddNewSecretVersion("my-secret", projectID, []byte("a new test"))
	//	fmt.Println(newResult)
	result2, _ := gcp_secret_manager.GetSecret("my-secret", "")
	fmt.Println(result2)
	//	result3 := gcp_secret_manager.GetSecretMetadata("my-secret", "1", projectID)
	//	fmt.Println(result3.GetName())

	//fmt.Println(gcp_secret_manager.SecretExists(projectID,"my-secret"))
}

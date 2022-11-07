/*
 * Copyright (c) 2019 WSO2 Inc. (http:www.wso2.org) All Rights Reserved.
 *
 * WSO2 Inc. licenses this file to you under the Apache License,
 * Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http:www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package siddhicontroller

// Status of a Siddhi process
type Status int

// Status array holds the string values of status
var status = []string{
	"Pending",
	"Ready",
	"Running",
	"Error",
	"Warning",
	"Normal",
	"Not Ready",
}

// getStatus return relevant status to a given integer. This uses status array and the constants list.
func getStatus(n Status) string {
	return status[n]
}

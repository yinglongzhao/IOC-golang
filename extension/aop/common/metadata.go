/*
 * Copyright (c) 2022, Alibaba Group;
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package common

import "github.com/alibaba/ioc-golang/autowire"

const AOPMetadataKey = "aop"

// AOPMetadata is aop extension metadata
type AOPMetadata map[string]interface{}

func ParseAOPMetadataFromSDMetadata(metadata autowire.Metadata) AOPMetadata {
	if metadata == nil {
		return nil
	}
	if aopMetadataVal, ok := metadata[AOPMetadataKey]; ok {
		if aopMetadata, ok2 := aopMetadataVal.(map[string]interface{}); ok2 {
			return aopMetadata
		}
	}
	return nil
}

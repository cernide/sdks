// Copyright 2018-2023 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/* tslint:disable */
/* eslint-disable */
/**
 * Polyaxon SDKs and REST API specification.
 * Polyaxon SDKs and REST API specification.
 *
 * The version of the OpenAPI document: 1.21.0
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

/**
 * - replace: Replaces the keys
 *  - isnull: Only set the keys if they do not exist or if they are null
 *  - post_merge: Merge the all keys and replace by new one
 *  - pre_merge: Merge the all keys and keep old ones
 * @export
 * @enum {string}
 */
export enum V1PatchStrategy {
    Replace = 'replace',
    Isnull = 'isnull',
    PostMerge = 'post_merge',
    PreMerge = 'pre_merge'
}

export function V1PatchStrategyFromJSON(json: any): V1PatchStrategy {
    return V1PatchStrategyFromJSONTyped(json, false);
}

export function V1PatchStrategyFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1PatchStrategy {
    return json as V1PatchStrategy;
}

export function V1PatchStrategyToJSON(value?: V1PatchStrategy | null): any {
    return value as any;
}


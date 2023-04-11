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
 * The version of the OpenAPI document: 2.0.0-rc5
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


/**
 *
 * @export
 */
export const V1RunPending = {
    Approval: 'approval',
    Upload: 'upload',
    Cache: 'cache',
    Build: 'build'
} as const;
export type V1RunPending = typeof V1RunPending[keyof typeof V1RunPending];


export function V1RunPendingFromJSON(json: any): V1RunPending {
    return V1RunPendingFromJSONTyped(json, false);
}

export function V1RunPendingFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1RunPending {
    return json as V1RunPending;
}

export function V1RunPendingToJSON(value?: V1RunPending | null): any {
    return value as any;
}


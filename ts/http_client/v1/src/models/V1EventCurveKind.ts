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
 * The version of the OpenAPI document: 1.22.0-rc0
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

/**
 * - roc: ROC curve
 *  - pr: Precision Recall curve
 *  - custom: Custom curve
 * @export
 * @enum {string}
 */
export enum V1EventCurveKind {
    Roc = 'roc',
    Pr = 'pr',
    Custom = 'custom'
}

export function V1EventCurveKindFromJSON(json: any): V1EventCurveKind {
    return V1EventCurveKindFromJSONTyped(json, false);
}

export function V1EventCurveKindFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1EventCurveKind {
    return json as V1EventCurveKind;
}

export function V1EventCurveKindToJSON(value?: V1EventCurveKind | null): any {
    return value as any;
}


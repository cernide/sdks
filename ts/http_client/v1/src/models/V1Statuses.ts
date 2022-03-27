// Copyright 2018-2021 Polyaxon, Inc.
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
 * The version of the OpenAPI document: 1.17.1
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

/**
 * 
 * @export
 * @enum {string}
 */
export enum V1Statuses {
    Created = 'created',
    Resuming = 'resuming',
    OnSchedule = 'on_schedule',
    Compiled = 'compiled',
    Queued = 'queued',
    Scheduled = 'scheduled',
    Starting = 'starting',
    Running = 'running',
    Processing = 'processing',
    Stopping = 'stopping',
    Failed = 'failed',
    Stopped = 'stopped',
    Succeeded = 'succeeded',
    Skipped = 'skipped',
    Warning = 'warning',
    Unschedulable = 'unschedulable',
    UpstreamFailed = 'upstream_failed',
    Retrying = 'retrying',
    Unknown = 'unknown',
    Done = 'done'
}

export function V1StatusesFromJSON(json: any): V1Statuses {
    return V1StatusesFromJSONTyped(json, false);
}

export function V1StatusesFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1Statuses {
    return json as V1Statuses;
}

export function V1StatusesToJSON(value?: V1Statuses | null): any {
    return value as any;
}


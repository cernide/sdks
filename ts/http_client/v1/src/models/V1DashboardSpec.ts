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
 * The version of the OpenAPI document: 1.17.0
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    V1SectionSpec,
    V1SectionSpecFromJSON,
    V1SectionSpecFromJSONTyped,
    V1SectionSpecToJSON,
} from './';

/**
 * 
 * @export
 * @interface V1DashboardSpec
 */
export interface V1DashboardSpec {
    /**
     * 
     * @type {Array<V1SectionSpec>}
     * @memberof V1DashboardSpec
     */
    sections?: Array<V1SectionSpec>;
}

export function V1DashboardSpecFromJSON(json: any): V1DashboardSpec {
    return V1DashboardSpecFromJSONTyped(json, false);
}

export function V1DashboardSpecFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1DashboardSpec {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'sections': !exists(json, 'sections') ? undefined : ((json['sections'] as Array<any>).map(V1SectionSpecFromJSON)),
    };
}

export function V1DashboardSpecToJSON(value?: V1DashboardSpec | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'sections': value.sections === undefined ? undefined : ((value.sections as Array<any>).map(V1SectionSpecToJSON)),
    };
}


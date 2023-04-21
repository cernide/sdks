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
 * The version of the OpenAPI document: 2.0.0-rc10
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 *
 * @export
 * @interface V1RunReferenceCatalog
 */
export interface V1RunReferenceCatalog {
    /**
     *
     * @type {string}
     * @memberof V1RunReferenceCatalog
     */
    owner?: string;
    /**
     *
     * @type {string}
     * @memberof V1RunReferenceCatalog
     */
    project?: string;
    /**
     *
     * @type {string}
     * @memberof V1RunReferenceCatalog
     */
    name?: string;
}

/**
 * Check if a given object implements the V1RunReferenceCatalog interface.
 */
export function instanceOfV1RunReferenceCatalog(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function V1RunReferenceCatalogFromJSON(json: any): V1RunReferenceCatalog {
    return V1RunReferenceCatalogFromJSONTyped(json, false);
}

export function V1RunReferenceCatalogFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1RunReferenceCatalog {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {

        'owner': !exists(json, 'owner') ? undefined : json['owner'],
        'project': !exists(json, 'project') ? undefined : json['project'],
        'name': !exists(json, 'name') ? undefined : json['name'],
    };
}

export function V1RunReferenceCatalogToJSON(value?: V1RunReferenceCatalog | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {

        'owner': value.owner,
        'project': value.project,
        'name': value.name,
    };
}


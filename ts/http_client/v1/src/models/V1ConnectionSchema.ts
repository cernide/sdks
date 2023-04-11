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

import { exists, mapValues } from '../runtime';
import type { V1BucketConnection } from './V1BucketConnection';
import {
    V1BucketConnectionFromJSON,
    V1BucketConnectionFromJSONTyped,
    V1BucketConnectionToJSON,
} from './V1BucketConnection';
import type { V1ClaimConnection } from './V1ClaimConnection';
import {
    V1ClaimConnectionFromJSON,
    V1ClaimConnectionFromJSONTyped,
    V1ClaimConnectionToJSON,
} from './V1ClaimConnection';
import type { V1GitConnection } from './V1GitConnection';
import {
    V1GitConnectionFromJSON,
    V1GitConnectionFromJSONTyped,
    V1GitConnectionToJSON,
} from './V1GitConnection';
import type { V1HostConnection } from './V1HostConnection';
import {
    V1HostConnectionFromJSON,
    V1HostConnectionFromJSONTyped,
    V1HostConnectionToJSON,
} from './V1HostConnection';
import type { V1HostPathConnection } from './V1HostPathConnection';
import {
    V1HostPathConnectionFromJSON,
    V1HostPathConnectionFromJSONTyped,
    V1HostPathConnectionToJSON,
} from './V1HostPathConnection';

/**
 *
 * @export
 * @interface V1ConnectionSchema
 */
export interface V1ConnectionSchema {
    /**
     *
     * @type {V1BucketConnection}
     * @memberof V1ConnectionSchema
     */
    bucketConnection?: V1BucketConnection;
    /**
     *
     * @type {V1HostPathConnection}
     * @memberof V1ConnectionSchema
     */
    hostPathConnection?: V1HostPathConnection;
    /**
     *
     * @type {V1ClaimConnection}
     * @memberof V1ConnectionSchema
     */
    claimConnection?: V1ClaimConnection;
    /**
     *
     * @type {V1HostConnection}
     * @memberof V1ConnectionSchema
     */
    hostConnection?: V1HostConnection;
    /**
     *
     * @type {V1GitConnection}
     * @memberof V1ConnectionSchema
     */
    gitConnection?: V1GitConnection;
}

/**
 * Check if a given object implements the V1ConnectionSchema interface.
 */
export function instanceOfV1ConnectionSchema(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function V1ConnectionSchemaFromJSON(json: any): V1ConnectionSchema {
    return V1ConnectionSchemaFromJSONTyped(json, false);
}

export function V1ConnectionSchemaFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1ConnectionSchema {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {

        'bucketConnection': !exists(json, 'bucketConnection') ? undefined : V1BucketConnectionFromJSON(json['bucketConnection']),
        'hostPathConnection': !exists(json, 'hostPathConnection') ? undefined : V1HostPathConnectionFromJSON(json['hostPathConnection']),
        'claimConnection': !exists(json, 'claimConnection') ? undefined : V1ClaimConnectionFromJSON(json['claimConnection']),
        'hostConnection': !exists(json, 'hostConnection') ? undefined : V1HostConnectionFromJSON(json['hostConnection']),
        'gitConnection': !exists(json, 'gitConnection') ? undefined : V1GitConnectionFromJSON(json['gitConnection']),
    };
}

export function V1ConnectionSchemaToJSON(value?: V1ConnectionSchema | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {

        'bucketConnection': V1BucketConnectionToJSON(value.bucketConnection),
        'hostPathConnection': V1HostPathConnectionToJSON(value.hostPathConnection),
        'claimConnection': V1ClaimConnectionToJSON(value.claimConnection),
        'hostConnection': V1HostConnectionToJSON(value.hostConnection),
        'gitConnection': V1GitConnectionToJSON(value.gitConnection),
    };
}


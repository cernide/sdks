/* tslint:disable */
/* eslint-disable */
/**
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.3.3
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { V1ArtifactKind } from './V1ArtifactKind';
import {
    V1ArtifactKindFromJSON,
    V1ArtifactKindFromJSONTyped,
    V1ArtifactKindToJSON,
} from './V1ArtifactKind';

/**
 *
 * @export
 * @interface V1RunArtifact
 */
export interface V1RunArtifact {
    /**
     *
     * @type {string}
     * @memberof V1RunArtifact
     */
    name?: string;
    /**
     *
     * @type {string}
     * @memberof V1RunArtifact
     */
    state?: string;
    /**
     *
     * @type {V1ArtifactKind}
     * @memberof V1RunArtifact
     */
    kind?: V1ArtifactKind;
    /**
     *
     * @type {string}
     * @memberof V1RunArtifact
     */
    path?: string;
    /**
     *
     * @type {string}
     * @memberof V1RunArtifact
     */
    connection?: string;
    /**
     *
     * @type {string}
     * @memberof V1RunArtifact
     */
    run?: string;
    /**
     *
     * @type {object}
     * @memberof V1RunArtifact
     */
    summary?: object;
    /**
     *
     * @type {boolean}
     * @memberof V1RunArtifact
     */
    is_input?: boolean;
    /**
     *
     * @type {object}
     * @memberof V1RunArtifact
     */
    meta_info?: object;
}

/**
 * Check if a given object implements the V1RunArtifact interface.
 */
export function instanceOfV1RunArtifact(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function V1RunArtifactFromJSON(json: any): V1RunArtifact {
    return V1RunArtifactFromJSONTyped(json, false);
}

export function V1RunArtifactFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1RunArtifact {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {

        'name': !exists(json, 'name') ? undefined : json['name'],
        'state': !exists(json, 'state') ? undefined : json['state'],
        'kind': !exists(json, 'kind') ? undefined : V1ArtifactKindFromJSON(json['kind']),
        'path': !exists(json, 'path') ? undefined : json['path'],
        'connection': !exists(json, 'connection') ? undefined : json['connection'],
        'run': !exists(json, 'run') ? undefined : json['run'],
        'summary': !exists(json, 'summary') ? undefined : json['summary'],
        'is_input': !exists(json, 'is_input') ? undefined : json['is_input'],
        'meta_info': !exists(json, 'meta_info') ? undefined : json['meta_info'],
    };
}

export function V1RunArtifactToJSON(value?: V1RunArtifact | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {

        'name': value.name,
        'state': value.state,
        'kind': V1ArtifactKindToJSON(value.kind),
        'path': value.path,
        'connection': value.connection,
        'run': value.run,
        'summary': value.summary,
        'is_input': value.is_input,
        'meta_info': value.meta_info,
    };
}


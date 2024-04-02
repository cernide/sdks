/* tslint:disable */
/* eslint-disable */
/**
 * Polyaxon SDKs and REST API specification.
 *    
 *
 * The version of the OpenAPI document: 2.1.5
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { V1ArtifactsType } from './V1ArtifactsType';
import {
    V1ArtifactsTypeFromJSON,
    V1ArtifactsTypeFromJSONTyped,
    V1ArtifactsTypeToJSON,
} from './V1ArtifactsType';
import type { V1DockerfileType } from './V1DockerfileType';
import {
    V1DockerfileTypeFromJSON,
    V1DockerfileTypeFromJSONTyped,
    V1DockerfileTypeToJSON,
} from './V1DockerfileType';
import type { V1FileType } from './V1FileType';
import {
    V1FileTypeFromJSON,
    V1FileTypeFromJSONTyped,
    V1FileTypeToJSON,
} from './V1FileType';
import type { V1GitType } from './V1GitType';
import {
    V1GitTypeFromJSON,
    V1GitTypeFromJSONTyped,
    V1GitTypeToJSON,
} from './V1GitType';
import type { V1TensorboardType } from './V1TensorboardType';
import {
    V1TensorboardTypeFromJSON,
    V1TensorboardTypeFromJSONTyped,
    V1TensorboardTypeToJSON,
} from './V1TensorboardType';

/**
 * 
 * @export
 * @interface V1Init
 */
export interface V1Init {
    /**
     * 
     * @type {V1ArtifactsType}
     * @memberof V1Init
     */
    artifacts?: V1ArtifactsType;
    /**
     * 
     * @type {Array<object>}
     * @memberof V1Init
     */
    paths?: Array<object>;
    /**
     * 
     * @type {V1GitType}
     * @memberof V1Init
     */
    git?: V1GitType;
    /**
     * 
     * @type {V1DockerfileType}
     * @memberof V1Init
     */
    dockerfile?: V1DockerfileType;
    /**
     * 
     * @type {V1FileType}
     * @memberof V1Init
     */
    file?: V1FileType;
    /**
     * 
     * @type {V1TensorboardType}
     * @memberof V1Init
     */
    tensorboard?: V1TensorboardType;
    /**
     * 
     * @type {string}
     * @memberof V1Init
     */
    lineageRef?: string;
    /**
     * 
     * @type {string}
     * @memberof V1Init
     */
    artifactRef?: string;
    /**
     * 
     * @type {string}
     * @memberof V1Init
     */
    modelRef?: string;
    /**
     * 
     * @type {string}
     * @memberof V1Init
     */
    connection?: string;
    /**
     * 
     * @type {string}
     * @memberof V1Init
     */
    path?: string;
    /**
     * 
     * @type {object}
     * @memberof V1Init
     */
    container?: object;
}

/**
 * Check if a given object implements the V1Init interface.
 */
export function instanceOfV1Init(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function V1InitFromJSON(json: any): V1Init {
    return V1InitFromJSONTyped(json, false);
}

export function V1InitFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1Init {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'artifacts': !exists(json, 'artifacts') ? undefined : V1ArtifactsTypeFromJSON(json['artifacts']),
        'paths': !exists(json, 'paths') ? undefined : json['paths'],
        'git': !exists(json, 'git') ? undefined : V1GitTypeFromJSON(json['git']),
        'dockerfile': !exists(json, 'dockerfile') ? undefined : V1DockerfileTypeFromJSON(json['dockerfile']),
        'file': !exists(json, 'file') ? undefined : V1FileTypeFromJSON(json['file']),
        'tensorboard': !exists(json, 'tensorboard') ? undefined : V1TensorboardTypeFromJSON(json['tensorboard']),
        'lineageRef': !exists(json, 'lineageRef') ? undefined : json['lineageRef'],
        'artifactRef': !exists(json, 'artifactRef') ? undefined : json['artifactRef'],
        'modelRef': !exists(json, 'modelRef') ? undefined : json['modelRef'],
        'connection': !exists(json, 'connection') ? undefined : json['connection'],
        'path': !exists(json, 'path') ? undefined : json['path'],
        'container': !exists(json, 'container') ? undefined : json['container'],
    };
}

export function V1InitToJSON(value?: V1Init | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'artifacts': V1ArtifactsTypeToJSON(value.artifacts),
        'paths': value.paths,
        'git': V1GitTypeToJSON(value.git),
        'dockerfile': V1DockerfileTypeToJSON(value.dockerfile),
        'file': V1FileTypeToJSON(value.file),
        'tensorboard': V1TensorboardTypeToJSON(value.tensorboard),
        'lineageRef': value.lineageRef,
        'artifactRef': value.artifactRef,
        'modelRef': value.modelRef,
        'connection': value.connection,
        'path': value.path,
        'container': value.container,
    };
}

